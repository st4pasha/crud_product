package main

import (
	"crud_product/app/handlers"
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
