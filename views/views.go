package views

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var indexTemplate *template.Template
var footerTemplate *template.Template

// Handles Index page
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "GET" {
		indexTemplate.Execute(w, nil)
	}
}

func HandleUrl(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == "GET" {
		shortUrl := ps.ByName("shorturl")
		fmt.Println(shortUrl)
	}
}

func LoadTemplates() {
	indexTemplate = template.Must(template.ParseFiles("./templates/index.html"))
}
