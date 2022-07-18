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
	articleBU := business.NewArticleBusiness(storage)
	tagBu := business.NewTagBusiness(storage)

	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	cwd, _ := os.Getwd()
	r.LoadHTMLGlob(path.Join(cwd, "templates/*"))
	r.GET("/", func(c *gin.Context) {
		articles, err := articleBU.GetAllArticles()
		if err != nil {
			fmt.Println("article list empty")
		}

		tags, err := tagBu.GetAllTags()

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":    "Crawl Web",
			"articles": articles,
			"tags":     tags,
		})
	})

	r.GET("/t/:tag", func(c *gin.Context) {
		tagName := c.Param("tag")

		tags, err := tagBu.GetAllTags()
		if err != nil {
			fmt.Println("tags list empty")
		}

		tag, err := tagBu.FindTag(map[string]interface{}{"Title": tagName})
		if err != nil {
			fmt.Println("tags list empty")
		}

		// get all article_tag by tag_id
		articleTagBU := business.NewArticleTagBusiness(storage)
		articleTags := articleTagBU.FindArticleIdByTagId(tag.Id)
		articles, err := articleBU.GetAllArticlesByIds(articleTags)
		if err != nil {
			fmt.Println("not load article by tag")
		}

		c.HTML(http.StatusOK, "tags.tmpl", gin.H{
			"title":    "Crawl Web",
			"articles": articles,
			"tags":     tags,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
