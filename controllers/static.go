package controllers

import "../views"

// Static struct
type Static struct {
	Home    *views.View
	Contact *views.View
	FAQ     *views.View
}

// NewStatic returns Static struct
func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bootstrap", "static/home"),
		Contact: views.NewView("bootstrap", "static/contact"),
		FAQ:     views.NewView("bootstrap", "static/faq"),
	}
}
