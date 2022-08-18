package crawl

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gosimple/slug"
	"strings"
)

const DOMAIN_LOGROCKET_CRAWL string = "https://blog.logrocket.com/"

func CrawlLogrocketWeb(ch chan []DataArticle) {
	result := func() []DataArticle {
		c := colly.NewCollector()

		var dataArticles []DataArticle
		c.OnHTML(".grid-item", func(e *colly.HTMLElement) {
			dataArticle := DataArticle{}
			dataArticle.Title = e.ChildText("h2.card-title a")
			link := e.ChildAttr("h2.card-title a", "href")
			image := e.ChildAttr("a.thumbimage", "style")

			// get tags
			dataArticle.Image = clearImageLogrocket(image)
			dataArticle.Slug = slug.Make(dataArticle.Title)
			dataArticle.Link = link
			dataArticle.WebsiteId = 5
			dataArticle.WebsiteSlug = "logrocket-com"
			dataArticles = append(dataArticles, dataArticle)
		})

		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting\n", r.URL)
		})

		c.Visit(DOMAIN_LOGROCKET_CRAWL)
		return dataArticles
	}()
	ch <- result
	defer close(ch)
}

func CrawlWebLogrocketContent(url string) DataArticle {
	c := colly.NewCollector()

	var dataArticle DataArticle
	c.OnHTML(".article-post", func(e *colly.HTMLElement) {
		content, _ := e.DOM.Html()
		dataArticle.Content = clearContent(content)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Print("Visiting\n", r.URL)
	})

	c.Visit(url)
	return dataArticle
}

func clearImageLogrocket(image string) string {
	image = strings.Replace(image, "background-image:url(", "", 1)
	image = strings.Replace(image, ");", "", 1)
	return image
}
