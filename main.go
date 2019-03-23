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

type article struct {
	gorm.Model
	Title      string `gorm:"size:255"`
	Content    string `gorm:"size:255"`
	Image_path string `gorm:"size:255"`
}

func getContents(c *gin.Context) {
	// 記事の配列を宣言
	articles := []article{}
	// 接続を確立
	db := getdbconn()
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&article{})
	// 接続を後で閉じる
	defer db.Close()
	// 全件取得
	db.Find(&articles)
	// view呼び出し
	c.HTML(http.StatusOK, "top.tmpl", gin.H{
		"articles": articles,
	})
}

func main() {
	router := gin.Default()
	// templateフォルダから読み込む
	router.LoadHTMLGlob("template/*")
	// imgフォルダのパスを通す
	router.Static("/img", "./img")
	// cssフォルダのパスを通す
	router.Static("/css", "./css")

	router.GET("/", getContents)

	router.Run(":8080")
}
