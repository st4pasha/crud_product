package main

import (
	repositories "crud_product/app/repositories/product_repositories"
	handlers "crud_product/app/transport/handlers/product_handler"
	"fmt"
	"net/http"
)

func main() {

	if err := repositories.Init(); err != nil {
		fmt.Println("error, ", err)
		return
	}
	router := handlers.RegisterHandlers()

	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		fmt.Println("error -", err)
		return
	}

}
