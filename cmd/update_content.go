package main

import (
	"crawl/business"
	"crawl/database"
	"crawl/pkg"
	"crawl/pkg/crawl"
	articleStorage "crawl/storage"
	"crawl/util"
	"fmt"
)

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

	var paging pkg.Pagination
	paging.Page = 1
	artiles, _ := biz.GetAllArticlesByIds([]int{1, 2, 3, 4, 5}, &paging)
	fmt.Println(artiles.Rows)
	for _, article := range artiles.Rows {
		content := crawl.CrawlWebDevContent(article.Link)
		fmt.Println(content)
	}
}
