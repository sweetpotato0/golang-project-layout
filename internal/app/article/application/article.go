package application

import (
	"context"

	"layout/internal/app/article/domain/entity"
	"layout/internal/app/article/domain/repository"
	"layout/internal/app/article/domain/vo"
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
func (i Article) GetArticle(ctx context.Context, id int) (*vo.Article, error) {
	a, err := i.articleRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &vo.Article{
		ID:      a.ID,
		Title:   a.Title,
		Content: a.Content,
	}, nil
}

// GetUsers returns user list
func (i Article) GetArticles(ctx context.Context) ([]*vo.Article, error) {
	as, err := i.articleRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	articles := []*vo.Article{}
	for _, a := range as {
		articles = append(articles, &vo.Article{
			ID:      a.ID,
			Title:   a.Title,
			Content: a.Content,
		})
	}

	return articles, nil
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
