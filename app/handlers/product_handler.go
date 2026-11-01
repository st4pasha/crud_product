package handlers

import (
	"crud_product/app/service"
	"crud_product/app/models"
	"crud_product/app/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func handleGetProducts(w http.ResponseWriter, _ *http.Request) {
	productsBytes, err := json.Marshal(repositories.Products)
	if err != nil {
		fmt.Printf("error, %v\n", err)
		return
	}
	if _, err = w.Write(productsBytes); err != nil {
		fmt.Printf("error, %v\n", err)
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

	validatedProduct, err := service.ValidProduct(product)
	if err != nil {
		fmt.Printf("error, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	repositories.Products[validatedProduct.Id] = &validatedProduct

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("Создан новый продукт с названием: , " + *product.Title))
	if err != nil {
		fmt.Printf("error, %v\n", err)
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

	for key, product := range repositories.Products {
		if product.Id == product_id {
			delete(repositories.Products, key)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	w.WriteHeader(http.StatusBadRequest)
}

func handleUpdateProduct(w http.ResponseWriter, r *http.Request) {

	var productUpdate models.Product = models.Product{}

	if err := json.NewDecoder(r.Body).Decode(&productUpdate); err != nil {
		fmt.Printf("error, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for key, product := range repositories.Products {
		if product.Id == productUpdate.Id {
			if productUpdate.Price != nil && *product.Price > 0 {
				repositories.Products[key].Price = productUpdate.Price
			}

			if productUpdate.Title != nil {
				repositories.Products[key].Title = productUpdate.Title
			}

			w.WriteHeader(http.StatusOK)
			if _, err := w.Write([]byte("Продукт успешно изменён")); err != nil {
				fmt.Printf("invalid id, %v\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			return

		}
	}

	w.WriteHeader(http.StatusBadRequest)
	if _, err := w.Write([]byte("Продукт не найден")); err != nil {
		fmt.Printf("invalid id, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func handleGetById(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)

	id := path["id"]

	product, ok := repositories.Products[id]
	if !ok {
		fmt.Println("invalid id")
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
