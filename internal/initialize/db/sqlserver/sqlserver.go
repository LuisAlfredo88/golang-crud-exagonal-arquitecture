package sqlserver

import (
	"fmt"
	"golang-crud-exagonal-arquitecture/internal/modules/vehicle/model/entity"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func NewSqlServer() *gorm.DB {
	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"sqlserver://%s:%s@%s:1433?database=%s",
		DB_USER,
		DB_PASSWORD,
		DB_HOST,
		DB_NAME,
	)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error connecting to the database")
	}

	db.AutoMigrate(&entity.Mark{})
	db.AutoMigrate(&entity.VehicleModel{})
	db.AutoMigrate(&entity.Vehicle{})

	return db
}
