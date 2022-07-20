package main

import (
	"crawl/business"
	"crawl/database"
	"crawl/pkg"
	mysqlStorage "crawl/storage"
	"crawl/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"strconv"
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
	r.LoadHTMLGlob(path.Join(cwd, "templates/*.tmpl"))
	r.GET("/", func(c *gin.Context) {
		var pagination pkg.Pagination
		page := c.Request.URL.Query().Get("page")
		pagination.Page, _ = strconv.Atoi(page)
		pagination.Link = "/"
		articles, err := articleBU.GetAllArticles(&pagination)
		if err != nil {
			fmt.Println("article list empty")
		}

		tags, err := tagBu.GetAllHotTags()

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":       "The Best Developer News",
			"pagination":  articles,
			"currentPage": articles.Page,
			"listPage":    articles.ListPages,
			"tags":        tags,
		})
	})

	r.GET("/tags", func(c *gin.Context) {
		tags, err := tagBu.GetAllTags()
		if err != nil {
			fmt.Println("tags list empty")
		}

		c.HTML(http.StatusOK, "tags.tmpl", gin.H{
			"title": "List tags",
			"tags":  tags,
		})
	})

	r.GET("/t/:tag", func(c *gin.Context) {
		tagName := c.Param("tag")
		page := c.Request.URL.Query().Get("page")

		tags, err := tagBu.GetAllHotTags()
		if err != nil {
			fmt.Println("tags list empty")
		}

		tag, err := tagBu.FindTag(map[string]interface{}{"slug": tagName})
		if err != nil {
			fmt.Println("tags list empty")
		}

		// get all article_tag by tag_id
		articleTagBU := business.NewArticleTagBusiness(storage)
		articleTags := articleTagBU.FindArticleIdByTagId(tag.Id)

		var pagination pkg.Pagination

		pagination.Page, _ = strconv.Atoi(page)
		pagination.Link = "/tag/"
		articles, err := articleBU.GetAllArticlesByIds(articleTags, &pagination)
		fmt.Println(articles)
		if err != nil {
			fmt.Println("not load article by tag")
		}

		c.HTML(http.StatusOK, "article_tags.tmpl", gin.H{
			"title":       tag.Title,
			"pagination":  articles,
			"currentPage": articles.Page,
			"listPage":    articles.ListPages,
			"tags":        tags,
			"tag":         tag,
		})
	})

	r.Run(":80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
