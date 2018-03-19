package database

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jinzhu/gorm"
)

func Open() (*gorm.DB) {
	if db, err := GetDb(); err != nil {
		panic(err)
	} else {
		return db
	}
}

func GetDb() (*gorm.DB, error) {
	if db, err := gorm.Open("sqlite3", "database/database.sqlite"); err != nil {
		return nil, err
	} else {
		db.LogMode(true)
		return db, nil
	}
}
