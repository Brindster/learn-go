package view

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	// LayoutDir is the default directory for layouts
	LayoutDir = "view/layout/"

	// TemplateExt is the default file extension for templates
	TemplateExt = ".gohtml"
)

// View holds the view details
type View struct {
	Layout   string
	Template *template.Template
}

// NewView returns a pointer to the view
func NewView(l string, files ...string) *View {
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   l,
	}
}

// Render renders the template held in the view
func (v *View) Render(w http.ResponseWriter, params interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout+TemplateExt, params)
}

// ServeHTTP will render the template
func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.Render(w, nil); err != nil {
		panic(err)
	}
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}
