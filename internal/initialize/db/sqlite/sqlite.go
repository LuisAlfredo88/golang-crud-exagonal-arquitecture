package sqlite

import (
	"golang-crud-exagonal-arquitecture/internal/initialize/db/dbentities"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		panic("Error connecting to the database")
	}

	db.AutoMigrate(&dbentities.Product{})

	return db
}
