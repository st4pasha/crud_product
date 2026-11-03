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

func CreateProduct(id string, title string, price float64) (models.Product, error) {
	product := models.Product{Id: id, Title: title, Price: price}

	validProduct, err := ValidProduct(product)
	if err != nil {
		return models.Product{}, err
	}

	return validProduct, nil
}

func GetProducts() map[string]models.Product {
	return repositories.GetProducts()
}

func AddValidProductToRepositories(validProduct models.Product) {
	repositories.UpdateProducts(validProduct)

}
func DeleteProduct(id string) error {

	if err := repositories.DeleteProductById(id); err != nil {
		return err
	}

	return nil
}

func GetProductById(id string) (models.Product, error) {
	product, err := repositories.GetProductById(id)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func UpdateProduct(productToUpdate models.Product) error {

	product, err := repositories.GetProductById(productToUpdate.Id)

	if err != nil {
		return err
	}

	if productToUpdate.Title != "" {
		if productToUpdate.Price != 0 && productToUpdate.Price > 0 {
			newProduct := models.Product{Id: productToUpdate.Id, Title: productToUpdate.Title, Price: productToUpdate.Price}

			if err := repositories.UpdateProducts(newProduct); err != nil {
				return err
			}

			return nil
		}

		newProduct := models.Product{Id: productToUpdate.Id, Title: productToUpdate.Title, Price: product.Price}

		repositories.UpdateProducts(newProduct)

		return nil
	} else if productToUpdate.Price != 0 && productToUpdate.Price > 0 {
		newProduct := models.Product{Id: productToUpdate.Id, Title: product.Title, Price: productToUpdate.Price}

		repositories.UpdateProducts(newProduct)

		return nil
	}

	return errors.New("only ID received")
}
