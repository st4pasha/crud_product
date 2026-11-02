package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"crud_product/app/models"
	"crud_product/app/service"
)

func handleGetProducts(w http.ResponseWriter, _ *http.Request) {
	products := service.GetProducts()

	if err := writeJSON(w, 200, products); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func handlerCreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	validProduct, err := service.CreateProduct(product.Id, product.Title, product.Price)
	if err != nil {
		http.Error(w, "Bad Request Error", http.StatusBadRequest)
		return
	}

	service.AddValidProductToRepositories(validProduct)

	w.WriteHeader(http.StatusCreated)

	if err := writeJSON(w, 200, "Создан продукт с id "+product.Id+" и названием "+product.Title); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func handleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)

	product_id, ok := path["product_id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := service.DeleteProduct(product_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func handleUpdateProduct(w http.ResponseWriter, r *http.Request) {

	var productUpdate models.Product = models.Product{}

	if err := json.NewDecoder(r.Body).Decode(&productUpdate); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := service.UpdateProduct(productUpdate); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := writeJSON(w, 200, "Продукт успешно обновлен"); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

func handleGetById(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)

	id, ok := path["product_id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := service.GetProductById(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := writeJSON(w, 200, product); err != nil {
		fmt.Printf("invalid id, %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func RegisterHandlers() *mux.Router {
	mux := mux.NewRouter()
	
	mux.HandleFunc("/products/{product_id}", handleGetById).Methods("GET")
	mux.HandleFunc("/products", handleGetProducts).Methods("GET")
	mux.HandleFunc("/products", handlerCreateProduct).Methods("POST")
	mux.HandleFunc("/products/{product_id}", handleDeleteProduct).Methods("DELETE")
	mux.HandleFunc("/products", handleUpdateProduct).Methods("PUT")

	return mux
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
