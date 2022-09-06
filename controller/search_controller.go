package controller

import (
	"crawl/business"
	"crawl/models"
	"crawl/pkg"
	"crawl/pkg/typesense"
	mysqlStorage "crawl/storage"
	"crawl/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math"
	"net/http"
	"strconv"
)

func (controller *Controller) Search(config util.Config, db *gorm.DB) gin.HandlerFunc {
	storage := mysqlStorage.NewMySQLStorage(db)
	tagBu := business.NewTagBusiness(storage)

	return func(c *gin.Context) {
		keyword := c.Request.URL.Query().Get("q")
		page := c.Request.URL.Query().Get("page")
		fmt.Println(page)

		typesenseService := typesense.NewTypesenseService(config)
		articles, _ := typesenseService.Search(keyword, "title")

		var pagination = pkg.Pagination{}
		pagination.Page, _ = strconv.Atoi(page)
		pagination.Link = fmt.Sprintf("/search?q=%s", keyword)
		pagination.Limit = 10
		pagination.TotalRows = int64(*articles.Found)
		pagination.TotalPages = int(math.Ceil(float64(pagination.TotalRows) / float64(pagination.GetLimit())))
		pagination.SetListPages()
		pagination.GetPage()

		var articleResponses []models.ArticleResponse
		for _, value := range *articles.Hits {
			document := value.Document
			highlights := value.Highlights

			articleResponse := models.ArticleResponse{}
			str, _ := (*document)["title"].(string)
			image, _ := (*document)["image"].(string)
			is_update_content, _ := (*document)["is_update_content"].(int)
			link, _ := (*document)["link"].(string)
			slug, _ := (*document)["slug"].(string)

			snippet := *(*highlights)[0].Snippet

			articleResponse.Title = str
			articleResponse.Snippet = snippet
			articleResponse.Image = image
			articleResponse.IsUpdateContent = is_update_content
			articleResponse.Link = link
			articleResponse.Slug = slug
			articleResponses = append(articleResponses, articleResponse)
		}

		tags, _ := tagBu.GetAllHotTags()
		c.HTML(http.StatusOK, "result_search.tmpl", gin.H{
			"title":       fmt.Sprintf("Search Results for %s - The Best Developer News", keyword),
			"description": "The Best Developer News is a website that aggregates all the latest articles on technology",
			"keywords":    "Angular, Aws, blockchain, ci/cd, css, Data Science, Django, GoLang, Java, Javascript, Laravel, Mmagento, Node.js, Php, Python, React, Rust, Serverless, Vuejs, Web Development",
			"articles":    articleResponses,
			"pagination":  pagination,
			"currentPage": pagination.Page,
			"listPage":    pagination.ListPages,
			"tags":        tags,
			"keyword":     keyword,
		})
	}
}
