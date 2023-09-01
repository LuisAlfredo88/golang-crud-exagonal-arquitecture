package service

import (
	"golang-crud-exagonal-arquitecture/internal/modules/product/model/entity"
)

func (l *product) GetAllProducts() ([]entity.Product, error) {
	return l.productRepository.GetAllProducts()
}
