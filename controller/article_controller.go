package controller

import (
	"crawl/business"
	"crawl/models"
	"crawl/pkg"
	"crawl/service"
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

		articleService := service.NewArticleService(articles)
		articleResponses := articleService.FormatData()

		tags, err := tagBu.GetAllHotTags()

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":       "The Best Developer News",
			"description": "The Best Developer News is a website that aggregates all the latest articles on technology",
			"keywords":    "Angular, Aws, blockchain, ci/cd, css, Data Science, Django, GoLang, Java, Javascript, Laravel, Mmagento, Node.js, Php, Python, React, Rust, Serverless, Vuejs, Web Development",
			"articles":    articleResponses,
			"pagination":  articles,
			"currentPage": articles.Page,
			"listPage":    articles.ListPages,
			"tags":        tags,
			"tabActive":   1,
		})
	}
}

func (controller *Controller) GetAllArticlesByTop(db *gorm.DB) gin.HandlerFunc {
	storage := mysqlStorage.NewMySQLStorage(db)
	articleBU := business.NewArticleBusiness(storage)
	tagBu := business.NewTagBusiness(storage)

	return func(c *gin.Context) {
		var pagination pkg.Pagination
		page := c.Request.URL.Query().Get("page")
		pagination.Page, _ = strconv.Atoi(page)
		pagination.Link = "/top"
		pagination.Limit = 40
		pagination.Sort = "viewed desc"
		articles, err := articleBU.GetAllArticles(&pagination)
		if err != nil {
			fmt.Println("article list empty")
		}

		articleService := service.NewArticleService(articles)
		articleResponses := articleService.FormatData()

		tags, err := tagBu.GetAllHotTags()

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":       "The Best Developer News",
			"description": "The Best Developer News is a website that aggregates all the latest articles on technology",
			"keywords":    "Angular, Aws, blockchain, ci/cd, css, Data Science, Django, GoLang, Java, Javascript, Laravel, Mmagento, Node.js, Php, Python, React, Rust, Serverless, Vuejs, Web Development",
			"articles":    articleResponses,
			"pagination":  articles,
			"currentPage": articles.Page,
			"listPage":    articles.ListPages,
			"tags":        tags,
			"tabActive":   2,
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

		articleBU.UpdateViewed(map[string]interface{}{"slug": slug}, article.Viewed+1)

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

		// articleDetail
		articleDetail := models.ArticleResponse{
			Title:           article.Title,
			Link:            article.Link,
			Slug:            article.Slug,
			CreatedAt:       article.CreatedAt.Format("Jan 02"),
			Tags:            article.Tags,
			Image:           article.Image,
			IsUpdateContent: article.IsUpdateContent,
			Website:         article.Website,
		}

		articleService := service.NewArticleService(articleOthers)
		articleResponses := articleService.FormatData()

		c.HTML(http.StatusOK, "article_detail.tmpl", gin.H{
			"title":          article.Title + "- The Best Developer News",
			"description":    clearContentDescription(article.Content, 170),
			"keywords":       "Angular, Aws, blockchain, ci/cd, css, Data Science, Django, GoLang, Java, Javascript, Laravel, Mmagento, Node.js, Php, Python, React, Rust, Serverless, Vuejs, Web Development",
			"article":        articleDetail,
			"articleOthers":  articleResponses,
			"ContentArticle": ContentArticle,
			"tags":           tags,
		})
	}
}

func (controller *Controller) GetArticleBySource(db *gorm.DB) gin.HandlerFunc {
	storage := mysqlStorage.NewMySQLStorage(db)
	articleBU := business.NewArticleBusiness(storage)
	tagBu := business.NewTagBusiness(storage)

	return func(c *gin.Context) {
		websiteSlug := c.Param("website_slug")
		var pagination pkg.Pagination
		page := c.Request.URL.Query().Get("page")
		pagination.Page, _ = strconv.Atoi(page)
		pagination.Link = fmt.Sprintf("/sources/%s", websiteSlug)
		pagination.Limit = 40
		pagination.Condition = map[string]interface{}{"website_slug": websiteSlug}
		articles, err := articleBU.GetAllArticles(&pagination)
		if err != nil {
			fmt.Println("article list empty")
		}

		articleService := service.NewArticleService(articles)
		articleResponses := articleService.FormatData()

		tags, err := tagBu.GetAllHotTags()
		c.HTML(http.StatusOK, "article_source.tmpl", gin.H{
			"title":       "The Best Developer News",
			"description": "The Best Developer News is a website that aggregates all the latest articles on technology",
			"keywords":    "Angular, Aws, blockchain, ci/cd, css, Data Science, Django, GoLang, Java, Javascript, Laravel, Mmagento, Node.js, Php, Python, React, Rust, Serverless, Vuejs, Web Development",
			"articles":    articleResponses,
			"pagination":  articles,
			"currentPage": articles.Page,
			"listPage":    articles.ListPages,
			"tags":        tags,
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
