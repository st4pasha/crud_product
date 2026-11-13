package models

import "errors"

type Product struct {
	Id    int `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

func (p Product) Valid() error {
	if p.Title == "" {
		return errors.New("invalid name")
	}

	if p.Price <= 0 {
		return errors.New("invalid price")
	}

	return nil

}
