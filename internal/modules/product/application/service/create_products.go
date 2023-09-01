package service

import (
	"golang-crud-exagonal-arquitecture/internal/modules/product/model/entity"
)

func (p *product) RegisterProduct(productEntity entity.Product) error {
	err := entity.ValidateProduct(productEntity)

	if err != nil {
		return err
	}

	logError := p.productRepository.RegisterProduct(productEntity)
	return logError
}
