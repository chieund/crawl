package controller

import (
	"crawl/business"
	"crawl/pkg"
	mysqlStorage "crawl/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/grokify/html-strip-tags-go"
	"gorm.io/gorm"
	"html/template"
	"net/http"
	"strconv"
	"strings"
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
		pagination.Limit = 40
		articles, err := articleBU.GetAllArticles(&pagination)
		if err != nil {
			fmt.Println("article list empty")
		}

		tags, err := tagBu.GetAllHotTags()

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":       "The Best Developer News",
			"description": "The Best Developer News is a website that aggregates all the latest articles on technology",
			"keywords":    "Angular, Aws, blockchain, ci/cd, css, Data Science, Django, GoLang, Java, Javascript, Laravel, Mmagento, Node.js, Php, Python, React, Rust, Serverless, Vuejs, Web Development",
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
	tagBu := business.NewTagBusiness(storage)

	return func(c *gin.Context) {
		slug := c.Param("slug")
		article, err := articleBU.FindArticle(map[string]interface{}{"slug": slug})
		if err != nil {
			c.Redirect(http.StatusNotFound, "/")
		}

		article.Viewed = article.Viewed + 1
		articleBU.UpdateArticle(map[string]interface{}{"slug": slug}, *article)

		if article.IsUpdateContent != 1 {
			c.Redirect(http.StatusFound, article.Link)
		}
		ContentArticle := template.HTML(article.Content)
		tags, err := tagBu.GetAllHotTags()

		// get list tags of article
		var tagId []int
		for _, tag := range article.Tags {
			tagId = append(tagId, tag.Id)
		}

		var pagination pkg.Pagination
		pagination.Limit = 42
		pagination.Condition = map[string]interface{}{"slug": slug}
		articleOthers, err := articleBU.FindArticleOther(tagId, &pagination)
		c.HTML(http.StatusOK, "article_detail.tmpl", gin.H{
			"title":          article.Title + "- The Best Developer News",
			"description":    clearContentDescription(article.Content, 170),
			"keywords":       "Angular, Aws, blockchain, ci/cd, css, Data Science, Django, GoLang, Java, Javascript, Laravel, Mmagento, Node.js, Php, Python, React, Rust, Serverless, Vuejs, Web Development",
			"article":        article,
			"articleOthers":  articleOthers,
			"ContentArticle": ContentArticle,
			"tags":           tags,
		})
	}
}

func clearContentDescription(content string, length int) string {
	content = strip.StripTags(content)
	if len(content) > length {
		content = content[0:length]
	}
	content = strings.TrimSuffix(content, "\r\n")
	content = strings.Replace(content, "\n", "", -1)
	content = strings.TrimSpace(content) + "..."
	return content
}
