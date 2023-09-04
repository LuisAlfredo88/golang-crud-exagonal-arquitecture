package service

import (
	"golang-crud-exagonal-arquitecture/internal/modules/product/model/entity"
)

func (p *product) UpdateProduct(productEntity entity.Product) error {
	err := entity.ValidateProductUpdate(productEntity)

	if err != nil {
		return err
	}

	logError := p.productRepository.SaveProduct(productEntity)
	return logError
}
