package basket

import basket "crud_product/services/order/cmd/internal/basket/repositories"

func AddToRepositories(productID, productCount int) error {
	if err := basket.AddToRepositories(productID, productCount); err != nil {
		return err
	}

	return nil
}
