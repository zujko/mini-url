package views

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/zujko/mini-url/db"
)

func AuthCallback(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		log.Fatal(err)
		return
	}
	session, _ := gothic.Store.Get(r, gothic.SessionName)
	fmt.Println("AUTH CALLBACK")
	if !UserExists(user.Email) {
		go StoreUser(user)
		fmt.Println("User does not exist")
	}
	fmt.Println("USER ID", user.UserID)
	session.Values["name"] = user.Name
	session.Values["email"] = user.Email
	err = session.Save(r, w)
	if err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

func StoreUser(user goth.User) {
	var lastInsertId int
	err := db.DBConn.QueryRow("INSERT INTO profile(username,full_name) VALUES($1,$2) RETURNING profile_id", user.NickName, user.Name).Scan(&lastInsertId)
	if err != nil {
		fmt.Println("USER SAVE FAIL")
		log.Fatalf("%+v\n", err)
	}
	fmt.Println("User saved")
}

func Logout(w http.ResponseWriter, r *http.Request) {
	gothic.Logout(w, r)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func Auth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AUTH")
	if _, err := gothic.CompleteUserAuth(w, r); err == nil {
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	} else {
		gothic.BeginAuthHandler(w, r)
	}
}

func UserExists(email string) bool {
	var exists bool
	err := db.DBConn.QueryRow("SELECT EXISTS(SELECT 1 FROM profile WHERE email=$1)", email).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
	return exists
}
