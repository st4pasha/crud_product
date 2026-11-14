package product

import (
	"crud_product/services/order/cmd/app/models"
	productrepo "crud_product/services/order/cmd/app/repositories/product_repositories"
	"crud_product/services/order/cmd/app/transport/dto"
	"database/sql"
	"errors"
)

func Create(title string, price float64) (models.Product, error) {

	product, err := productrepo.AddToRepositories(title, price)

	if err != nil {
		return models.Product{}, err
	}

	return product, nil

}

func Get() ([]models.Product, error) {
	return productrepo.Get()
}

func Delete(id int) error {

	if err := productrepo.DeleteById(id); err != nil {
		return err
	}

	return nil
}

func GetById(id int) (models.Product, error) {
	product, err := productrepo.GetById(id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Product{}, nil
		}
		return models.Product{}, err
	}

	return product, nil
}

func Update(productDTO dto.ProductToCreate, id int) (models.Product, error) {

	product, err := productrepo.GetById(id)

	if err != nil {
		return models.Product{}, err
	}

	if productDTO.Title != "" {
		if productDTO.Price != 0 && productDTO.Price > 0 {
			newProduct := models.Product{Id: id, Title: productDTO.Title, Price: productDTO.Price}

			productToUpdate, err := productrepo.Update(newProduct)
			if err != nil {
				return models.Product{}, err
			}

			return productToUpdate, nil
		}

		newProduct := models.Product{Id: id, Title: productDTO.Title, Price: product.Price}

		productUpdated, err := productrepo.Update(newProduct)
		if err != nil {
			return models.Product{}, err
		}

		return productUpdated, nil

	} else if productDTO.Price != 0 && productDTO.Price > 0 {
		newProduct := models.Product{Id: id, Title: product.Title, Price: productDTO.Price}

		productUpdated, err := productrepo.Update(newProduct)
		if err != nil {
			return models.Product{}, err
		}

		return productUpdated, nil
	}

	return models.Product{}, errors.New("only ID received")
}
