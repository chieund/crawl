package crawl

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gosimple/slug"
)

const URL_FREECODECAMP = "https://www.freecodecamp.org"

func CrawlWebFreeCodeCamp(ch chan []DataArticle) {
	c := colly.NewCollector()

	var dataArticles []DataArticle
	c.OnHTML("article.post-card", func(e *colly.HTMLElement) {
		dataArticle := DataArticle{}

		dataArticle.Title = e.ChildText("h2.post-card-title a")
		dataArticle.Image = e.ChildAttr("img.post-card-image", "src")
		link := e.ChildAttr("h2.post-card-title a", "href")
		dataArticle.Link = URL_FREECODECAMP + link
		dataArticle.Slug = slug.Make(dataArticle.Title)
		dataArticles = append(dataArticles, dataArticle)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Print("Visiting\n", r.URL)
	})

	c.Visit(URL_FREECODECAMP + "/news")
	ch <- dataArticles
	defer close(ch)
}
