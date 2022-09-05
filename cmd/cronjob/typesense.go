package cronjob

import (
	"crawl/pkg/typesense"
	"crawl/util"
	"fmt"
	"github.com/spf13/cobra"
)

var TypeSenseCmd = &cobra.Command{
	Use:   "typesense",
	Short: "TypeSense",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := util.LoadConfig(".")
		if err != nil {
			fmt.Println("not load config", err)
			panic(err)
		}

		fmt.Println(args)
		CreateSchema(config)
		CreateDocument(config)
	},
}

func CreateSchema(config util.Config) {
	typesenseService := typesense.NewTypesenseService(config)
	typesenseService.CreateSchema()
}

func CreateDocument(config util.Config) {
	typesenseService := typesense.NewTypesenseService(config)

	typeDocument := typesense.DocumentTypesense{
		Id:    1,
		Title: "test",
		Slug:  "slug",
		Image: "img",
		Link:  "Link",
	}
	typesenseService.CreateDocument(typeDocument)
}
