package models

type ArticleTag struct {
	ArticleId int `gorm:"primaryKey" column:"article_id"`
	TagId     int `gorm:"primaryKey" column:"article_id"`
}

func (articleTag *ArticleTag) TableName() string {
	return "article_tag"
}
