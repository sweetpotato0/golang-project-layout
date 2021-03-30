package application

import (
	"context"

	"layout/internal/app/article/domain/entity"
	"layout/internal/app/article/domain/repository"
)

// Article provides use-case
type Article struct {
	articleRepo repository.ArticleRepository
	// userRepo    repository.UserRepository
}

func NewArticleUseCase(rep repository.ArticleRepository) *Article {
	return &Article{articleRepo: rep}
}

// GetUser returns user
func (i Article) GetArticle(ctx context.Context, id int) (*entity.Article, error) {
	return i.articleRepo.Get(ctx, id)
}

// GetUsers returns user list
func (i Article) GetArticles(ctx context.Context) ([]*entity.Article, error) {
	return i.articleRepo.GetAll(ctx)
}

// AddArticle saves new article
func (i Article) AddArticle(ctx context.Context, title, content string) error {
	a, err := entity.NewArticle(title, content)
	if err != nil {
		return err
	}
	return i.articleRepo.Save(ctx, a)
}

// NewArticle create a article use case
func NewArticle(repo repository.ArticleRepository) *Article {
	return &Article{
		articleRepo: repo,
	}
}
