package repository

import (
	"golang-crud-exagonal-arquitecture/internal/modules/product/model/entity"
)

type ProductRepository interface {
	RegisterProduct(product entity.Product) error
	GetAllProducts() ([]entity.Product, error)
}
