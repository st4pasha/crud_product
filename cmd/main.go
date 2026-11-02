package main

import (
	"crud_product/app/handlers"
	"fmt"
	"net/http"
)

func main() {
	mux := handlers.RegisterHandlers()

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println("error -", err)
		return
	}
}
