package entity

import (
	"fmt"
	"strconv"

	"layout/internal/app/article/domain/po"
	"layout/internal/pkg/entity"
)

// Article represent entity of the article
type Article struct {
	entity.Entity
	po.Article
}

// Identity .
func (a *Article) Identity() string {
	return strconv.Itoa(a.ID)
}

// NewArticle initialize Article
func NewArticle(title, content string) (*Article, error) {
	if title == "" || content == "" {
		return nil, fmt.Errorf("invalid title or content")
	}

	a := &Article{Article: po.Article{Title: title, Content: content}}
	return a, nil
}
