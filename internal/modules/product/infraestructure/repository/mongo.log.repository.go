package repository

import (
	"golang-crud-exagonal-arquitecture/internal/modules/product/model/entity"
	productModel "golang-crud-exagonal-arquitecture/internal/modules/product/model/repository"

	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewLogRepository(db *gorm.DB) productModel.ProductRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) RegisterProduct(productEntity entity.Product) error {
	result := r.db.Create(productEntity)
	return result.Error
}

func (r *repo) GetAllProducts() ([]entity.Product, error) {
	var products []entity.Product
	r.db.Find(&products)
	return products, nil
}
