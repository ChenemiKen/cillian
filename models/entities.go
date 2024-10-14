package models

type Product struct {
	Title       string `json:"title" validate:"required,max=10"`
	Description string `json:"description" validate:"max=200"`
	Type        string `json:"type" validate:"required"`
	Price       int    `json:"price" validate:"required"`
}
