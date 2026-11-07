package main

import (
	"crud_product/app/repositories"
	"crud_product/app/transport/handlers"
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
