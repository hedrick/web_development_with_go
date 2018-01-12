package controllers

import (
	"../models"
	"../views"
)

// Galleries struct
type Galleries struct {
	New *views.View
	gs  models.GalleryService
}

// NewGalleries returns a new Galleries type to be rendered
func NewGalleries(gs models.GalleryService) *Galleries {
	return &Galleries{
		New: views.NewView("bootstrap", "galleries/new"),
		gs:  gs,
	}
}
