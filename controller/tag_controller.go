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

func (controller *Controller) GetAllTags(db *gorm.DB) gin.HandlerFunc {
	storage := mysqlStorage.NewMySQLStorage(db)
	tagBu := business.NewTagBusiness(storage)

	return func(c *gin.Context) {
		tags, err := tagBu.GetAllTags()
		if err != nil {
			fmt.Println("tags list empty")
		}

		c.HTML(http.StatusOK, "tags.tmpl", gin.H{
			"title": "List tags",
			"tags":  tags,
		})
	}
}

func (controller *Controller) GetArticleByTag(db *gorm.DB) gin.HandlerFunc {
	storage := mysqlStorage.NewMySQLStorage(db)
	articleBU := business.NewArticleBusiness(storage)
	tagBu := business.NewTagBusiness(storage)

	return func(c *gin.Context) {
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
		pagination.Link = "/t/"
		articles, err := articleBU.GetAllArticlesByIds(articleTags, &pagination)
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
	}
}
