package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
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

	if err := writeJSON(w, 200, validProduct); err != nil {
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

	if err := writeJSON(w, 200, productUpdate); err != nil {
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
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func RegisterHandlers() *chi.Mux {
	mux := chi.NewRouter()

	mux.Get("/products/{product_id}", handleGetById)
	mux.Get("/products", handleGetProducts)
	mux.Post("/products", handlerCreateProduct)
	mux.Delete("/products/{product_id}", handleDeleteProduct)
	mux.Put("/products", handleUpdateProduct)

	return mux
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
