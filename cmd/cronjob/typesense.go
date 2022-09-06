package cronjob

import (
	"crawl/pkg/typesense"
	"crawl/util"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/typesense/typesense-go/typesense/api"
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

		fmt.Println("args input:", args)
		if len(args) > 0 {
			task := args[0]
			switch task {
			case "create-schema":
				CreateSchema(config)
			case "create-doc":
				CreateDocument(config)
			case "import-json":
				ImportFileJson(config)
			case "search":
				fmt.Println(Search(config, "java", "title"))
			default:
				fmt.Println("Params not contain [create-schema, create-doc, import-json, search]")
			}
		} else {
			fmt.Println("Require args contain [create-schema, create-doc, import-json, search]")
		}
	},
}

func CreateSchema(config util.Config) {
	typesenseService := typesense.NewTypesenseService(config)
	typesenseService.CreateSchema()
}

func CreateDocument(config util.Config) {
	typesenseService := typesense.NewTypesenseService(config)
	typeDocument := typesense.ArticleJson{
		ID:      "2",
		Title:   "test",
		Slug:    "slug",
		Image:   "img",
		Link:    "Link",
		Tags:    []string{"java", "c#", "golang"},
		Website: "dev-to",
	}
	typesenseService.CreateDocument(typeDocument)
}

func ImportFileJson(config util.Config) {
	typesenseService := typesense.NewTypesenseService(config)
	typesenseService.ImportJson(filePathJson)
}

func Search(config util.Config, keyword string, queryFiledBy string) (*api.SearchResult, error) {
	typesenseService := typesense.NewTypesenseService(config)
	return typesenseService.Search(keyword, queryFiledBy)
}
