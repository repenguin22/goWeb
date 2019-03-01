package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type human struct {
	name, message string
}

func (human human) Said() string {
	return human.name + "さんが" + human.message + "と申しております"
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("template/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "ようこそ！！!",
		})
	})
	router.POST("/post", func(c *gin.Context) {
		name := c.PostForm("name")
		messageVal := c.PostForm("message")
		human := human{name, messageVal}
		said := human.Said()
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "ようこそ！！",
			"msg":   said,
		})
	})
	router.Run(":8080")
}
