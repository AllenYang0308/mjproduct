package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Product object.
type Product struct {
	gorm.Model
	ProductName        string
	ProductPrice       string
	ProductDescription string
	ProductMid         string
	ProductCode        string
}

// For query
type Products struct {
	Prod []Product
}
