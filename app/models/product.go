package models

import (
	"errors"
	"math/rand"
)

type Product struct {
	Id    string
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

func ValidProduct(product Product) (Product, error) {
	if product.Title == "" {
		return Product{}, errors.New("invalid name")
	}

	if product.Price < 0 {
		return Product{}, errors.New("invalid price")
	}

	return Product{string(rand.Intn(100000)), product.Title, product.Price}, nil

}
