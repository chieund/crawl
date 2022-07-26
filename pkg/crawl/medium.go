package crawl

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gosimple/slug"
	"strings"
)

const (
	URL_MEDIUM = "https://medium.com"
)

var listTags = []string{
	"programming",
	"python",
	"technology",
	"software-engineering",
	"web-development",
	"javascript",
	"software-development",
}

func CrawlWebMedium(ch chan []DataArticle) {
	result := func() []DataArticle {
		c := colly.NewCollector()
		var dataArticles []DataArticle
		c.OnHTML("article", func(e *colly.HTMLElement) {
			dataArticle := DataArticle{}

			dataArticle.Title = e.ChildText("h2")
			link := e.ChildAttr("a[aria-label='Post Preview Title']", "href")

			newLink := strings.Split(link, "?")

			image := e.ChildAttr("a[aria-label='Post Preview Image'] div img", "src")
			dataArticle.Image = strings.ReplaceAll(image, "fit/c/224/224", "max/1400")
			dataArticle.Link = URL_MEDIUM + newLink[0]
			dataArticle.Slug = slug.Make(dataArticle.Title)

			// get tags
			var tags []DataTag
			e.ForEach("div.lc.dq.ho.it.iu.iv.bn.b.ld.bp.fu.iw", func(_ int, e *colly.HTMLElement) {
				tagText := e.Text
				title := strings.Replace(tagText, "#", "", 1)
				tag := DataTag{}
				tag.Title = title
				tag.Slug = slug.Make(title)
				tags = append(tags, tag)
			})
			dataArticle.Tags = tags

			dataArticles = append(dataArticles, dataArticle)
		})

		c.OnRequest(func(r *colly.Request) {
			fmt.Print("Visiting\n", r.URL)
		})

		for _, tag := range listTags {
			c.Visit(URL_MEDIUM + "/tag/" + tag)
		}
		return dataArticles
	}()
	ch <- result
	defer close(ch)
}
