package pkg

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gosimple/slug"
)

const DOMAIN_CRAWL string = "https://dev.to"

type DataTag struct {
	Name string
}

type DataArticle struct {
	Title string
	Slug  string
	Link  string
	Image string
	Tags  []DataTag
}

func CrawlWeb(url string) []DataArticle {
	c := colly.NewCollector()

	var dataArticles []DataArticle
	c.OnHTML(".crayons-story", func(e *colly.HTMLElement) {
		dataArticle := DataArticle{}
		dataArticle.Title = e.ChildText("h2.crayons-story__title a")
		link := e.ChildAttr("h2.crayons-story__title a", "href")
		dataArticle.Image = e.ChildAttr("h2.crayons-story__title a", "data-preload-image")
		if dataArticle.Image == "" {
			dataArticle.Image = "https://thepracticaldev.s3.amazonaws.com/i/6hqmcjaxbgbon8ydw93z.png"
		}

		e.ForEach("a.crayons-tag__prefix", func(_ int, e *colly.HTMLElement) {

		})

		// get tags
		dataArticle.Slug = slug.Make(dataArticle.Title)
		dataArticle.Link = DOMAIN_CRAWL + link
		dataArticles = append(dataArticles, dataArticle)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Print("Visiting\n", r.URL)
	})

	c.Visit(url)
	return dataArticles
}
