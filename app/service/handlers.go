package service

import (
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

	validated_product, err := models.ValidProduct(product)
	if err != nil {
		fmt.Printf("error, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	repositories.Products[validated_product.Title] = &validated_product

	_, err = w.Write([]byte("Создан новый продукт с названием: , " + product.Title))
	if err != nil {
		fmt.Printf("error, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func handleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	mapPath := mux.Vars(r)

	product_id, ok := mapPath["product_id"]
	if !ok {
		fmt.Println("invalid id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for key, product := range repositories.Products {
		if product.Id == product_id {
			delete(repositories.Products, key)
			w.WriteHeader(http.StatusNoContent)
		}
	}
}

func handleUpdateProduct(w http.ResponseWriter, r *http.Request) {

	var productUpdate models.ProductUpdate = models.ProductUpdate{}

	if err := json.NewDecoder(r.Body).Decode(&productUpdate); err != nil {
		fmt.Printf("error, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for key, product := range repositories.Products {
		if product.Id == productUpdate.Id {
			if productUpdate.Price != nil {
				repositories.Products[key].Price = *productUpdate.Price
			}

			if productUpdate.Title != nil {
				repositories.Products[key].Title = *productUpdate.Title
			}

			w.WriteHeader(http.StatusOK)
			if _, err := w.Write([]byte("Продукт успешно изменён")); err != nil {
				fmt.Printf("invalid id, %v\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			return

		}

		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte("Продукт не найден")); err != nil {
			fmt.Printf("invalid id, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	}

}
