package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func serveHTTP() error {
	router := gin.Default()
	var err error

	router.HTMLRender, err = createPageRender()
	if err != nil {
		return err
	}

	router.Static("/static", "./static")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index", gin.H{})
	})
	router.GET("/blog", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "blogArticleList", gin.H{
			"pageTitle": "記事一覧",
		})
	})

	// Error response
	router.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "404", gin.H{
			"pageTitle": "404 Error",
		})
	})

	return router.Run(":80")
}
