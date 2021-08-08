package repository

import (
	"context"
	"database/sql"

	"layout/configs"
	"layout/internal/app/user/domain/entity"
	"layout/internal/app/user/domain/po"
	"layout/internal/app/user/domain/repository"
	fentity "layout/internal/pkg/entity"
	"layout/internal/pkg/store"

	// mysql driver init
	_ "github.com/go-sql-driver/mysql"
)

//实现领域模型内的依赖倒置
var _ repository.UserRepository = (*User)(nil)

// User Implements repository.UserRepository
type User struct {
	conn  *sql.DB
	cache store.EntityCache //实体缓存组件
}

// NewUser create user repository.
func NewUser(c *configs.Configuration) repository.UserRepository {
	conn, err := sql.Open("mysql", c.DB.Addr)
	if err != nil {
		panic(err)
	}

	a := &User{conn: conn, cache: store.New(c.Redis)}
	a.GetCacheFunc()

	return a
}

// GetCacheFunc .
func (repo *User) GetCacheFunc() {
	//设置缓存的持久化数据源
	repo.cache.SetSource(func(result fentity.Entity) error {
		a := result.(*entity.User)
		ctx := context.Background()
		row, err := repo.queryRow(ctx, "select id, name from user where id=?", a.ID)
		if err != nil {
			return err
		}

		if err := row.Scan(&a.ID, &a.Name); err != nil {
			return err
		}
		return nil
	})
}

// GetUser .
func (repo *User) GetUser(ctx context.Context, id int64) (*entity.User, error) {

	user := &entity.User{User: po.User{ID: id}}

	return user, repo.cache.GetEntity(user)
}

func (repo *User) queryRow(ctx context.Context, q string, args ...interface{}) (*sql.Row, error) {
	stmt, err := repo.conn.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.QueryRowContext(ctx, args...), nil
}
