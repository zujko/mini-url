package views

import (
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

func LoadTemplates() {
	indexTemplate = template.Must(template.ParseFiles("./templates/index.html"))
}
