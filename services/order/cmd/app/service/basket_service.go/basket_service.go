package basket

import basket "crud_product/services/order/cmd/app/repositories/basket_repositories.go"

func AddToRepositories(productID, productCount int) error {
	if err := basket.AddToRepositories(productID, productCount); err != nil {
		return err
	}

	return nil
}
