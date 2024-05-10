package main

<<<<<<< HEAD
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")
	engine.Static("/static", "./static")
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message": "hello gin",
		})
	})
	engine.Run(":80")
=======
import "github.com/gin-gonic/gin"

import "net/http"

func main() {
    engine:= gin.Default()
    engine.LoadHTMLGlob("templates/*")
    engine.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{
            "message": "hello gin",
        })
    })
    engine.Run(":80")
>>>>>>> origin/main
}
