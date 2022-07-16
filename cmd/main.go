package main

import (
	"crawl/business"
	"crawl/database"
	"crawl/models"
	"crawl/pkg"
	articleStorage "crawl/storage"
	"crawl/util"
	"crawl/web_crawl"
	"fmt"
)

//const PAGE_LASTER = "latest"
const DOMAIN_CRAWL string = "https://dev.to"

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
	storage := articleStorage.NewMySQLStorage(db)
	biz := business.NewArticleBusiness(storage)
	fmt.Println(storage)

	dataResult := pkg.CrawlWeb(DOMAIN_CRAWL)
	for _, data := range dataResult {
		article, err := biz.FindArticle(map[string]interface{}{"slug": data.Slug})
		if err != nil {
			fmt.Println("insert article: ", data.Title)
			article := models.Article{
				Title: data.Title,
				Slug:  data.Slug,
				Image: data.Image,
				Link:  data.Link,
			}
			biz.CreateArticle(article)
		} else {
			fmt.Println("update article: ", article.Title)
			biz.UpdateArticle(map[string]interface{}{"slug": data.Slug}, *article)
		}
	}
}
