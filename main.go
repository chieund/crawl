package main

import (
	IndexAction "crawl/controller"
	"crawl/database"
	"crawl/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path"
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

	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	cwd, _ := os.Getwd()
	r.LoadHTMLGlob(path.Join(cwd, "templates/*.tmpl"))
	controller := IndexAction.Controller{}
	r.GET("/", controller.GetAllArticles(db))
	r.GET("/tags", controller.GetAllTags(db))
	// add sitemap
	r.GET("/sitemap.xml", func(c *gin.Context) {
		var pagination pkg.Pagination
		page := c.Request.URL.Query().Get("page")
		pagination.Page, _ = strconv.Atoi(page)
		pagination.Link = "/"
		articles, err := articleBU.GetAllArticles(&pagination)
		if err != nil {
			fmt.Println("article list empty")
		}
		fmt.Println(articles)

		t := template.Must(template.New("sitemap.xml").ParseFiles("./templates/sitemap.xml"))
		var b bytes.Buffer
		t.Execute(&b, gin.H{
			"title":    "List tags",
			"articles": articles,
		})

		c.Data(http.StatusOK, "application/xml", b.Bytes())
	})

	r.GET("/t/:tag", controller.GetArticleByTag(db))
	r.GET("/:slug", controller.GetArticleBySlug(db))

	r.Run(":80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
