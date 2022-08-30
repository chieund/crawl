package cronjob

import (
	"crawl/business"
	"crawl/database"
	"crawl/pkg"
	"crawl/pkg/crawl"
	articleStorage "crawl/storage"
	"crawl/util"
	"fmt"
	"github.com/spf13/cobra"
)

var CrawlArticleDetailCmd = &cobra.Command{
	Use:   "crawl-article-detail",
	Short: "Crawl Article Detail",
	Run: func(cmd *cobra.Command, args []string) {
		CrawlArticleDetail()
	},
}

//const PAGE_LASTER = "latest"
func CrawlArticleDetail() {
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

	var paging pkg.Pagination
	paging.Page = 1
	paging.Sort = "created_at desc"
	paging.Condition = map[string]interface{}{
		"website_slug": []string{
			"dev-to",
			"freecodecamp-org",
			"hashnode-com",
			"logrocket-com",
			"infoq-com",
		},
		"is_update_content": 0,
	}
	artiles, _ := biz.GetAllArticles(&paging)
	for _, article := range artiles.Rows {
		var content crawl.DataArticle
		switch article.WebsiteSlug {
		case "dev-to":
			content = crawl.CrawlWebDevContent(article.Link)
		case "freecodecamp-org":
			content = crawl.CrawlWebFreeCodeCampContent(article.Link)
		case "hashnode-com":
			content = crawl.CrawlWebHashNodeContent(article.Link)
		case "logrocket-com":
			content = crawl.CrawlWebLogrocketContent(article.Link)
		case "infoq-com":
			content = crawl.CrawlWebInfoQContent(article.Link)
		}

		//// find article by id
		if len(content.Content) > 0 {
			articleFind, err := biz.FindArticle(map[string]interface{}{"id": article.Id})
			if err != nil {
				fmt.Println("article not found", article.Id, article.Slug, article.Title)
			}

			articleFind.Content = content.Content
			if article.WebsiteSlug == "infoq-com" {
				articleFind.Image = content.Image
			}

			articleFind.IsUpdateContent = 1
			biz.UpdateArticle(map[string]interface{}{"id": article.Id}, *articleFind)
		} else {
			fmt.Println("Content url", article.Link, " empty")
		}
	}
}
