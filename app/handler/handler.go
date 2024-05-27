package handler

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
)

func TopPageHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index", nil)
}

// TODO: 正規ページ化
func BlogArticlePageHandler(ctx *gin.Context) {
	mdBytes, err := os.ReadFile("./static/md/" + ctx.Param("id") + ".md")
	if err != nil {
		ctx.HTML(http.StatusNotFound, "404", gin.H{
			"pageTitle": "404 Error",
		})
		log.Print(err)
		return
	}
	html := string(blackfriday.MarkdownCommon(mdBytes))

	ctx.HTML(http.StatusOK, "blogArticle", gin.H{
		"pageTitle":   "テスト記事",
		"articleBody": template.HTML(html),
	})
}

func BlogListPageHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "blogArticleList", gin.H{
		"pageTitle": "記事一覧",
	})
}

func NotFoundHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusNotFound, "404", gin.H{
		"pageTitle": "404 Error",
	})
}
