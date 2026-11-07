package service

import (
	"crud_product/app/models"
	"crud_product/app/repositories"
	"crud_product/app/transport/dto"
	"database/sql"
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

	if err == sql.ErrNoRows {
		return models.Product{}, nil
	}

	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func UpdateProduct(productDTO dto.ProductToCreate, id string) (models.Product, error) {

	product, err := repositories.GetProductById(id)

	if err != nil {
		return models.Product{}, err
	}

	if productDTO.Title != "" {
		if productDTO.Price != 0 && productDTO.Price > 0 {
			newProduct := models.Product{Id: id, Title: productDTO.Title, Price: productDTO.Price}

			productToUpdate, err := repositories.UpdateProducts(newProduct)
			if err != nil {
				return models.Product{}, err
			}

			return productToUpdate, nil
		}

		newProduct := models.Product{Id: id, Title: productDTO.Title, Price: product.Price}

		productUpdated, err := repositories.UpdateProducts(newProduct)
		if err != nil {
			return models.Product{}, err
		}

		return productUpdated, nil

	} else if productDTO.Price != 0 && productDTO.Price > 0 {
		newProduct := models.Product{Id: id, Title: product.Title, Price: productDTO.Price}

		productUpdated, err := repositories.UpdateProducts(newProduct)
		if err != nil {
			return models.Product{}, err
		}

		return productUpdated, nil
	}

	return models.Product{}, errors.New("only ID received")
}
