package repository

import (
	"context"
	"database/sql"

	"layout/api/user/v1"
	"layout/configs"
	"layout/internal/app/article/domain/entity"
	"layout/internal/app/article/domain/po"
	"layout/internal/app/article/domain/repository"
	fentity "layout/internal/pkg/entity"
	"layout/internal/pkg/store"
	"layout/pkg/transport/grpc"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

//实现领域模型内的依赖倒置
var _ repository.ArticleRepository = (*Article)(nil)

// Article Implements repository.ArticleRepository
type Article struct {
	conn   *sql.DB
	userv1 v1.UserClient
	cache  store.EntityCache //实体缓存组件
}

// NewArticle .
func NewArticle(c *configs.Configuration) repository.ArticleRepository {
	conn, err := sql.Open("mysql", c.DB.Addr)
	if err != nil {
		panic(err)
	}

	a := &Article{conn: conn, cache: store.New(c.Redis), userv1: newUserClient(c)}
	a.GetCacheFunc()

	return a
}

func newUserClient(c *configs.Configuration) v1.UserClient {
	ctx := context.Background()
	conn, err := grpc.NewGrpcConn(ctx, "user_service", c)
	if err != nil {
		panic(err)
	}

	return v1.NewUserClient(conn)
}

// GetCacheFunc .
func (repo *Article) GetCacheFunc() {
	//设置缓存的持久化数据源
	repo.cache.SetSource(func(result fentity.Entity) error {
		a := result.(*entity.Article)
		ctx := context.Background()
		row, err := repo.queryRow(ctx, "select id, title, content from articles where id=?", a.ID)
		if err != nil {
			return err
		}

		if err := row.Scan(&a.ID, &a.Title, &a.Content); err != nil {
			return err
		}
		return nil
	})
}

// Get .
func (repo *Article) Get(ctx context.Context, id int) (*entity.Article, error) {

	article := &entity.Article{Article: po.Article{ID: id}}

	return article, repo.cache.GetEntity(article)
}

// GetArticleUser .
func (repo *Article) GetArticleUser(ctx context.Context, id int64) (*v1.GetUserReply, error) {
	req := &v1.GetUserReq{
		Id: id,
	}
	return repo.userv1.GetUser(ctx, req)
}

// GetAll .
func (repo *Article) GetAll(ctx context.Context) ([]*entity.Article, error) {
	rows, err := repo.query(ctx, "select id, title, content from articles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	as := make([]*entity.Article, 0)
	for rows.Next() {
		a := &entity.Article{}
		err = rows.Scan(&a.ID, &a.Title, &a.Content)
		if err != nil {
			return nil, err
		}
		as = append(as, a)
	}
	return as, nil
}

// Save 保存订单实体
func (repo *Article) Save(ctx context.Context, articleEntity *entity.Article) error {

	if articleEntity.ID == 0 { //新建

		stmt, err := repo.conn.Prepare("insert into articles (title, content) values (?, ?)")
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.ExecContext(ctx, articleEntity.Title, articleEntity.Content)
		repo.cache.Delete(articleEntity)
		return err
	}

	stmt, err := repo.conn.Prepare("update articles set title=?, content=? where id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, articleEntity.Title, articleEntity.Content, articleEntity.ID)
	return err

	repo.cache.Delete(articleEntity)
	return nil
}

func (repo *Article) query(ctx context.Context, q string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := repo.conn.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.QueryContext(ctx, args...)
}

func (repo *Article) queryRow(ctx context.Context, q string, args ...interface{}) (*sql.Row, error) {
	stmt, err := repo.conn.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.QueryRowContext(ctx, args...), nil
}
