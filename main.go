package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"chrisbrindley.co.uk/view"
	"github.com/gorilla/mux"
)

var templates map[string]*view.View

var (
	host   = "db"
	port   = "3306"
	user   = os.Getenv("MYSQL_USER")
	pass   = os.Getenv("MYSQL_PASSWORD")
	dbname = os.Getenv("MYSQL_DATABASE")
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

	data := struct {
		Name string
	}{"Chris Brindley"}

	render(w, "home", data)
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

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbname)
	db, err := sql.Open("mysql", conn)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	http.ListenAndServe(":8000", r)
}
