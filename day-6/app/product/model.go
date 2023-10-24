package product

import "errors"

var (
	ErrEmptyName     = errors.New("name can't be empty")
	ErrEmptyCategory = errors.New("category can't be empty")
	ErrEmptyPrice    = errors.New("price can't be empty")
	ErrEmptyStock    = errors.New("stock can't be empty")
)

type Product struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Category string `json:"category" db:"category"`
	Price    int    `json:"price" db:"price"`
	Stock    int    `json:"stock" db:"stock"`
}

func (p Product) Validate() (err error) {
	if p.Name == "" {
		return ErrEmptyName
	}
	if p.Category == "" {
		return ErrEmptyCategory
	}
	if p.Price == 0 {
		return ErrEmptyPrice
	}
	if p.Stock == 0 {
		return ErrEmptyStock
	}
	return err
}
