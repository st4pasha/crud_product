package service

import (
	"crud_product/app/models"
	"crud_product/app/repositories"
	"errors"
)

func CreateProduct(title string, price float64) (models.Product, error) {

	product, err := repositories.AddToRepositories(title, price)

	if err != nil {
		return models.Product{}, err
	}

	return product, nil

}

func GetProducts() ([]models.Product, error) {
	return repositories.GetProducts()
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

func UpdateProduct(productToUpdate models.Product) (models.Product, error) {

	product, err := repositories.GetProductById(productToUpdate.Id)

	if err != nil {
		return models.Product{}, err
	}

	if productToUpdate.Title != "" {
		if productToUpdate.Price != 0 && productToUpdate.Price > 0 {
			newProduct := models.Product{Id: productToUpdate.Id, Title: productToUpdate.Title, Price: productToUpdate.Price}

			productUpdated, err := repositories.UpdateProducts(newProduct); 
			if err != nil {
				return models.Product{}, err
			}

			return productUpdated, nil
		}

		newProduct := models.Product{Id: productToUpdate.Id, Title: productToUpdate.Title, Price: product.Price}

		productUpdated, err := repositories.UpdateProducts(newProduct); 
			if err != nil {
				return models.Product{}, err
			}

			return productUpdated, nil

	} else if productToUpdate.Price != 0 && productToUpdate.Price > 0 {
		newProduct := models.Product{Id: productToUpdate.Id, Title: product.Title, Price: productToUpdate.Price}

		productUpdated, err := repositories.UpdateProducts(newProduct); 
			if err != nil {
				return models.Product{}, err
			}

			return productUpdated, nil
	}

	return models.Product{}, errors.New("only ID received")
}
