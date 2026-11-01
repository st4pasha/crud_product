package service

import (
	"crud_product/app/models"
	"errors"
)

func ValidProduct(product models.Product) (models.Product, error) {
	if *product.Title == "" {
		return models.Product{}, errors.New("invalid name")
	}

	if *product.Price < 0 {
		return models.Product{}, errors.New("invalid price")
	}

	if product.Id == "" {
		return models.Product{}, errors.New("invalid id")
	}

	return models.Product{product.Id, product.Title, product.Price}, nil

}
