package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type renderParams struct {
	t string
	p interface{}
}

func render(w http.ResponseWriter, p renderParams) {
	t, err := template.ParseFiles(p.t)
	if err != nil {
		panic(err)
	}

	if err = t.Execute(w, p.p); err != nil {
		panic(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	data := struct {
		Name string
	}{"Chris Brindley"}

	render(w, renderParams{t: "view/index.gohtml", p: data})
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	render(w, renderParams{t: "view/contact.gohtml"})
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

	render(w, renderParams{t: "view/faq.gohtml", p: d})
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	render(w, renderParams{t: "view/404.gohtml"})
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/contact", contactHandler)
	r.HandleFunc("/faq", faqHandler)

	var h http.Handler = http.HandlerFunc(notFoundHandler)
	r.NotFoundHandler = h

	http.ListenAndServe(":8000", r)
}
