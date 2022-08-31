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
	r.StaticFile("/css/styles.css", "./templates/css/styles.css")
	r.StaticFile("/robots.txt", "./templates/robots.txt")
	r.StaticFile("/images/image.webp", "./templates/images/image.webp")
	r.StaticFile("/images/icon/dev-to.png", "./templates/images/icon/dev-to.png")
	r.StaticFile("/images/icon/hashnode.jpeg", "./templates/images/icon/hashnode.jpeg")
	r.StaticFile("/images/icon/freecodecam.jpeg", "./templates/images/icon/freecodecam.jpeg")
	r.StaticFile("/images/icon/logrocket.png", "./templates/images/icon/logrocket.png")
	r.StaticFile("/images/icon/infoq.png", "./templates/images/icon/infoq.png")
	r.StaticFile("/google18bb3160694ad28a.html", "./templates/google18bb3160694ad28a.html")
	r.GET("/", controller.GetAllArticles(db))
	r.GET("/tags", controller.GetAllTags(db))
	r.GET("/sitemap.xml", controller.Sitemap(db))
	r.GET("/tags.xml", controller.SitemapTags(db))
	r.GET("/sources/:website_slug", controller.GetArticleBySource(db))
	r.GET("/t/:tag", controller.GetArticleByTag(db))
	r.GET("/:slug", controller.GetArticleBySlug(db))

	r.Run(":9000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
