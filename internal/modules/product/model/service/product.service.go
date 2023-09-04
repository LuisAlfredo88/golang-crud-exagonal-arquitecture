package service

import (
	"golang-crud-exagonal-arquitecture/internal/modules/product/model/entity"
)

type ProductService interface {
	SaveProduct(product entity.Product) error
	GetAllProducts() ([]entity.Product, error)
	DeleteProduct(productId int32) error
}
