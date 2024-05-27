package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
)

func blogPageHandler(ctx *gin.Context) {
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
}
