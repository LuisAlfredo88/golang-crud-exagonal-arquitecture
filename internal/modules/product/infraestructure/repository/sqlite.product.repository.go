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

func (r *repo) SaveProduct(productEntity entity.Product) error {
	result := r.db.Omit("ID").Create(productEntity)
	return result.Error
}

func (r *repo) GetAllProducts() ([]entity.Product, error) {
	var products []entity.Product
	r.db.Find(&products)
	return products, nil
}

func (r *repo) DeleteProduct(productId int32) error {
	result := r.db.Delete(&entity.Product{}, productId)
	return result.Error
}
