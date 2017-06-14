package main

import (
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"
	"github.com/zujko/mini-url/db"
	"github.com/zujko/mini-url/views"
)

func main() {
	var err error
	PORT := ":8080"
	df := func(network, addr string) (*redis.Client, error) {
		client, err := redis.Dial(network, addr)
		if err != nil {
			return nil, err
		}
		if err = client.Cmd("AUTH", os.Getenv("REDISKEY")).Err; err != nil {
			client.Close()
			return nil, err
		}
		return client, nil
	}

	db.RedisPool, err = pool.NewCustom("tcp", "localhost:6379", 10, df)
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
