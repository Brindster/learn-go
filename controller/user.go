package controller

import (
	"net/http"

	"chrisbrindley.co.uk/model"
	"chrisbrindley.co.uk/service"
	"chrisbrindley.co.uk/view"
)

// User is a controller to handle user actions
type User struct {
	auth  *model.UserAuth
	views map[string]*view.View
}

// NewUser returns a User controller
func NewUser(c service.Container) (interface{}, error) {
	views := make(map[string]*view.View)
	views["new"] = view.NewView("main", "view/user/new.gohtml")

	auth, ok := c.MustGet("Model/UserAuth").(*model.UserAuth)
	if !ok {
		return nil, service.ErrInvalidType
	}

	return &User{
		auth:  auth,
		views: views,
	}, nil
}

// NewHandler is the HTTP handler for a new User
func (u *User) NewHandler(w http.ResponseWriter, r *http.Request) {
	if err := u.views["new"].Render(w, nil); err != nil {
		panic(err)
	}
}
