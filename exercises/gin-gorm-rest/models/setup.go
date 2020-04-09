package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// SetupModels - Connect with database & migrate model's schema
func SetupModels() *gorm.DB {
	//Opens connection to DB
	db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	//Migrate database schema (keep schema up to date)
	db.AutoMigrate(&Book{})

	return db
}
