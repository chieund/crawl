package controller

import (
	"bytes"
	"crawl/business"
	"crawl/models"
	"crawl/pkg"
	mysqlStorage "crawl/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"text/template"
	"time"
)

func (controller *Controller) Sitemap(db *gorm.DB) gin.HandlerFunc {
	storage := mysqlStorage.NewMySQLStorage(db)
	articleBU := business.NewArticleBusiness(storage)

	return func(c *gin.Context) {
		var pagination pkg.Pagination
		pagination.Limit = 10000
		articles, err := articleBU.GetAllArticles(&pagination)
		if err != nil {
			fmt.Println("article list empty")
		}

		var articleResponses []models.ArticleResponse
		for _, article := range articles.Rows {
			articleResponse := models.ArticleResponse{}
			articleResponse.Link = article.Link
			articleResponse.Slug = article.Slug
			articleResponse.UpdateAt = article.UpdatedAt.Format(time.RFC3339)
			articleResponses = append(articleResponses, articleResponse)
		}

		t := template.Must(template.New("sitemap.xml").ParseFiles("./templates/sitemap/sitemap.xml"))
		var b bytes.Buffer
		t.Execute(&b, gin.H{
			"articles": articleResponses,
		})

		c.Data(http.StatusOK, "text/xml", b.Bytes())
	}
}

func (controller *Controller) SitemapTags(db *gorm.DB) gin.HandlerFunc {
	storage := mysqlStorage.NewMySQLStorage(db)
	tagBu := business.NewTagBusiness(storage)

	return func(c *gin.Context) {
		tags, err := tagBu.GetAllTags()
		if err != nil {
			fmt.Println("article list empty")
		}

		var tagResponses []models.TagResponse
		for _, article := range tags {
			tagResponse := models.TagResponse{}
			tagResponse.Slug = article.Slug
			tagResponse.UpdateAt = article.UpdatedAt.Format(time.RFC3339)
			tagResponses = append(tagResponses, tagResponse)
		}

		t := template.Must(template.New("tags.xml").ParseFiles("./templates/sitemap/tags.xml"))
		var b bytes.Buffer
		t.Execute(&b, gin.H{
			"tags": tagResponses,
		})

		c.Data(http.StatusOK, "text/xml", b.Bytes())
	}
}
