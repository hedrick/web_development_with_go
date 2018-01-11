package models

import (
	"github.com/jinzhu/gorm"
)

// Services struct for services
type Services struct {
	Gallery GalleryService
	User    UserService
}

// NewServices opens a database connection, checks for
// errors, sets log mode to true and the uses that db
// connection to construct individual services
func NewServices(connectionInfo string) (*Services, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return &Services{
		User:    NewUserService(db),
		Gallery: &galleryGorm{},
	}, nil
}
