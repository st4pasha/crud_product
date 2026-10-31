package models

type ProductUpdate struct {
	Id    string   `json:"id"`
	Title *string  `json:"title"`
	Price *float64 `json:"price"`
}
