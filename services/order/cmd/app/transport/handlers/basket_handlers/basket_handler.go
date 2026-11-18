package basket

import (
	product "crud_product/services/order/cmd/app/repositories/product_repositories"
	basket "crud_product/services/order/cmd/app/service/basket_service.go"
	"crud_product/services/order/cmd/app/transport/dto"
	"crud_product/services/order/cmd/app/util"
	"encoding/json"
	"net/http"
)

func AddProductHandler(w http.ResponseWriter, r *http.Request) {
	var productToBasket dto.ProductToBasket = dto.ProductToBasket{}

	if err := json.NewDecoder(r.Body).Decode(&productToBasket); err != nil {
		http.Error(w, "Bad Request error", http.StatusBadRequest)
		return
	}

	if _, err := product.GetById(productToBasket.ProductID); err != nil {
		http.Error(w, "Bad Request error", http.StatusBadRequest)
		return
	}

	if err := basket.AddToRepositories(productToBasket.ProductID, productToBasket.Count); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := util.WriteJSON(w, 201, productToBasket); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	
}
