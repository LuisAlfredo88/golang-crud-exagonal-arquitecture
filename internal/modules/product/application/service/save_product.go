package service

import (
	"golang-crud-exagonal-arquitecture/internal/modules/product/model/entity"
)

func (p *product) SaveProduct(productEntity entity.Product) error {
	err := error(nil)

	if productEntity.ID > 0 {
		err = entity.ValidateProductUpdate(productEntity)
	} else {
		err = entity.ValidateProduct(productEntity)
	}

	if err != nil {
		return err
	}

	logError := p.productRepository.SaveProduct(productEntity)
	return logError
}
