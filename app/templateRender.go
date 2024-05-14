package main

import "github.com/gin-contrib/multitemplate"

func createPageRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	r.AddFromFiles("index", "templates/base.html", "templates/content/topPage.html")
	r.AddFromFiles("blogArticleList", "templates/base.html", "templates/content/customContent.html", "templates/content/contentBody/blogArticleList.html")

	r.AddFromFiles("404", "templates/base.html", "templates/content/customContent.html", "templates/content/contentBody/404.html")

	return r
}