package models

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// ImageService type to Create and store images
type ImageService interface {
	Create(galleryID uint, r io.Reader, filename string) error
	ByGalleryID(galleryID uint) ([]string, error)
}

type imageService struct{}

// NewImageService returns a new ImageService type
func NewImageService() ImageService {
	return &imageService{}
}

// Create creates a new image given a gallery ID and Reader
func (is *imageService) Create(galleryID uint, r io.Reader,
	filename string) error {
	path, err := is.mkImagePath(galleryID)
	if err != nil {
		return err
	}
	// Create a destination file
	dst, err := os.Create(filepath.Join(path, filename))
	if err != nil {
		return err
	}
	defer dst.Close()
	// Copy reader data to the destination file
	_, err = io.Copy(dst, r)
	if err != nil {
		return err
	}
	return nil
}

func (is *imageService) ByGalleryID(galleryID uint) ([]string, error) {
	path := is.imagePath(galleryID)
	strings, err := filepath.Glob(filepath.Join(path, "*"))
	if err != nil {
		return nil, err
	}
	return strings, nil
}

func (is *imageService) mkImagePath(galleryID uint) (string, error) {
	galleryPath := is.imagePath(galleryID)
	err := os.MkdirAll(galleryPath, 0755)
	if err != nil {
		return "", err
	}
	return galleryPath, nil
}

// Going to need this when we know it is already made
func (is *imageService) imagePath(galleryID uint) string {
	return filepath.Join("images", "galleries", fmt.Sprintf("%v", galleryID))
}
