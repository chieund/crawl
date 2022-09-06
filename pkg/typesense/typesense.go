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
			{
				Name: "is_update_content",
				Type: "int32",
			},
			{
				Name: "tags",
				Type: "string[]",
			},
			{
				Name: "website",
				Type: "string",
			},
		},
	}

	data, err := typesenseService.client.Collections().Create(schema)
	fmt.Println(data, err)
}

func (typesenseService *TypesenseService) CreateDocument(dType ArticleJson) (map[string]interface{}, error) {
	result, err := typesenseService.client.Collection(schemaArticle).Documents().Create(dType)
	if err != nil {
		fmt.Println(err)
	}
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

func (typesenseService *TypesenseService) Search(keyword string, queryByField string) (*api.SearchResult, error) {
	searchParameters := &api.SearchCollectionParams{
		Q:       keyword,
		QueryBy: queryByField,
	}
	result, _ := typesenseService.client.Collection(schemaArticle).Documents().Search(searchParameters)
	return result, nil
}
