package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	// TODO: 正規ページ化
	router.GET("/blog/:id", blogPageHandler)

	// Error response
	router.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "404", gin.H{
			"pageTitle": "404 Error",
		})
	})

	srv := &http.Server{
		Addr:    ":80",
		Handler: router.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Panicf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")

	return nil
}
