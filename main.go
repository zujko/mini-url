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
	router.ServeFiles("/static/*filepath", http.Dir("./static"))
	log.Fatal(http.ListenAndServe(PORT, router))
}
