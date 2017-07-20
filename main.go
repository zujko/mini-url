package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/julienschmidt/httprouter"
	"github.com/zujko/mini-url/db"
	"github.com/zujko/mini-url/views"
)

var (
	host    = os.Getenv("DB_HOST")
	port    = os.Getenv("DB_PORT")
	user    = os.Getenv("DB_USER")
	dbName  = os.Getenv("DB_NAME")
	sslmode = os.Getenv("DB_SSL")
)

func main() {
	var err error
	PORT := ":8080"

	// Create db connection
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s", host, port, user, dbName, sslmode)
	db.DBConn, err = sql.Open("postgres", psqlConn)
	if err != nil {
		log.Fatal(err)
	}
	err = db.DBConn.Ping()
	if err != nil {
		log.Fatal("Could not ping DB")
	}
	defer db.DBConn.Close()

	// Create router
	router := httprouter.New()
	views.LoadTemplates()
	router.GET("/", views.Index)
	router.GET("/:shorturl", views.HandleUrl)
	router.POST("/shorten", views.Shorten)
	// This sidesteps the core rules of httprouter, but this is only for a dev env so eh
	router.NotFound = http.FileServer(http.Dir("./static"))
	log.Fatal(http.ListenAndServe(PORT, router))
}
