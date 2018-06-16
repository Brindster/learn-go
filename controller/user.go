package controller

import (
	"net/http"

	"chrisbrindley.co.uk/view"
)

// User is a controller to handle user actions
type User struct {
	views map[string]*view.View
}

// NewUser returns a User controller
func NewUser() *User {
	views := make(map[string]*view.View)
	views["new"] = view.NewView("main", "view/user/new.gohtml")

	return &User{
		views: views,
	}
}

// NewHandler is the HTTP handler for a new User
func (u *User) NewHandler(w http.ResponseWriter, r *http.Request) {
	if err := u.views["new"].Render(w, nil); err != nil {
		panic(err)
	}
}
