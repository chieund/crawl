package controller

import (
	"crawl/business"
	"crawl/models"
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
			"title":       fmt.Sprintf("%s - The Best Developer News", "List tags"),
			"description": fmt.Sprintf("%s a website that aggregates all the latest articles on technology", "list tags"),
			"keywords":    "Angular, Aws, blockchain, ci/cd, css, Data Science, Django, GoLang, Java, Javascript, Laravel, Mmagento, Node.js, Php, Python, React, Rust, Serverless, Vuejs, Web Development",
			"tags":        tags,
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
		pagination.Link = fmt.Sprintf("/t/%s", tag.Slug)
		articles, err := articleBU.GetAllArticlesByIds(articleTags, &pagination)
		if err != nil {
			fmt.Println("not load article by tag")
		}

		var articleResponses []models.ArticleResponse
		for _, article := range articles.Rows {
			articleResponse := models.ArticleResponse{}
			articleResponse.Title = article.Title
			articleResponse.Link = article.Link
			articleResponse.Slug = article.Slug
			articleResponse.CreatedAt = article.CreatedAt.Format("Jan 02")
			articleResponse.Tags = article.Tags
			articleResponse.Image = article.Image
			articleResponse.IsUpdateContent = article.IsUpdateContent
			articleResponse.Website = article.Website
			articleResponses = append(articleResponses, articleResponse)
		}

		c.HTML(http.StatusOK, "article_tags.tmpl", gin.H{
			"title":       fmt.Sprintf("%s - The Best Developer News", tag.Title),
			"description": fmt.Sprintf("%s a website that aggregates all the latest articles on technology", tag.Title),
			"keywords":    fmt.Sprintf("Software development, engineering, Web Development, %s", tag.Title),
			"articles":    articleResponses,
			"pagination":  articles,
			"currentPage": articles.Page,
			"listPage":    articles.ListPages,
			"tags":        tags,
			"tag":         tag,
		})
	}
}
