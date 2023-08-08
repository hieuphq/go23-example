package model

// Product is the model for a product.
type Product struct {
	Base
	Name        string
	Description string
	SKU         string
	Price       float64
	Currency    string
	Vendors     []Vendor `gorm:"many2many:product_vendors;"`
}

type Vendor struct {
	Base
	Name        string
	Description string
	Email       string
	Products    []Product `gorm:"many2many:product_vendors;"`
}
