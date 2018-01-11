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

var _ GalleryDB = &galleryGorm{}

// Create a gallery
func (gg *galleryGorm) Create(gallery *Gallery) error {
	return gg.db.Create(gallery).Error
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
