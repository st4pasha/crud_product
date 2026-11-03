package repositories

import (
	"crud_product/app/models"
	"errors"
)

var products map[string]models.Product = map[string]models.Product{}

func GetProducts() map[string]models.Product {
	return products
}

func GetProductById(id string) (models.Product, error) {
	product, ok := products[id]
	if !ok {
		return models.Product{}, errors.New("product not found")
	}

	return product, nil
}

func UpdateProducts(productUpdated models.Product) error {
	_, err := GetProductById(productUpdated.Id)
	if err != nil {
		return err
	}
	products[productUpdated.Id] = productUpdated

	return nil
}

func DeleteProductById(id string) error {
	_, ok := GetProducts()[id]
	if !ok {
		return errors.New("product not found")
	}

	delete(GetProducts(), id)

	return nil
}
