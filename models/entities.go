package models

type Product struct {
	Title       string      `json:"title" validate:"required,max=10"`
	Description string      `json:"description" validate:"max=200"`
	Type        ProductType `json:"type" validate:"required"`
	Price       int         `json:"price" validate:"required"`
}

type ProductType string

const (
	item    ProductType = "item"
	service ProductType = "service"
)
