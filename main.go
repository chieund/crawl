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
	r.StaticFile("/favicon.ico", "./templates/favicon.ico")
	r.StaticFile("/robots.txt", "./templates/robots.txt")
	r.StaticFile("/images/image.webp", "./templates/images/image.webp")
	r.GET("/", controller.GetAllArticles(db))
	r.GET("/tags", controller.GetAllTags(db))
	r.GET("/sitemap.xml", controller.Sitemap(db))
	r.GET("/tags.xml", controller.SitemapTags(db))
	r.GET("/t/:tag", controller.GetArticleByTag(db))
	r.GET("/:slug", controller.GetArticleBySlug(db))

	r.Run(":80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
