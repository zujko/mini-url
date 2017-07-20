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

// Index Displays the index page.
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	indexTemplate.Execute(w, nil)
}

// HandleUrl Handles redirecting a client to the associated page.
func HandleUrl(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	shortUrl := ps.ByName("shorturl")
	fmt.Println("Handling URL", shortUrl)
	if shortUrl == "favicon.ico" {
		http.ServeFile(w, r, "static/favicon.ico")
		return
	}

	// Get a redis connection
	//redis, err := db.RedisPool.Get()
	//if err != nil {
	//	log.Fatal(err)
	//}

	// Grab the URL and check if it exists
	//resp, err := redis.Cmd("GET", fmt.Sprintf("url:%s", shortUrl)).Str()
	//db.RedisPool.Put(redis)
	//if err != nil {
	//	fmt.Println("This URL does not exist")
	//	return
	//}
	fmt.Println("redirecting")
	http.Redirect(w, r, "testurl", http.StatusMovedPermanently)
}

// Shorten Handles grabbing the long url from the client, checking if it is a valid URL
// then calling ShortenURL to encode and store the URL.
func Shorten(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

// LoadTemplates Handles loading all HTML files.
func LoadTemplates() {
	indexTemplate = template.Must(template.ParseFiles("./templates/index.html"))
}
