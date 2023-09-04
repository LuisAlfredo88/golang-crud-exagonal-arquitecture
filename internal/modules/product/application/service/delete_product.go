package service

import (
	"errors"
)

func (p *product) DeleteProduct(productId int32) error {
	if productId <= 0 {
		return errors.New("must specify product id")
	}

	logError := p.productRepository.DeleteProduct(productId)
	return logError
}
