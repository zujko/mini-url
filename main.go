package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mediocregopher/radix.v2/pool"
	"github.com/zujko/mini-url/db"
	"github.com/zujko/mini-url/views"
)

func main() {
	var err error
	PORT := ":8080"
	db.RedisPool, err = pool.New("tcp", "localhost:6379", 10)
	if err != nil {
		log.Fatal(err)
	}

	router := httprouter.New()
	views.LoadTemplates()
	router.GET("/", views.Index)
	router.GET("/:shorturl", views.HandleUrl)
	router.POST("/shorten", views.Shorten)
	// This sidesteps the core rules of httprouter, but this is only for a dev env so eh
	router.NotFound = http.FileServer(http.Dir("./static"))
	log.Fatal(http.ListenAndServe(PORT, router))
}
