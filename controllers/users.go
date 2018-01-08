package controllers

import (
	"net/http"

	"../views"
)

// Users struct
type Users struct {
	NewView *views.View
}

// NewUsers - returns a Users Struct
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

// New - handler to handle web requests when a user visits
// the signup page
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}
