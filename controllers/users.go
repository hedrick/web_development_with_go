package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/schema"

	"../views"
)

// Users struct
type Users struct {
	NewView *views.View
}

// SignupForm struct
type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
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

// Create is used to process the signup form when a user
// tries to create a new user account.

// Create POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	dec := schema.NewDecoder()
	form := SignupForm{}
	if err := dec.Decode(&form, r.PostForm); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, form)
}
