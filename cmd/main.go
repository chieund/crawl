package main

import (
	"crawl/business"
	"crawl/database"
	"crawl/models"
	articleStorage "crawl/storage"
	"crawl/util"
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gosimple/slug"
)

//const PAGE_LASTER = "latest"
const DOMAIN_CRAWL = "https://dev.to"

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

	c := colly.NewCollector()
	c.OnHTML(".crayons-story", func(e *colly.HTMLElement) {
		title := e.ChildText("h2.crayons-story__title a")
		link := e.ChildAttr("h2.crayons-story__title a", "href")
		image := e.ChildAttr("h2.crayons-story__title a", "data-preload-image")

		if image == "" {
			image = "https://thepracticaldev.s3.amazonaws.com/i/6hqmcjaxbgbon8ydw93z.png"
		}

		// get tags
		slug := slug.Make(title)
		fmt.Println(title, slug, DOMAIN_CRAWL+link)

		biz := business.NewArticleBusiness(storage)
		article, err := biz.FindArticle(map[string]interface{}{"slug": slug})
		if err != nil {
			fmt.Println("insert article: ", title)

			article := models.Article{
				Title: title,
				Slug:  slug,
				Image: image,
				Link:  DOMAIN_CRAWL + link,
			}
			biz.CreateArticle(article)
		} else {
			fmt.Println("update article: ", article.Title)
			biz.UpdateArticle(map[string]interface{}{"slug": slug}, *article)
		}
		//c.Visit(DOMAIN_CRAWL + link)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Print("Visiting\n", r.URL)
	})

	c.Visit(DOMAIN_CRAWL)
}
