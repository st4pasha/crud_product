package product

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	productservice "crud_product/app/service/product_service"
	"crud_product/app/transport/dto"
)

func handleGet(w http.ResponseWriter, _ *http.Request) {
	products, err := productservice.Get()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := writeJSON(w, 200, products); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func handlerCreate(w http.ResponseWriter, r *http.Request) {
	var product dto.ProductToCreate

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	validProduct, err := productservice.Create(product.Title, product.Price)
	if err != nil {
		http.Error(w, "Bad Request Error", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

	if err := writeJSON(w, 200, validProduct); err != nil {
		return
	}
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	product_id := chi.URLParam(r, "product_id")

	if product_id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := productservice.Delete(product_id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func handleUpdate(w http.ResponseWriter, r *http.Request) {

	var productDTO dto.ProductToCreate = dto.ProductToCreate{}

	product_id := chi.URLParam(r, "product_id")
	if product_id == "" {
		http.Error(w, "Bad Request Error", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&productDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	productUpdated, err := productservice.Update(productDTO, product_id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := writeJSON(w, 200, productUpdated); err != nil {
		return
	}

}

func handleGetById(w http.ResponseWriter, r *http.Request) {
	product_id := chi.URLParam(r, "product_id")

	if product_id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	productservice, err := productservice.GetById(product_id)

	if err != nil {
		
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "{'error': productservice not found}", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := writeJSON(w, 200, productservice); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func RegisterHandlers() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/products/{product_id}", handleGetById)
	router.Get("/products", handleGet)
	router.Post("/products", handlerCreate)
	router.Delete("/products/{product_id}", handleDelete)
	router.Patch("/products/{product_id}", handleUpdate)

	return router
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
