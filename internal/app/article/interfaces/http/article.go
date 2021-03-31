package http

import (
	"net/http"
	"strconv"

	"layout/internal/app/article/application"
	"layout/internal/app/article/domain/vo"

	fhttp "layout/pkg/transport/http"

	"github.com/gin-gonic/gin"
)

// Handler article handler
type Handler struct {
	UseCase *application.Article
}

// ArticleRoutes returns the initialized article router
func ArticleRoutes(engine *fhttp.Server, h Handler) {
	engine.GET("/article/:id", h.getArticle)
	engine.GET("/articles", h.getArticles)
	engine.POST("/article", h.createArticle)
}

func (h Handler) getArticle(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

	article, err := h.UseCase.GetArticle(ctx, id)
	if err != nil {
		ctx.String(http.StatusNotFound, err.Error()+" failed to get article")
		return
	}
	ctx.JSON(http.StatusOK, article)
}

func (h Handler) getArticles(ctx *gin.Context) {

	articles, err := h.UseCase.GetArticles(ctx)
	if err != nil {
		ctx.String(http.StatusNotFound, err.Error()+" failed to get article list")
		return
	}
	type payload struct {
		Articles []*vo.Article `json:"articles"`
	}
	ctx.JSON(http.StatusOK, payload{Articles: articles})
}

func (h Handler) createArticle(ctx *gin.Context) {

	type payload struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	var p payload
	if err := ctx.ShouldBind(&p); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if err := h.UseCase.AddArticle(ctx, p.Title, p.Content); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error()+" failed to create article")
		return
	}
	ctx.JSON(http.StatusCreated, nil)
}
