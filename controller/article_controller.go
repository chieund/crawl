package controller

import (
	"crawl/business"
	"crawl/pkg"
	mysqlStorage "crawl/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func (controller *Controller) GetAllArticles(db *gorm.DB) gin.HandlerFunc {
	storage := mysqlStorage.NewMySQLStorage(db)
	articleBU := business.NewArticleBusiness(storage)
	tagBu := business.NewTagBusiness(storage)

	return func(c *gin.Context) {
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
	}
}

func (controller *Controller) GetArticleBySlug(db *gorm.DB) gin.HandlerFunc {
	storage := mysqlStorage.NewMySQLStorage(db)
	articleBU := business.NewArticleBusiness(storage)

	return func(c *gin.Context) {
		slug := c.Param("slug")
		article, err := articleBU.FindArticle(map[string]interface{}{"slug": slug})
		if err != nil {
			c.Redirect(http.StatusNotFound, "/")
		}

		article.Viewed = article.Viewed + 1
		articleBU.UpdateArticle(map[string]interface{}{"slug": slug}, *article)
		c.Redirect(http.StatusFound, article.Link)
	}
}
