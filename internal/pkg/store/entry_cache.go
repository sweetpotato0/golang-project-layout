package store

/*
	缓存组件，实现了一级缓存，二级缓存，防击穿.
*/
import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"time"

	"layout/configs"
	"layout/internal/pkg/entity"
	"layout/pkg/log"

	"github.com/go-redis/redis"
	"golang.org/x/sync/singleflight"
)

// EntityCache .
type EntityCache interface {
	//获取实体
	GetEntity(entity.Entity) error
	//删除实体缓存
	Delete(result entity.Entity, async ...bool) error
	//设置数据源
	SetSource(func(entity.Entity) error) EntityCache
	//设置前缀
	SetPrefix(string) EntityCache
	//设置缓存时间，默认5分钟
	SetExpiration(time.Duration) EntityCache
	//设置异步反写缓存。默认关闭，缓存未命中读取数据源后的异步反写缓存
	SetAsyncWrite(bool) EntityCache
	//设置防击穿，默认开启
	SetSingleFlight(bool) EntityCache
	//关闭二级缓存. 关闭后只有一级缓存生效
	CloseRedis() EntityCache
}

var _ EntityCache = (*EntityCacheImpl)(nil)

var group singleflight.Group

// EntityCacheImpl .
type EntityCacheImpl struct {
	asyncWrite   bool
	prefix       string
	expiration   time.Duration
	getEntity    func(result entity.Entity) error
	singleFlight bool
	client       redis.Cmdable
}

func New(cfg *configs.RedisConf) *EntityCacheImpl {

	opt := &redis.Options{
		Addr:               cfg.Addr,
		Password:           cfg.Password,
		DB:                 cfg.DB,
		MaxRetries:         cfg.MaxRetries,
		PoolSize:           cfg.PoolSize,
		ReadTimeout:        time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout:       time.Duration(cfg.WriteTimeout) * time.Second,
		IdleTimeout:        time.Duration(cfg.IdleTimeout) * time.Second,
		IdleCheckFrequency: time.Duration(cfg.IdleCheckFrequency) * time.Second,
		MaxConnAge:         time.Duration(cfg.MaxConnAge) * time.Second,
		PoolTimeout:        time.Duration(cfg.PoolTimeout) * time.Second,
	}
	redisClient := redis.NewClient(opt)
	if e := redisClient.Ping().Err(); e != nil {
		log.Fatal(e.Error())
	}

	return &EntityCacheImpl{
		asyncWrite:   false,
		expiration:   5 * time.Minute,
		singleFlight: true,
		client:       redisClient,
	}
}

// GetEntity 读取实体缓存
func (cache *EntityCacheImpl) GetEntity(result entity.Entity) error {
	value := reflect.ValueOf(result)
	name := cache.getName(value.Type()) + ":" + result.Identity()

	//可以做一级缓存读取：本地缓存

	//缓存读取
	entityBytes, err := cache.getRedis(name)
	if err != nil && err != redis.Nil {
		return err
	}
	if err != redis.Nil {
		err = json.Unmarshal(entityBytes, result)
		if err != nil {
			return err
		}
		return nil
	}

	//持久化数据源读取
	entityBytes, err = cache.getCall(name, result)
	if err != nil {
		return err
	}

	//反写缓存
	expiration := cache.expiration
	client := cache.client
	if client == nil {
		return nil
	}
	if !cache.asyncWrite {
		return client.Set(name, entityBytes, expiration).Err()
	}
	go func() {
		var err error
		defer func() {
			if perr := recover(); perr != nil {
				err = fmt.Errorf(fmt.Sprint(perr))
			}
			if err != nil {
				log.Errorf("Failed to set entity cache, name:%s err:%v, ", name, err)
			}
		}()
		err = client.Set(name, entityBytes, expiration).Err()
	}()

	return nil
}

// Delete 删除实体缓存
func (cache *EntityCacheImpl) Delete(result entity.Entity, async ...bool) error {
	name := cache.getName(reflect.ValueOf(result).Type()) + ":" + result.Identity()
	client := cache.client
	if client == nil {
		return nil
	}
	if len(async) == 0 {
		return client.Del(name).Err()
	}

	go func() {
		var err error
		defer func() {
			if perr := recover(); perr != nil {
				err = fmt.Errorf(fmt.Sprint(perr))
			}
			if err != nil {
				log.Errorf("Failed to delete entity cache, name:%s err:%v, ", name, err)
			}
		}()
		err = client.Del(name).Err()
	}()

	return nil
}

// SetSource 设置数据源
func (cache *EntityCacheImpl) SetSource(getEntity func(result entity.Entity) error) EntityCache {
	cache.getEntity = getEntity
	return cache
}

// SetAsyncWrite 设置异步写入,默认同步写入缓存。 当缓存未命中读取数据源后是否异步写入缓存
func (cache *EntityCacheImpl) SetAsyncWrite(open bool) EntityCache {
	cache.asyncWrite = open
	return cache
}

// SetPrefix 设置缓存实体前缀
func (cache *EntityCacheImpl) SetPrefix(prefix string) EntityCache {
	cache.prefix = prefix
	return cache
}

// SetExpiration 设置缓存实体时间 默认5分钟
func (cache *EntityCacheImpl) SetExpiration(expiration time.Duration) EntityCache {
	cache.expiration = expiration
	return cache
}

// SetSingleFlight 默认开启
func (cache *EntityCacheImpl) SetSingleFlight(open bool) EntityCache {
	cache.singleFlight = open
	return cache
}

// CloseRedis 关闭缓存
func (cache *EntityCacheImpl) CloseRedis() EntityCache {
	cache.client = nil
	return cache
}

func (cache *EntityCacheImpl) getName(entityType reflect.Type) string {
	for entityType.Kind() == reflect.Ptr {
		entityType = entityType.Elem()
	}
	if cache.prefix != "" {
		return cache.prefix + ":" + entityType.Name()
	}
	return entityType.Name()
}

func (cache *EntityCacheImpl) getRedis(name string) ([]byte, error) {
	if cache.client == nil {
		return nil, redis.Nil
	}
	client := cache.client
	if cache.singleFlight {
		entityData, err, _ := group.Do("cache:"+name, func() (interface{}, error) {
			return client.Get(name).Bytes()
		})
		if err != nil {
			return nil, err
		}
		return entityData.([]byte), err
	}
	return client.Get(name).Bytes()
}

func (cache *EntityCacheImpl) getCall(name string, result entity.Entity) ([]byte, error) {
	if cache.getEntity == nil {
		return nil, errors.New("Undefined source")
	}
	if cache.singleFlight {
		entityData, err, shared := group.Do("getEntity:"+name, func() (interface{}, error) {
			e := cache.getEntity(result)
			if e != nil {
				return nil, e
			}
			return json.Marshal(result)
		})
		if err != nil {
			return nil, err
		}
		if shared {
			entityByte, _ := entityData.([]byte)
			err = json.Unmarshal(entityByte, result)
			resultByte := make([]byte, len(entityByte))
			copy(resultByte, entityByte)
			return resultByte, err
		}
		return entityData.([]byte), err
	}
	err := cache.getEntity(result)
	if err != nil {
		return nil, err
	}
	return json.Marshal(result)
}
