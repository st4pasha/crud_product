package basket

import (
	"crud_product/services/order/cmd/internal/basket/dto"
	basket "crud_product/services/order/cmd/internal/basket/service"
	product "crud_product/services/order/cmd/internal/product/repositories"
	"crud_product/services/order/cmd/internal/util"
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