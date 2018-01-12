package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"../context"
	"../models"
	"../views"
	"github.com/gorilla/mux"
)

// Galleries struct
type Galleries struct {
	New      *views.View
	ShowView *views.View
	gs       models.GalleryService
}

// GalleryForm type
type GalleryForm struct {
	Title string `schema:"title"`
}

// NewGalleries returns a new Galleries type to be rendered
func NewGalleries(gs models.GalleryService) *Galleries {
	return &Galleries{
		New:      views.NewView("bootstrap", "galleries/new"),
		ShowView: views.NewView("bootstrap", "galleries/show"),
		gs:       gs,
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

// Show GET /galleries/:id
func (g *Galleries) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid gallery ID", http.StatusNotFound)
		return
	}
	_ = id

	gallery := models.Gallery{
		Title: "A temporary fake gallery with ID: " + idStr,
	}

	var vd views.Data
	vd.Yield = gallery
	g.ShowView.Render(w, vd)
}
