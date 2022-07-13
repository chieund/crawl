package main

import (
	"crawl/business"
	"crawl/database"
	articleStorage "crawl/storage"
	"crawl/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		fmt.Println("not load config", err)
	}

	db := database.DBConn(config)
	storage := articleStorage.NewMySQLStorage(db)
	biz := business.NewArticleBusiness(storage)

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {

		articles, err := biz.GetAllArticles()
		if err != nil {
			fmt.Println("article list empty")
		}
		fmt.Println(articles)

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
