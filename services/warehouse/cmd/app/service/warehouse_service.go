package service

import (
	"crud_product/services/warehouse/cmd/app/models"
	"crud_product/services/warehouse/cmd/app/repositories"
)

func CheckProductAvailability(product models.Product) error {
	if err := repositories.CheckCountProduct(product); err != nil {
		return err
	}

	return nil
}

func ReserveProduct(product models.Product) error {

	if err := repositories.CheckCountProduct(product); err != nil {
		return err
	}

	if err := repositories.UpdateCountProduct(product); err != nil {
		return err
	}

	return nil
}