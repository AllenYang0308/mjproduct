package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"os"
	"text/template"
)

var (
	app *mux.Router
)

type Product struct {
	ProductName        string
	ProductPrice       string
	ProductDescription string
	ProductMid         string
	ProductCode        string
}

func FindString(bts [][]byte, srange int) []string {
	var rts []string
	for i := srange; i < len(bts); i++ {
		rts = append(rts, string(bts[i]))
	}
	return rts
}

// *mux.Router
func intercept(w http.ResponseWriter, r *http.Request) {
	var app *mux.Router
	app.ServeHTTP(w, r)
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

// Views func handler.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.New("Product List")
	tpl, _ = tpl.Parse("Product name: {{{.ProductName}}")
	p := Product{ProductName: "GO parse"}
	tpl.Execute(os.Stdout, p)
}

//func main() {
//	app = mux.NewRouter()
//	http.HandleFunc("/", intercept)
//	app.HandleFunc("/home", HomeHandler)
//
//	http.ListenAndServe(":9090", nil)
//}

func main() {

	var apiparam = url.Values{}
	var apiurl string
	apiparam.Set("CodeID", "4550002056247")

	apiurl = "http://www.muji.tw/item_detail.aspx?CatID=1&PdtID=2&" + apiparam.Encode()
	fmt.Println(apiurl)
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

	fmt.Println(FindString(prod_name, 1))
	fmt.Println(FindString(prod_desc, 1))
	fmt.Println(FindString(prod_price, 1))
	fmt.Println(FindString(prod_code, 1))
	fmt.Println(FindString(prod_mid, 1))

}
