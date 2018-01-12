package controllers

import (
	"fmt"
	"net/http"

	"../context"
	"../models"
	"../views"
)

// Galleries struct
type Galleries struct {
	New *views.View
	gs  models.GalleryService
}

// GalleryForm type
type GalleryForm struct {
	Title string `schema:"title"`
}

// NewGalleries returns a new Galleries type to be rendered
func NewGalleries(gs models.GalleryService) *Galleries {
	return &Galleries{
		New: views.NewView("bootstrap", "galleries/new"),
		gs:  gs,
	}
}

// Create POST /galleries
func (g *Galleries) Create(w http.ResponseWriter, r *http.Request) {
	var vd views.Data
	var form GalleryForm
	if err := parseForm(r, &form); err != nil {
		vd.SetAlert(err)
		g.New.Render(w, vd)
		return
	}
	user := context.User(r.Context())
	gallery := models.Gallery{
		Title:  form.Title,
		UserID: user.ID,
	}
	if err := g.gs.Create(&gallery); err != nil {
		vd.SetAlert(err)
		g.New.Render(w, vd)
		return
	}
	fmt.Fprintln(w, gallery)
}
