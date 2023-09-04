package repository

import (
	"golang-crud-exagonal-arquitecture/internal/modules/product/model/entity"
)

type ProductRepository interface {
	SaveProduct(product entity.Product) error
	GetAllProducts() ([]entity.Product, error)
	DeleteProduct(productId int32) error
}
