package views

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func setup() {
	indexTemplate = template.Must(template.ParseFiles("../templates/index.html"))
}
func TestIndex(t *testing.T) {
	setup()
	indexRouter := httprouter.New()
	indexRouter.GET("/", Index)

	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	indexRouter.ServeHTTP(w, request)

	if w.Code != http.StatusOK {
		t.Fatal("Status code is not 200")
	}
}
