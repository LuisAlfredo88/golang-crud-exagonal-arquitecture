package entity

import (
	"errors"
	"strings"
)

func ValidateProduct(product Product) error {
	if strings.TrimSpace(product.Name) == "" {
		return errors.New("must specify product name")
	}

	if strings.TrimSpace(product.Description) == "" {
		return errors.New("must specify product description")
	}

	if product.Price <= 0 {
		return errors.New("must specify product price")
	}

	return nil
}

func ValidateProductUpdate(product Product) error {
	err := ValidateProduct(product)

	if err != nil {
		return err
	}

	// if product.ID <= 0 {
	// 	return errors.New("must specify product id")
	// }

	return nil
}
