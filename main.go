package main

import (
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"chrisbrindley.co.uk/model"
	"chrisbrindley.co.uk/view"
	"github.com/gorilla/mux"
)

var templates map[string]*view.View

var (
	conn = os.Getenv("DB_CONN")
)

func initTemplate(alias, path string) {
	if templates == nil {
		templates = make(map[string]*view.View)
	}
	templates[alias] = view.NewView("main", path)
}

func render(w http.ResponseWriter, alias string, params interface{}) {
	if t, ok := templates[alias]; ok {
		if err := t.Render(w, params); err != nil {
			panic(err)
		}
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	us, err := model.NewUserService(conn)
	if err != nil {
		panic(err)
	}

	us.Truncate()

	user := model.User{
		Name:  "Chris Brindley",
		Email: "chris@chrisbrindley.co.uk",
	}

	if err = us.Create(&user); err != nil {
		panic(err)
	}

	fUser, err := us.GetByID(1)
	if err != nil {
		panic(err)
	}

	render(w, "home", fUser)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	render(w, "contact", nil)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	type Question struct {
		Q, A string
	}

	type Data = struct {
		Qs []Question
	}

	qs := []Question{}
	qs = append(qs, Question{Q: "What is this site?", A: "Just a test"})
	qs = append(qs, Question{Q: "Is Go a good language", A: "Too early to tell"})

	d := Data{qs}

	render(w, "faq", d)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	render(w, "error/404", nil)
}

func main() {
	initTemplate("error/404", "view/404.gohtml")
	initTemplate("home", "view/index.gohtml")
	initTemplate("contact", "view/contact.gohtml")
	initTemplate("faq", "view/faq.gohtml")

	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/contact", contactHandler)
	r.HandleFunc("/faq", faqHandler)

	var h http.Handler = http.HandlerFunc(notFoundHandler)
	r.NotFoundHandler = h

	http.ListenAndServe(":8000", r)
}
