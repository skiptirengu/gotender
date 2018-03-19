package models

import (
	"github.com/jinzhu/gorm"
	"github.com/skiptirengu/gotender/database"
)

type Model interface {
	Create() *gorm.DB
}

func Migrate() (*gorm.DB, error) {
	db, err := database.GetDb()
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(User{})
	if db.Error != nil {
		return nil, db.Error
	}

	db.AutoMigrate(Token{})
	if db.Error != nil {
		return nil, db.Error
	}

	return db, nil
}
