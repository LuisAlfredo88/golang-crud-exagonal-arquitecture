package sqlite

import (
	productEntity "golang-crud-exagonal-arquitecture/internal/modules/product/model/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		panic("Error connecting to the database")
	}

	db.AutoMigrate(&productEntity.Product{})

	return db
}
