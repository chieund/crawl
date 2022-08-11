package crawl

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gosimple/slug"
	"strings"
)

const (
	URL_HASHNODE = "https://hashnode.com"
)

var listCrawlTags = []string{
	"javascript",
	"web-development",
	"css",
	"reactjs",
	"python",
	"javascript",
	"html5",
	"devops",
	"go",
	"html5",
}

func CrawlWebHashNode(ch chan []DataArticle) {
	result := func() []DataArticle {
		c := colly.NewCollector()
		var dataArticles []DataArticle
		c.OnHTML("div.css-4gdbui", func(e *colly.HTMLElement) {
			dataArticle := DataArticle{}

			dataArticle.Title = e.ChildText("h1.css-1j1qyv3 a.css-4zleql")
			link := e.ChildAttr("h1.css-1j1qyv3 a.css-4zleql", "href")
			image := e.ChildAttr("div.css-qnvenm img", "style")
			if image != "" {
				imageNew := strings.Split(image, "background-image:url(")
				imageNew = strings.Split(imageNew[1], "?")
				imageNewReplace := imageNew[0]
				imageNewPlace := strings.ReplaceAll(imageNewReplace, `"`, "")
				dataArticle.Image = imageNewPlace
			}
			dataArticle.Link = link
			dataArticle.Slug = slug.Make(dataArticle.Title)

			// get tags
			var tags []DataTag
			e.ForEach("div.css-1r9abvi a.css-83n4vj", func(_ int, e *colly.HTMLElement) {
				title := e.Text
				tag := DataTag{}
				tag.Title = title
				tag.Slug = slug.Make(title)
				tags = append(tags, tag)
			})
			dataArticle.Tags = tags
			dataArticles = append(dataArticles, dataArticle)
		})

		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting\n", r.URL)
		})

		//c.Visit(URL_HASHNODE + "/community")
		for _, tag := range listCrawlTags {
			c.Visit(URL_HASHNODE + "/n/" + tag)
		}
		return dataArticles
	}()
	ch <- result
	defer close(ch)
}

func CrawlWebHashNodeContent(url string) DataArticle {
	c := colly.NewCollector()

	var dataArticle DataArticle
	c.OnHTML("#post-content-wrapper", func(e *colly.HTMLElement) {
		content, _ := e.DOM.Html()
		dataArticle.Content = content
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting\n", r.URL)
	})

	c.Visit(url)
	return dataArticle
}
