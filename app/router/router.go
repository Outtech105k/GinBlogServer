package router

import (
	"app/handler"
	"app/template"

	"github.com/gin-gonic/gin"
)

func SetRouting(router *gin.Engine) error {
	var err error

	router.HTMLRender, err = template.CreatePageRender()
	if err != nil {
		return err
	}

	router.GET("/", handler.TopPageHandler)
	router.GET("/blog", handler.BlogListPageHandler)
	router.GET("/blog/:id", handler.BlogArticlePageHandler)

	// Error response
	router.NoRoute(handler.NotFoundHandler)

	return nil
}
