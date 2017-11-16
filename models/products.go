package models

import (
	"bytes"
	"encoding/gob"
	//	"fmt"
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

type ProductModel struct {
	ProductName        interface{}
	ProductPrice       string
	ProductDescription interface{}
	ProductMid         string
	ProductCode        string
}

// For query
type Products struct {
	Prod []ProductModel
}

func GetBytes(key interface{}) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(key)
	return buf.Bytes()
}
