package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("template/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "ようこそ！！",
		})
	})
	router.POST("/post", func(c *gin.Context) {
		name := c.PostForm("name")
		message := c.PostForm("message")
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":   "ようこそ！！",
			"name":    name,
			"message": message,
		})
	})
	router.Run(":8080")
}
