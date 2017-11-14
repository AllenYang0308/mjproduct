package main

import (
	"./product"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"time"
)

func main() {
	db, err := gorm.Open("postgres", "user=allenyang password=musasi9936 dbname=mjproduct sslmode=disable")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	db.DropTableIfExists(&Owner{}, &Book{}, &Author{})
	db.CreateTable(&Owner{}, &Book{}, &Author{})

	var p product.Product
	p.Name = "product name"
	fmt.Println(p)
}

type Owner struct {
	gorm.Model
	FirstName string
	LastName  string
	Books     []Book
}

type Book struct {
	gorm.Model
	Name        string
	PublishDate time.Time
	OwnerID     uint     `sql:"index"`
	Authors     []Author `gorm:"many2many:books_authors"`
}

type Author struct {
	gorm.Model
	FirstName string
	LastName  string
}
