package controller

import "chrisbrindley.co.uk/view"

// Static is a controller handling static pages
type Static struct {
	Views map[string]*view.View
}

// NewStaticController returns a new instance of the static controller
func NewStaticController() *Static {
	views := make(map[string]*view.View)
	views["index"] = view.NewView("main", "view/static/index.gohtml")
	views["contact"] = view.NewView("main", "view/static/contact.gohtml")
	views["faq"] = view.NewView("main", "view/static/faq.gohtml")

	return &Static{
		Views: views,
	}
}
