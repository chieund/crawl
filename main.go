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
		panic(err)
	}

	db, err := database.DBConn(config)
	if err != nil {
		panic(err)
	}
	storage := articleStorage.NewMySQLStorage(db)
	biz := business.NewArticleBusiness(storage)

	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		articles, err := biz.GetAllArticles()
		if err != nil {
			fmt.Println("article list empty")
		}

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":    "Crawl Web",
			"articles": articles,
		})
	})
	r.Run(":80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
