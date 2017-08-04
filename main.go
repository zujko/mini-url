package main

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"

	"github.com/gorilla/context"
	"github.com/gorilla/pat"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/twitter"

	"github.com/markbates/goth"
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

func init() {
	store := sessions.NewFilesystemStore(os.TempDir(), []byte("goth-example"))

	store.MaxLength(math.MaxInt64)

	gothic.Store = store
}

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
	goth.UseProviders(
		github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), "http://localhost/auth/github/callback"),
		twitter.New(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), "http://localhost/auth/twitter/callback"),
		facebook.New(os.Getenv("FACEBOOK_KEY"), os.Getenv("FACEBOOK_SECRET"), "http://localhost/auth/facebook/callback"),
	)

	// Create router
	router := pat.New()
	views.LoadTemplates()
	router.Post("/shorten", views.Shorten)
	router.Get("/static/", views.StaticHandler)
	router.Get("/auth/{provider}/callback", views.AuthCallback)
	router.Get("/auth/{provider}", views.Auth)
	router.Get("/logout/{provider}", views.Logout)
	router.Get("/{shorturl}", views.HandleUrl)
	router.Get("/", views.Index)
	log.Fatal(http.ListenAndServe(PORT, context.ClearHandler(router)))
}
