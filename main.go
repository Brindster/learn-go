package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type renderParams interface{}

func render(w http.ResponseWriter, tpl string, data renderParams) {
	t, err := template.ParseFiles(tpl)
	if err != nil {
		panic(err)
	}

	if err = t.Execute(w, data); err != nil {
		panic(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	data := struct {
		Name string
	}{"Chris Brindley"}

	render(w, "view/index.gohtml", data)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<p>Contact me on <a href=\"mailto:chris@chrisbrindley.co.uk\">chris@chrisbrindley.co.uk</a>!</p>")
}
func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h3>Frequently Asked Questions</h3><dl><dt>What the heck is this site</dt><dd>I dunno!</dd></dl>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<p>The page you are looking for could not be found.</p>")
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
