package controller

import (
	"net/http"

	"chrisbrindley.co.uk/model"
	"chrisbrindley.co.uk/view"
)

// User is a controller to handle user actions
type User struct {
	auth  *model.UserAuth
	views map[string]*view.View
}

// NewUser returns a User controller
func NewUser(connInfo string) *User {
	views := make(map[string]*view.View)
	views["new"] = view.NewView("main", "view/user/new.gohtml")

	/**
	 * @todo Inject UserAuth as a dependancy
	 */
	auth, err := model.NewUserAuth(connInfo)
	if err != nil {
		panic(err)
	}

	return &User{
		auth:  auth,
		views: views,
	}
}

// NewHandler is the HTTP handler for a new User
func (u *User) NewHandler(w http.ResponseWriter, r *http.Request) {
	if err := u.views["new"].Render(w, nil); err != nil {
		panic(err)
	}
}
