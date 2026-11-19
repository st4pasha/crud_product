package main

import (
	productHandlers "crud_product/services/order/cmd/internal/product/handlers"
	"fmt"
	"net/http"
)

func main() {
	router := productHandlers.RegisterHandlers()

	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		fmt.Println("error -", err)
		return
	}

}
