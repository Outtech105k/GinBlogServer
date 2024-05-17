package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
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

	// TODO: 正規ページ化
	router.GET("/blog/:id", func(ctx *gin.Context) {
		mdBytes, err := os.ReadFile("./static/md/" + ctx.Param("id") + ".md")
		if err != nil {
			ctx.HTML(http.StatusNotFound, "404", nil)
			log.Print(err)
			return
		}
		html := string(blackfriday.MarkdownCommon(mdBytes))

		ctx.HTML(http.StatusOK, "blogArticle", gin.H{
			"pageTitle":   "テスト記事",
			"articleBody": template.HTML(html),
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
