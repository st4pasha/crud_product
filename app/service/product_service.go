package service

import (
	"crud_product/app/models"
	"crud_product/app/repositories"
	"errors"
)

func ValidProduct(product models.Product) (models.Product, error) {
	if product.Title == "" {
		return models.Product{}, errors.New("invalid name")
	}

	if product.Price <= 0 {
		return models.Product{}, errors.New("invalid price")
	}

	if product.Id == "" {
		return models.Product{}, errors.New("invalid id")
	}

	return models.Product{Id: product.Id, Title: product.Title, Price: product.Price}, nil

}

func CreateProduct(id string, title string, price float64) models.Product {
	return models.Product{Id: id, Title: title, Price: price}
}

func GetProducts() map[string]models.Product {
	return repositories.Products
}

func AddValidProductToRepositories(product models.Product) error {
	validProduct, err := ValidProduct(product)
	if err != nil {
		return err
	}

	repositories.Products[validProduct.Id] = validProduct

	return nil
}
func DeleteProduct(id string) (models.Product, error) {
	product, ok := repositories.Products[id]
	if !ok {
		return models.Product{}, errors.New("product not found")
	}

	delete(repositories.Products, id)

	return product, nil
}

func GetProductById(id string) (models.Product, error) {
	product, ok := repositories.Products[id]
	if !ok {
		return models.Product{}, errors.New("product not found")
	}

	return product, nil
}

func UpdateProduct(productToUpdate models.Product) error {

	product, ok := repositories.Products[productToUpdate.Id]
	if !ok {
		return errors.New("product not found")
	}
	if productToUpdate.Title != "" {
		if productToUpdate.Price != 0 && productToUpdate.Price > 0 {
			newProduct := CreateProduct(productToUpdate.Id, productToUpdate.Title, productToUpdate.Price)

			repositories.Products[productToUpdate.Id] = newProduct

			return nil
		}

		newProduct := CreateProduct(productToUpdate.Id, productToUpdate.Title, product.Price)

		repositories.Products[productToUpdate.Id] = newProduct

		return nil
	} else if productToUpdate.Price != 0 && productToUpdate.Price > 0 {
		newProduct := CreateProduct(productToUpdate.Id, product.Title, productToUpdate.Price)

		repositories.Products[productToUpdate.Id] = newProduct

		return nil
	}

	return errors.New("only ID received")
}
