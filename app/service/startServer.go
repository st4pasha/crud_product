package service

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Start_server() {

	mux := mux.NewRouter()

	mux.HandleFunc("/products", handleGetProducts).Methods("GET")
	mux.HandleFunc("/products", handlerCreateProduct).Methods("POST")
	mux.HandleFunc("/products/{product_id}", handleDeleteProduct).Methods("DELETE")
	mux.HandleFunc("/products", handleUpdateProduct).Methods("PUT")

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println("error -", err)
		return
	}
}
