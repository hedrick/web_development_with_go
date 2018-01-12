package models

import (
	"github.com/jinzhu/gorm"
)

// Gallery type
type Gallery struct {
	gorm.Model
	UserID uint   `gorm:"not_null;index"`
	Title  string `gorm:"not_null"`
}

// GalleryService interface
type GalleryService interface {
	GalleryDB
}

// GalleryDB interface
type GalleryDB interface {
	Create(gallery *Gallery) error
}

type galleryGorm struct {
	db *gorm.DB
}

type galleryService struct {
	GalleryDB
}

type galleryValidator struct {
	GalleryDB
}

type galleryValFn func(*Gallery) error

var _ GalleryDB = &galleryGorm{}

const (
	// ErrUserIDRequired error
	ErrUserIDRequired modelError = "models: user ID is required"
	// ErrTitleRequired error
	ErrTitleRequired modelError = "models: title is required"
)

// Create a gallery
func (gg *galleryGorm) Create(gallery *Gallery) error {
	return gg.db.Create(gallery).Error
}

// Create a gallery with validation and then call create to the DB
func (gv *galleryValidator) Create(gallery *Gallery) error {
	err := runGalleryValFns(gallery,
		gv.userIDRequired,
		gv.titleRequired)
	if err != nil {
		return err
	}
	return gv.GalleryDB.Create(gallery)
}

// NewGalleryService returns a new instance of a GalleryService
func NewGalleryService(db *gorm.DB) GalleryService {
	return &galleryService{
		GalleryDB: &galleryValidator{
			GalleryDB: &galleryGorm{
				db: db,
			},
		},
	}
}

func (gv *galleryValidator) userIDRequired(g *Gallery) error {
	if g.UserID <= 0 {
		return ErrUserIDRequired
	}
	return nil
}

func (gv *galleryValidator) titleRequired(g *Gallery) error {
	if g.Title == "" {
		return ErrTitleRequired
	}
	return nil
}

func runGalleryValFns(gallery *Gallery, fns ...galleryValFn) error {
	for _, fn := range fns {
		if err := fn(gallery); err != nil {
			return err
		}
	}
	return nil
}
