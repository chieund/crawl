package cronjob

import (
	"crawl/business"
	"crawl/database"
	"crawl/models"
	"crawl/pkg"
	"crawl/pkg/crawl"
	//"crawl/pkg/typesense"

	//"crawl/pkg/typesense"
	articleStorage "crawl/storage"
	"crawl/util"
	"fmt"
	"github.com/spf13/cobra"
	//"strconv"
	"strings"
)

var CrawlArticleCmd = &cobra.Command{
	Use:   "crawl-article",
	Short: "Crawl web",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		CrawlArticle()
	},
}

const DOMAIN_CRAWL string = "https://dev.to"

func CrawlArticle() {
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

	bizTag := business.NewTagBusiness(storage)

	articleTagBiz := business.NewArticleTagBusiness(storage)

	devToChan := make(chan []crawl.DataArticle)
	hashNode := make(chan []crawl.DataArticle)
	webFreeCodeCamp := make(chan []crawl.DataArticle)
	medium := make(chan []crawl.DataArticle)
	logrocket := make(chan []crawl.DataArticle)

	infoq := make(chan []crawl.DataArticle)
	go crawl.CrawlInfoQWeb(infoq)
	go crawl.CrawlWeb(devToChan)
	go crawl.CrawlWebFreeCodeCamp(webFreeCodeCamp)
	go crawl.CrawlWebMedium(medium)
	go crawl.CrawlWebHashNode(hashNode)
	go crawl.CrawlLogrocketWeb(logrocket)

	insertData(config, <-webFreeCodeCamp, biz, bizTag, articleTagBiz)
	insertData(config, <-devToChan, biz, bizTag, articleTagBiz)
	insertData(config, <-medium, biz, bizTag, articleTagBiz)
	insertData(config, <-hashNode, biz, bizTag, articleTagBiz)
	insertData(config, <-logrocket, biz, bizTag, articleTagBiz)
	insertData(config, <-infoq, biz, bizTag, articleTagBiz)
}

func insertData(config util.Config, dataResult []crawl.DataArticle, biz *business.ArticleBusiness, bizTag *business.TagBusiness, articleTagBiz *business.ArticleTagBusiness) {
	count := 0

	//typesenseService := typesense.NewTypesenseService(config)

	for _, data := range dataResult {
		if len(data.Tags) > 0 {
			for _, dataTag := range data.Tags {
				tag, err := bizTag.FindTag(map[string]interface{}{"slug": dataTag.Slug})
				if err != nil {
					tag := models.Tag{
						Title: dataTag.Title,
						Slug:  dataTag.Slug,
					}
					bizTag.CreateTag(tag)
				} else {
					bizTag.UpdateTag(map[string]interface{}{"slug": dataTag.Slug}, *tag)
				}
			}
		}

		check := strings.Contains(data.Slug, "go")
		if check && count < 5 {
			pkg.BotPushNewGoToDiscord(config, data.Title, data.Link, data.Image)
		}

		article, err := biz.FindArticleCron(map[string]interface{}{"slug": data.Slug})
		if err != nil {
			//fmt.Println("insert article: ", data.Title)
			article := models.Article{
				Title:       data.Title,
				Slug:        data.Slug,
				Image:       data.Image,
				Link:        data.Link,
				WebsiteId:   data.WebsiteId,
				WebsiteSlug: data.WebsiteSlug,
			}
			biz.CreateArticle(&article)
			// insert article_tag
			var tagJsons = []string{}
			if len(data.Tags) > 0 {
				for _, dataTag := range data.Tags {
					tag, err := bizTag.FindTag(map[string]interface{}{"slug": dataTag.Slug})
					if err == nil {
						// insert article with tag
						articleTag := models.ArticleTag{
							ArticleId: article.Id,
							TagId:     tag.Id,
						}
						articleTagBiz.CreateArticleTag(&articleTag)
					}

					tagJsons = append(tagJsons, tag.Title)
				}
			}

			//typeDocument := typesense.ArticleJson{
			//	ID:        strconv.Itoa(article.Id),
			//	Title:     article.Title,
			//	Slug:      article.Slug,
			//	Image:     article.Image,
			//	Link:      article.Link,
			//	Tags:      tagJsons,
			//	Website:   article.Website.Slug,
			//	CreatedAt: article.CreatedAt.Format("2006-01-02 15:04:05"),
			//	UpdatedAt: article.CreatedAt.Format("2006-01-02 15:04:05"),
			//}
			//_, err := typesenseService.CreateDocument(typeDocument)
			//if err != nil {
			//	fmt.Println("not create typesense", article.Slug)
			//}
		} else {
			//fmt.Println("update article: ", article.Title)
			biz.UpdateArticle(map[string]interface{}{"slug": data.Slug}, *article)
		}

		count++
	}
}
