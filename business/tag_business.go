package business

import (
	"crawl/models"
)

type TagStorageInterface interface {
	FindTag(map[string]interface{}) (*models.Tag, error)
	UpdateTag(map[string]interface{}, models.Tag) bool
	CreateTag(models.Tag)
	GetAllTags() ([]models.Tag, error)
	Test() (models.Tag, error)
}

type TagBusiness struct {
	tagStore TagStorageInterface
}

func NewTagBusiness(tagStore TagStorageInterface) *TagBusiness {
	return &TagBusiness{
		tagStore: tagStore,
	}
}

func (tagBusiness *TagBusiness) FindTag(condition map[string]interface{}) (*models.Tag, error) {
	tag, err := tagBusiness.tagStore.FindTag(condition)

	if err != nil {
		return nil, err
	}

	return tag, nil
}

func (tagBusiness *TagBusiness) UpdateTag(condition map[string]interface{}, tag models.Tag) bool {
	return tagBusiness.tagStore.UpdateTag(condition, tag)
}

func (tagBusiness *TagBusiness) CreateTag(tag models.Tag) {
	tagBusiness.tagStore.CreateTag(tag)
}

func (tagBusiness *TagBusiness) GetAllTags() ([]models.Tag, error) {
	tags, err := tagBusiness.tagStore.GetAllTags()
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (tagBusiness *TagBusiness) Test() (models.Tag, error) {
	// get by tag
	return models.Tag{}, nil
}
