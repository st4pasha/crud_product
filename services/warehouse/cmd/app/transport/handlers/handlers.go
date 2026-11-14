package handlers

import (
	"crud_product/services/warehouse/cmd/app/models"
	"crud_product/services/warehouse/cmd/app/service"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func CheckAvailableProductHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product = models.Product{}
	var available bool

	if err := service.CheckProductAvailability(product); err != nil {
		if err := writeJSON(w, 409, available); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		return
	}

	available = true

	if err := writeJSON(w, 200, available); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func ReserveProductHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product = models.Product{}
	var sucess bool

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := service.ReserveProduct(product); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	sucess = true

	if err := writeJSON(w, 200, sucess); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

func RegisterHanders() *chi.Mux {
	chi := chi.NewRouter()

	chi.Get("/products/check", CheckAvailableProductHandler)
	chi.Post("/products/reserve", ReserveProductHandler)
	
	return chi
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
