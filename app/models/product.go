package models

import (
	"errors"
)

type Product struct {
	Id    string  `json:"id"`
	Title *string  `json:"title"`
	Price *float64 `json:"price"`
}

func ValidProduct(product Product) (Product, error) {
	if *product.Title == "" {
		return Product{}, errors.New("invalid name")
	}

	if *product.Price < 0 {
		return Product{}, errors.New("invalid price")
	}

	if product.Id == "" {
		return Product{}, errors.New("invalid id")
	}

	return Product{product.Id, product.Title, product.Price}, nil

}
