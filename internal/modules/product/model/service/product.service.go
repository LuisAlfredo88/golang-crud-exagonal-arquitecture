package service

import (
	"golang-crud-exagonal-arquitecture/internal/modules/product/model/entity"
)

type ProductService interface {
	RegisterProduct(product entity.Product) error
	GetAllProducts() ([]entity.Product, error)
	DeleteProduct(productId int32) error
	UpdateProduct(product entity.Product) error
}
