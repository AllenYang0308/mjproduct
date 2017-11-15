package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/leekchan/gtf"
	"net/http"

	"./models"
	"html/template"
	"io/ioutil"
	"log"
	"net/url"
	"regexp"
)

//type Product struct {
//	gorm.Model
//	ProductName        string
//	ProductPrice       string
//	ProductDescription string
//	ProductMid         string
//	ProductCode        string
//}
//
//type Products struct {
//	Prod []Product
//}

func FindString(bts [][]byte, srange int) []string {
	var rts []string
	for i := srange; i < len(bts); i++ {
		rts = append(rts, string(bts[i]))
	}
	return rts
}

func show(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect database.")
	}

	var Prod models.Products
	var prod []models.Product
	db.Find(&prod)
	Prod.Prod = prod
	t, _ := template.ParseFiles("./templates/index.html")
	t = t.Funcs(gtf.GtfFuncMap)

	err = t.Execute(w, Prod)
	if err != nil {
		panic(err)
	}

}

func search(w http.ResponseWriter, r *http.Request) {

	// NOTE: For test gorm data store.
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect database.")
	}
	defer db.Close()

	db.AutoMigrate(&models.Product{})

	var product models.Product
	var apiparam = url.Values{}
	var apiurl string
	r.ParseForm()
	for k, v := range r.Form {
		for _, sv := range v {
			apiparam.Add(k, sv)
		}
	}

	fmt.Println("apiparam: ", apiparam)
	apiurl = "http://www.muji.tw/item_detail.aspx?CatID=1&PdtID=2&" + apiparam.Encode()
	fmt.Println("apiurl: ", apiurl)
	res, err := http.Get(apiurl)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	bts, _ := ioutil.ReadAll(res.Body)
	pattern := make(map[string]string)

	pattern["prod_name"] = `<li\sclass="[a-z]+">\s+<h3>(.*?)</h3>\s*</li>`
	pattern["prod_desc"] = `<li\sclass="[a-z]+">\s+<h3>.*?</h3>\s*</li>\s*<li>(.*?)</li>`
	pattern["prod_price"] = `<ol>\s*<li\sclass="[a-z]+">.*\p{Han}+.*>(\d+).*\s*<li.*>\p{Han}+.*>\d+</li>`
	pattern["prod_code"] = `<ol>\s*<li\sclass="[a-z]+">.*\p{Han}+.*>\d+.*\s*<li.*>\p{Han}+.*>(\d+)</li>`
	pattern["prod_mid"] = `<p\sclass="num">\p{Han}+&nbsp;*(.*?)</p>`

	prod_name_pattern := regexp.MustCompile(pattern["prod_name"])
	prod_name := prod_name_pattern.FindSubmatch(bts)

	prod_desc_pattern := regexp.MustCompile(pattern["prod_desc"])
	prod_desc := prod_desc_pattern.FindSubmatch(bts)

	prod_price_pattern := regexp.MustCompile(pattern["prod_price"])
	prod_price := prod_price_pattern.FindSubmatch(bts)

	prod_code_pattern := regexp.MustCompile(pattern["prod_code"])
	prod_code := prod_code_pattern.FindSubmatch(bts)

	prod_mid_pattern := regexp.MustCompile(pattern["prod_mid"])
	prod_mid := prod_mid_pattern.FindSubmatch(bts)

	if len(prod_name) > 0 {
		product.ProductName = string(prod_name[1])
	}
	if len(prod_price) > 0 {
		product.ProductPrice = string(prod_price[1])
	}
	if len(prod_code) > 0 {
		product.ProductCode = string(prod_code[1])
	}
	if len(prod_mid) > 0 {
		product.ProductMid = string(prod_mid[1])
	}
	if len(prod_desc) > 0 {
		product.ProductDescription = string(prod_desc[1])
	}

	if (product.ProductName != "") && (product.ProductPrice != "") && (product.ProductCode != "") {
		db.Create(&product)
	}

	http.Redirect(w, r, "/show", http.StatusFound)
	return

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", show)
	r.HandleFunc("/search", search)
	r.HandleFunc("/show", show)
	log.Fatal(http.ListenAndServe(":8080", r))
}
