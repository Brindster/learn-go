package main

import (
	"net/http"
	"os"

	"chrisbrindley.co.uk/controller"
	"chrisbrindley.co.uk/model"
	"chrisbrindley.co.uk/service"
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
	var s map[string]service.Factory
	s = make(map[string]service.Factory)

	s["Model/Db"] = model.NewDbConnection
	s["Model/UserAuth"] = model.NewUserAuth
	s["Model/UserService"] = model.NewUserService
	s["Controller/UserController"] = controller.NewUser

	srv := service.NewServices(s)

	staticController := controller.NewStaticController()
	userController, _ := srv.MustGet("Controller/UserController").(*controller.User)

	r := mux.NewRouter()

	r.Handle("/", staticController.Views["index"])
	r.Handle("/contact", staticController.Views["contact"])
	r.Handle("/faq", staticController.Views["faq"])
	r.HandleFunc("/new", userController.NewHandler)

	var h http.Handler = http.HandlerFunc(notFoundHandler)
	r.NotFoundHandler = h

	http.ListenAndServe(":8000", r)
}
