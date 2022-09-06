package controller

import (
	"crawl/business"
	"crawl/models"
	"crawl/pkg/typesense"
	mysqlStorage "crawl/storage"
	"crawl/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func (controller *Controller) Search(config util.Config, db *gorm.DB) gin.HandlerFunc {
	storage := mysqlStorage.NewMySQLStorage(db)
	tagBu := business.NewTagBusiness(storage)

	return func(c *gin.Context) {
		keyword := c.Request.URL.Query().Get("q")

		typesenseService := typesense.NewTypesenseService(config)
		articles, _ := typesenseService.Search(keyword, "title")

		var articleResponses []models.ArticleResponse
		for _, value := range *articles.Hits {
			document := value.Document

			articleResponse := models.ArticleResponse{}
			str, _ := (*document)["title"].(string)
			image, _ := (*document)["image"].(string)

			articleResponse.Title = str
			articleResponse.Image = image
			articleResponses = append(articleResponses, articleResponse)
		}

		tags, _ := tagBu.GetAllHotTags()
		c.HTML(http.StatusOK, "result_search.tmpl", gin.H{
			"title":       "The Best Developer News",
			"description": "The Best Developer News is a website that aggregates all the latest articles on technology",
			"keywords":    "Angular, Aws, blockchain, ci/cd, css, Data Science, Django, GoLang, Java, Javascript, Laravel, Mmagento, Node.js, Php, Python, React, Rust, Serverless, Vuejs, Web Development",
			"articles":    articleResponses,
			"pagination":  articles,
			"currentPage": articles.Page,
			"tags":        tags,
		})
	}
}
