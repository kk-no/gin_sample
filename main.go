package main

import (
	"gin/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.tmpl")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "gin-gonic sample"})
	})

	user := r.Group("/user")
	{
		user.GET("/", handler.GetUser)
	}
	r.Run(":8001")
}
