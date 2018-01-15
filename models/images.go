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

func (is *imageService) mkImagePath(galleryID uint) (string, error) {
	galleryPath := filepath.Join("images", "galleries",
		fmt.Sprintf("%v", galleryID))
	err := os.MkdirAll(galleryPath, 0755)
	if err != nil {
		return "", err
	}
	return galleryPath, nil
}