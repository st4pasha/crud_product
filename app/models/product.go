package models

import "errors"

type Product struct {
	Id    string  `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

func (p Product) ValidProduct() error {
	if p.Title == "" {
		return errors.New("invalid name")
	}

	if p.Price <= 0 {
		return errors.New("invalid price")
	}

	if p.Id == "" {
		return errors.New("invalid id")
	}

	return nil

}
