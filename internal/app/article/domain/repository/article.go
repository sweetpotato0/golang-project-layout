package repository

import (
	"context"

	"layout/api/user/v1"
	"layout/internal/app/article/domain/entity"
)

// ArticleRepository represent repository of the article
// Expect implementation by the infrastructure layer
type ArticleRepository interface {
	Get(ctx context.Context, id int) (*entity.Article, error)
	GetArticleUser(ctx context.Context, id int64) (*v1.GetUserReply, error)
	GetAll(ctx context.Context) ([]*entity.Article, error)
	Save(ctx context.Context, article *entity.Article) error
}
