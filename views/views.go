package views

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/markbates/goth/gothic"
	"github.com/zujko/mini-url/db"
	"github.com/zujko/mini-url/util"
)

var indexTemplate *template.Template
var footerTemplate *template.Template

// URLRsp is the format expected from the client.
type URLRsp struct {
	Data string `json:"url"`
}

// ShortURL is the format sent to the client.
type ShortURL struct {
	ShortUrl string `json:"shortUrl"`
	LongUrl  string `json:"longUrl"`
	Error    bool   `json:"error"`
}
type User struct {
	Name string
}

// Index Displays the index page.
func Index(w http.ResponseWriter, r *http.Request) {
	session, err := gothic.Store.Get(r, gothic.SessionName)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("USER", session.Values["name"])
	var username string
	if session.Values["name"] != nil {
		username = session.Values["name"].(string)
	}

	indexTemplate.Execute(w, &User{username})
}

// HandleUrl Handles redirecting a client to the associated page.
func HandleUrl(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.URL.Query().Get(":shorturl")
	fmt.Println("Handling URL", shortUrl)
	if shortUrl == "favicon.ico" {
		http.ServeFile(w, r, "static/favicon.ico")
		return
	}
	var longURL string
	err := db.DBConn.QueryRow("SELECT long_url FROM url WHERE short_url = $1", shortUrl).Scan(&longURL)
	if err != nil {
		log.Fatal("Failed to get longurl")
	}
	fmt.Println("redirecting")
	http.Redirect(w, r, longURL, http.StatusMovedPermanently)
}

// Shorten Handles grabbing the long url from the client, checking if it is a valid URL
// then calling ShortenURL to encode and store the URL.
func Shorten(w http.ResponseWriter, r *http.Request) {
	var url URLRsp
	json.NewDecoder(r.Body).Decode(&url)
	var resp = &ShortURL{"", "", true}
	// Check if the url is valid
	if !util.IsURL(url.Data) {
		resp = &ShortURL{"", "", true}
	} else {
		result := util.ShortenURL(url.Data)
		resp = &ShortURL{result, url.Data, false}
	}
	json.NewEncoder(w).Encode(resp)
}

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

// LoadTemplates Handles loading all HTML files.
func LoadTemplates() {
	indexTemplate = template.Must(template.ParseFiles("./templates/index.html"))
}
