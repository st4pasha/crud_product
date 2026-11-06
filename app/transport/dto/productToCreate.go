package dto

type ProductToCreate struct {
	Title string `json:"title"`
	Price float64 `json:"price"`
}