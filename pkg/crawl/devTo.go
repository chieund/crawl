package crawl

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gosimple/slug"
	"strings"
)

const DOMAIN_CRAWL string = "https://dev.to"

var listCrawlDevTags = []string{
	"javascript",
	"webdev",
	"react",
	"python",
	"devops",
	"css",
	"typescript",
	"java",
	"php",
	"blockchain",
	"database",
	"go",
	"aws",
}

func CrawlWeb(ch chan []DataArticle) {
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

		var tags []DataTag
		e.ForEach("a.crayons-tag", func(_ int, e *colly.HTMLElement) {
			tagText := e.Text
			title := strings.Replace(tagText, "#", "", 1)
			tag := DataTag{}
			tag.Title = title
			tag.Slug = slug.Make(title)
			tags = append(tags, tag)
		})
		dataArticle.Tags = tags

		// get tags
		dataArticle.Slug = slug.Make(dataArticle.Title)
		dataArticle.Link = DOMAIN_CRAWL + link
		dataArticles = append(dataArticles, dataArticle)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Print("Visiting\n", r.URL)
	})

	c.Visit(DOMAIN_CRAWL)
	for _, tag := range listCrawlDevTags {
		c.Visit(DOMAIN_CRAWL + "/t/" + tag)
	}
	ch <- dataArticles
	defer close(ch)
}
