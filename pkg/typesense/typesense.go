package typesense

import (
	"crawl/util"
	"fmt"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
	"os"
)

var (
	schemaArticle = "articles"
)

type TypesenseService struct {
	client *typesense.Client
}

func NewTypesenseService(config util.Config) *TypesenseService {
	client := typesense.NewClient(
		typesense.WithServer(config.TYPESENSE_URL),
		typesense.WithAPIKey(config.TYPESENSE_API_KEY))

	return &TypesenseService{
		client: client,
	}
}

func (typesenseService *TypesenseService) CreateSchema() {
	//test := "id"
	schema := &api.CollectionSchema{
		Name: schemaArticle,
		Fields: []api.Field{
			{
				Name: "id",
				Type: "int32",
			},
			{
				Name: "title",
				Type: "string",
			},
			{
				Name: "slug",
				Type: "string",
			},
			{
				Name: "image",
				Type: "string",
			},
			{
				Name: "link",
				Type: "string",
			},
		},
		//DefaultSortingField: &test,
	}
	fmt.Println(schema)

	data, err := typesenseService.client.Collections().Create(schema)
	fmt.Println(data, err)
}

type DocumentTypesense struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Slug  string `json:"slug"`
	Image string `json:"image"`
	Link  string `json:"link"`
}

func (typesenseService *TypesenseService) CreateDocument(dType DocumentTypesense) (map[string]interface{}, error) {
	result, err := typesenseService.client.Collection(schemaArticle).Documents().Create(dType)
	if err != nil {
		fmt.Println(err)
	}
	return result, nil
}

func (typesenseService *TypesenseService) Search(params string) (*api.SearchResult, error) {
	searchParameters := &api.SearchCollectionParams{
		Q:       params,
		QueryBy: "title",
	}
	result, _ := typesenseService.client.Collection(schemaArticle).Documents().Search(searchParameters)
	return result, nil
}

func (typesenseService *TypesenseService) ImportJson(fileName string) {
	action := "create"
	batchSize := 40
	params := &api.ImportDocumentsParams{
		Action:    &action,
		BatchSize: &batchSize,
	}
	importBody, err := os.Open(fileName)
	fmt.Println(importBody, err)
	fmt.Println(typesenseService.client.Collection(schemaArticle).Documents().ImportJsonl(importBody, params))
}
