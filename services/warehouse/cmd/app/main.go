package main

import (
	"crud_product/services/warehouse/internal/warehouse/handlers"
	"net/http"
)

func main(){
	router := handlers.RegisterHanders()

	if err := http.ListenAndServe(":8081", router); err != nil {
		return
	}
}