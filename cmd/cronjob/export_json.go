package cronjob

import (
	"crawl/business"
	"crawl/database"
	"crawl/pkg"
	"crawl/service"
	mysqlStorage "crawl/storage"
	"crawl/util"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var ExportJsonCmd = &cobra.Command{
	Use:   "export-json",
	Short: "export json",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := util.LoadConfig(".")
		if err != nil {
			fmt.Println("not load config", err)
			panic(err)
		}

		db, err := database.DBConn(config)
		if err != nil {
			panic(err)
		}

		storage := mysqlStorage.NewMySQLStorage(db)
		articleBU := business.NewArticleBusiness(storage)
		var pagination pkg.Pagination
		pagination.Limit = 5
		articles, err := articleBU.GetAllArticles(&pagination)
		if err != nil {
			fmt.Println("article list empty")
		}

		articleService := service.NewArticleService(articles)
		articleResponses := articleService.FormatData()

		for _, article := range articleResponses {
			articles1, _ := json.Marshal(article)

			//err = ioutil.WriteFile("output.jsonl", articles1, 0644)
			f, err := os.OpenFile("text.jsonl",
				os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}
			defer f.Close()
			if _, err := f.WriteString(string(articles1) + "\n"); err != nil {
				log.Println(err)
			}
		}
	},
}
