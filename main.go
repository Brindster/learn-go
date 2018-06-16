package main

import (
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"chrisbrindley.co.uk/controller"
	"chrisbrindley.co.uk/view"
	"github.com/gorilla/mux"
)

var (
	conn = os.Getenv("DB_CONN")
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)

	view := view.NewView("main", "view/404.gohtml")
	view.Render(w, nil)
}

func main() {
	staticController := controller.NewStaticController()
	userController := controller.NewUser(conn)

	r := mux.NewRouter()

	r.Handle("/", staticController.Views["index"])
	r.Handle("/contact", staticController.Views["contact"])
	r.Handle("/faq", staticController.Views["faq"])
	r.HandleFunc("/new", userController.NewHandler)

	var h http.Handler = http.HandlerFunc(notFoundHandler)
	r.NotFoundHandler = h

	http.ListenAndServe(":8000", r)
}
