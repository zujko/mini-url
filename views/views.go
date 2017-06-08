package views

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/zujko/mini-url/util"
)

var indexTemplate *template.Template
var footerTemplate *template.Template

type URLRsp struct {
	Data string `json:"url"`
}

type ShortURL struct {
	Data string `json:"data"`
}

// Handles Index page
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	indexTemplate.Execute(w, nil)
}

func HandleUrl(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	shortUrl := ps.ByName("shorturl")
	fmt.Println(shortUrl)
}

func Shorten(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var url URLRsp
	json.NewDecoder(r.Body).Decode(&url)
	var resp = &ShortURL{""}
	// Get Short url
	if !util.IsURL(url.Data) {
		resp = &ShortURL{"invalid"}
	} else {
		result := util.ShortenURL(url.Data)
		resp = &ShortURL{result}
	}
	json.NewEncoder(w).Encode(resp)
}

func LoadTemplates() {
	indexTemplate = template.Must(template.ParseFiles("./templates/index.html"))
}
