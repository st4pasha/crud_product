package models

type Product struct {
	Id    int `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
	Count int `json:"count"`
}