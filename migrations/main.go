package main

import (
	"crawl/database"
	"crawl/models"
	"crawl/util"
	"fmt"
)

func main() {
	Migrate()
}

func Migrate() {
	config, err := util.LoadConfig(".")
	if err != nil {
		fmt.Println("not load config", err)
		panic(err)
	}

	db, err := database.DBConn(config)
	db.AutoMigrate(models.Article{}, models.Tag{}, models.Website{})
}
