package repositories

import (
	"crud_product/app/models"
	"errors"
)

var Products map[string]models.Product = map[string]models.Product{}

func GetProducts() map[string]models.Product {
	return Products
}

func GetProductById(id string) (models.Product, error) {
	product, ok := Products[id]
	if !ok {
		return models.Product{}, errors.New("product not found")
	}

	return product, nil
}

func UpdateProducts(productUpdated models.Product) {
	Products[productUpdated.Id] = productUpdated
}
