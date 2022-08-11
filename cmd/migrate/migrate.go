package migrate

import (
	"crawl/database"
	"crawl/models"
	"crawl/util"
	"fmt"
	"github.com/spf13/cobra"
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate",
	Run: func(cmd *cobra.Command, args []string) {
		Migrate()
	},
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
