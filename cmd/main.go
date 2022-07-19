package main

import (
	"crawl/business"
	"crawl/database"
	"crawl/models"
	"crawl/pkg/crawl"
	articleStorage "crawl/storage"
	"crawl/util"
	"fmt"
)

//const PAGE_LASTER = "latest"
const DOMAIN_CRAWL string = "https://dev.to"

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
	biz := business.NewArticleBusiness(storage)

	bizTag := business.NewTagBusiness(storage)

	articleTagBiz := business.NewArticleTagBusiness(storage)

	dataResult := crawl.CrawlWeb(DOMAIN_CRAWL)
	insertData(dataResult, biz, bizTag, articleTagBiz)

	// crawl freecodecamp
	dataResultFreeCodeCamp := crawl.CrawlWebFreeCodeCamp()
	insertData(dataResultFreeCodeCamp, biz, bizTag, articleTagBiz)

	medium := crawl.CrawlWebMedium()
	insertData(medium, biz, bizTag, articleTagBiz)
}

func insertData(dataResult []crawl.DataArticle, biz *business.ArticleBusiness, bizTag *business.TagBusiness, articleTagBiz *business.ArticleTagBusiness) {
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

		article, err := biz.FindArticle(map[string]interface{}{"slug": data.Slug})
		if err != nil {
			//fmt.Println("insert article: ", data.Title)
			article := models.Article{
				Title: data.Title,
				Slug:  data.Slug,
				Image: data.Image,
				Link:  data.Link,
			}
			biz.CreateArticle(&article)
			// insert article_tag
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
				}
			}
		} else {
			//fmt.Println("update article: ", article.Title)
			biz.UpdateArticle(map[string]interface{}{"slug": data.Slug}, *article)
		}
	}
}
