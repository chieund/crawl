package crawl

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gosimple/slug"
	"strings"
)

const DOMAIN_INFOQ_CRAWL string = "https://www.infoq.com"

var listCrawlInfoQTags = []string{
	//"news",
	"articles",
}

func CrawlInfoQWeb(ch chan []DataArticle) {
	result := func() []DataArticle {
		c := colly.NewCollector()

		var dataArticles []DataArticle
		c.OnHTML(".card__content", func(e *colly.HTMLElement) {
			dataArticle := DataArticle{}
			dataArticle.Title = e.ChildText("h3.card__title a")
			link := e.ChildAttr("h3.card__title a", "href")

			var tags []DataTag
			e.ForEach(".card__topics", func(_ int, e *colly.HTMLElement) {
				tagText := e.Text
				title := strings.Replace(tagText, "#", "", 1)
				title = strings.TrimSpace(title)
				tag := DataTag{}
				tag.Title = title
				tag.Slug = slug.Make(title)
				tags = append(tags, tag)
			})
			dataArticle.Tags = tags

			// get tags
			dataArticle.Slug = slug.Make(dataArticle.Title)
			dataArticle.Link = DOMAIN_INFOQ_CRAWL + link
			dataArticle.WebsiteId = 6
			dataArticle.WebsiteSlug = "infoq-com"
			dataArticles = append(dataArticles, dataArticle)
			//c.Visit(DOMAIN_INFOQ_CRAWL + link)
		})

		//c.OnHTML(".article__data", func(e *colly.HTMLElement) {
		//	content := e.DOM
		//	content.Find(".related__group").Remove()
		//	content.Find("script").Remove()
		//	content.Find(":input").Remove()
		//	fmt.Println(content.Html())
		//	return
		//})

		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting\n", r.URL)
		})

		for _, tag := range listCrawlInfoQTags {
			c.Visit(DOMAIN_INFOQ_CRAWL + "/" + tag)
		}
		return dataArticles
	}()

	ch <- result
	defer close(ch)
}

func CrawlWebInfoQContent(url string) DataArticle {
	c := colly.NewCollector()

	var dataArticle DataArticle
	c.OnHTML("html", func(e *colly.HTMLElement) {
		content := e.DOM
		content = content.Find(".article__data")
		content.Find(".related__group").Remove()
		content.Find("script").Remove()
		content.Find(":input").Remove()
		contentClean, _ := content.Html()
		dataArticle.Content = strings.TrimSpace(contentClean)

		imgURL, _ := e.DOM.Find(`meta[property="og:image"]`).Attr("content")
		dataArticle.Image = imgURL
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Print("Visiting\n", r.URL)
	})

	c.Visit(url)
	return dataArticle
}
