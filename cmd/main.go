package main

import (
	handlers "crud_product/app/transport/handlers/product_handler"
	"fmt"
	"net/http"
)

func main() {
	router := handlers.RegisterHandlers()

	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		fmt.Println("error -", err)
		return
	}

}
