package handlers

import (
	"crud_product/app/models"
	"crud_product/app/service"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handleGetProducts(w http.ResponseWriter, _ *http.Request) {
	products := service.GetProducts()

	productsBytes, err := json.Marshal(products)
	if err != nil {
		fmt.Printf("error, %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(productsBytes); err != nil {
		fmt.Printf("error, %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func handlerCreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		fmt.Printf("error, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := service.AddValidProductToRepositories(product); err != nil {
		fmt.Printf("error, %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

	if _, err := w.Write([]byte("Создан продукт с id " + product.Id + " и названием " + product.Title)); err != nil {
		fmt.Printf("error, %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func handleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)

	product_id, ok := path["product_id"]
	if !ok {
		fmt.Println("invalid id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := service.DeleteProduct(product_id)
	if err != nil {
		fmt.Printf("error, %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func handleUpdateProduct(w http.ResponseWriter, r *http.Request) {

	var productUpdate models.Product = models.Product{}

	if err := json.NewDecoder(r.Body).Decode(&productUpdate); err != nil {
		fmt.Printf("error, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := service.UpdateProduct(productUpdate); err != nil {
		fmt.Printf("error, %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, err := w.Write([]byte("Продукт успешно обновлен")); err != nil {
		fmt.Printf("error, %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func handleGetById(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)

	id, ok := path["id"]
	if !ok {
		fmt.Println("invalid id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := service.GetProductById(id)
	if err != nil {
		fmt.Printf("error, %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(product); err != nil {
		fmt.Printf("invalid id, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func StartServer() {

	mux := mux.NewRouter()

	mux.HandleFunc("/products/{id}", handleGetById).Methods("GET")
	mux.HandleFunc("/products", handleGetProducts).Methods("GET")
	mux.HandleFunc("/products", handlerCreateProduct).Methods("POST")
	mux.HandleFunc("/products/{product_id}", handleDeleteProduct).Methods("DELETE")
	mux.HandleFunc("/products", handleUpdateProduct).Methods("PUT")

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println("error -", err)
		return
	}
}
