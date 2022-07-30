package controller

import (
	"bytes"
	"crawl/business"
	"crawl/pkg"
	mysqlStorage "crawl/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"text/template"
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
		fmt.Println(articles)

		t := template.Must(template.New("sitemap.xml").ParseFiles("./templates/sitemap.xml"))

		var b bytes.Buffer
		t.Execute(&b, gin.H{
			"articles": articles,
		})

		c.Data(http.StatusOK, "application/xml", b.Bytes())
	}
}
