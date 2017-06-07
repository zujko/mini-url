package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/zujko/mini-url/views"
)

func main() {
	PORT := ":8080"
	router := httprouter.New()
	views.LoadTemplates()
	router.GET("/", views.Index)
	router.GET("/:shorturl", views.HandleUrl)
	// This sidesteps the core rules of httprouter, but this is only for a dev env so eh
	router.NotFound = http.FileServer(http.Dir("./static"))
	log.Fatal(http.ListenAndServe(PORT, router))
}
