package service

import (
	productRepo "golang-crud-exagonal-arquitecture/internal/modules/product/model/repository"
	productService "golang-crud-exagonal-arquitecture/internal/modules/product/model/service"
)

type product struct {
	productRepository productRepo.ProductRepository
}

func NewProductService(productR productRepo.ProductRepository) productService.ProductService {
	return &product{
		productRepository: productR,
	}
}
