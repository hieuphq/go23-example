package model

// Product is the model for a product.
type Product struct {
	Base
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	SKU         string   `json:"sku,omitempty"`
	Price       float64  `json:"price,omitempty"`
	Currency    string   `json:"currency,omitempty"`
	Vendors     []Vendor `gorm:"many2many:product_vendors;" json:"vendors,omitempty"`
}

type CreateProductRequest struct {
	Name        string  `json:"name,omitempty" binding:"required"`
	Description string  `json:"description,omitempty"`
	SKU         string  `json:"sku,omitempty" binding:"required"`
	Price       float64 `json:"price,omitempty" binding:"required"`
	Currency    string  `json:"currency,omitempty"`
}

type Vendor struct {
	Base
	Name        string
	Description string
	Email       string
	Products    []Product `gorm:"many2many:product_vendors;"`
}
