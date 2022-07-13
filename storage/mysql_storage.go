package storage

import "gorm.io/gorm"

type mysqlStorage struct {
	db *gorm.DB
}

func NewMySQLStorage(db *gorm.DB) *mysqlStorage {
	return &mysqlStorage{db: db}
}
