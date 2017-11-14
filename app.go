package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"

	"html/template"
	"log"
)

type Product struct {
	ProductName        string
	ProductPrice       string
	ProductDescription string
	ProductMid         string
	ProductCode        string
}

// NOTE: For gorm object query results.
type Products struct {
	Prod []Product
}

func FindString(bts [][]byte, srange int) []string {
	var rts []string
	for i := srange; i < len(bts); i++ {
		rts = append(rts, string(bts[i]))
	}
	return rts
}

func thandler(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("./templates/index.html")

	// NOTE: For template test.
	data := Products{
		Prod: []Product{
			{ProductName: "p1", ProductPrice: "100", ProductMid: "001", ProductCode: "A001"},
			{ProductName: "p1", ProductPrice: "100", ProductMid: "001", ProductCode: "A001"},
			{ProductName: "p1", ProductPrice: "100", ProductMid: "001", ProductCode: "A001"},
			{ProductName: "p1", ProductPrice: "100", ProductMid: "001", ProductCode: "A001"},
			{ProductName: "p1", ProductPrice: "100", ProductMid: "001", ProductCode: "A001"},
		},
	}
	fmt.Println(data)

	err := t.Execute(w, data)
	if err != nil {
		panic(err)
	}

}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:musasi.yang@gmail.com\">Allen Yang</a>")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/thandle", thandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
