package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func getdbconn() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/go?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db.LogMode(true)
	return db
}

type Message struct {
	ID      int    `gorm:"AUTO_INCREMENT"`
	Message string `gorm:"size:255"`
}

type human struct {
	Name    string `form:"name" binding:"required"`
	Message string `form:"message" binding:"required"`
}

func (human human) Said() string {
	return human.Name + "さんが" + human.Message + "と申しております"
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
		/*name := c.PostForm("name")
		messageVal := c.PostForm("message")*/

		var humanObj human
		// bindingする
		if err := c.ShouldBind(&humanObj); err != nil {
			// エラーだったらエラーで返す
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"title": "エラー",
			})
			return
		}
		// エラー処理
		if humanObj.Name == "" || humanObj.Message == "" {
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"title": "入力がされていません",
			})
			return
		}

		/*humanObj := human{
		Name:    name,
		Message: messageVal}*/
		said := humanObj.Said()
		db := getdbconn()
		defer db.Close()
		// テーブルの作成
		db.AutoMigrate(&Message{})
		message := Message{}
		message.ID = 0
		message.Message = said
		db.Create(&message)
		messages := []Message{}
		db.Find(&messages)
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "ようこそ！！",
			"msg":   messages,
		})
	})
	router.Run(":8080")
}
