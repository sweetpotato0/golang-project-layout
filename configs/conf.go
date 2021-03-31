package configs

import (
	"runtime"
	"sync"

	"github.com/BurntSushi/toml"
)

var once sync.Once
var cfg *Configuration

// Get .
func Get() *Configuration {
	once.Do(func() {
		cfg = &Configuration{
			Server: newServerConf(),
			DB:     newDBConf(),
			Redis:  newRedisConf(),
		}
	})
	return cfg
}

// Configuration .
type Configuration struct {
	Server *ServerConf
	DB     *DBConf
	Redis  *RedisConf
}

type ServerConf struct {
	Mode string
	Http *HTTPConf
	Grpc *GRPCConf
}

type HTTPConf struct {
	Addr    string
	Port    int
	Timeout int
}

type GRPCConf struct {
	Addr    string
	Port    int
	Timeout int
}

// DBConf .
type DBConf struct {
	Addr            string `toml:"addr"`
	MaxOpenConns    int    `toml:"max_open_conns"`
	MaxIdleConns    int    `toml:"max_idle_conns"`
	ConnMaxLifeTime int    `toml:"conn_max_life_time"`
}

// RedisConf .
type RedisConf struct {
	Addr               string `toml:"addr"`
	Password           string `toml:"password"`
	DB                 int    `toml:"db"`
	MaxRetries         int    `toml:"max_retries"`
	PoolSize           int    `toml:"pool_size"`
	ReadTimeout        int    `toml:"read_timeout"`
	WriteTimeout       int    `toml:"write_timeout"`
	IdleTimeout        int    `toml:"idle_timeout"`
	IdleCheckFrequency int    `toml:"idle_check_frequency"`
	MaxConnAge         int    `toml:"max_conn_age"`
	PoolTimeout        int    `toml:"pool_timeout"`
}

func newServerConf() *ServerConf {
	result := &ServerConf{}
	if _, err := toml.DecodeFile("configs/server.toml", result); err != nil {
		panic(err)
	}
	return result
}

func newDBConf() *DBConf {
	result := &DBConf{}
	if _, err := toml.DecodeFile("configs/db.toml", result); err != nil {
		panic(err)
	}
	return result
}

func newRedisConf() *RedisConf {
	result := &RedisConf{
		MaxRetries:         0,
		PoolSize:           10 * runtime.NumCPU(),
		ReadTimeout:        3,
		WriteTimeout:       3,
		IdleTimeout:        300,
		IdleCheckFrequency: 60,
	}
	if _, err := toml.DecodeFile("configs/redis.toml", result); err != nil {
		panic(err)
	}

	return result
}
