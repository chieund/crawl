package main

import (
	"crawl/business"
	"crawl/database"
	mysqlStorage "crawl/storage"
	"crawl/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
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
	storage := mysqlStorage.NewMySQLStorage(db)
	biz := business.NewArticleBusiness(storage)
	tagBiz := business.NewTagBusiness(storage)

	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	cwd, _ := os.Getwd()
	r.LoadHTMLGlob(path.Join(cwd, "templates/*"))
	r.GET("/", func(c *gin.Context) {
		articles, err := biz.GetAllArticles()
		if err != nil {
			fmt.Println("article list empty")
		}

		tags, err := tagBiz.GetAllTags()

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":    "Crawl Web",
			"articles": articles,
			"tags":     tags,
		})
	})

	r.GET("/t/:tag", func(c *gin.Context) {
		tagName := c.Param("tag")

		tags, err := tagBiz.GetAllTags()
		if err != nil {
			fmt.Println("tags list empty")
		}

		articleTagBiz := business.NewArticleTagBusiness(storage)
		tags, err := tagBiz.GetArticleByTag(tagName, articleTagBiz)

		c.HTML(http.StatusOK, "tag.tmpl", gin.H{
			"title":    "Crawl Web",
			"articles": tags,
			"tags":     tags,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
