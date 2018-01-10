package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

var (
	// ErrNotFound is returned when a resource cannot be found
	// in the database
	ErrNotFound = errors.New("models: resource not found")
	// ErrInvalidID is returned when an invalid ID is provided
	// to a method like Delete.
	ErrInvalidID = errors.New("models: ID provided was invalid")
)

// User struct for GORM model
type User struct {
	gorm.Model
	Name         string
	Email        string `gorm:"not null;unique_index"`
	Password     string `gorm:"-"`
	PasswordHash string `gorm:"not null"`
}

// UserService - Abstraction Layer for Users DB
type UserService struct {
	db *gorm.DB
}

// Create will create the provided user and backfill data
// like the ID, CreatedAt, and UpdatedAt fields.
func (us *UserService) Create(user *User) error {
	hashedBytes, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedBytes)
	user.Password = ""
	return us.db.Create(user).Error
}

// Update will update the provided user with all of the data
// in the provided user object.
func (us *UserService) Update(user *User) error {
	return us.db.Save(user).Error
}

// Delete will delte the user with the provided ID
func (us *UserService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidID
	}
	user := User{Model: gorm.Model{ID: id}}
	return us.db.Delete(&user).Error
}

// AutoMigrate will attempt to automatically migrate the
// users table
func (us *UserService) AutoMigrate() error {
	if err := us.db.AutoMigrate(&User{}).Error; err != nil {
		return err
	}
	return nil
}

/*
	ById will look up a user with the provided ID.
	If the user is found, we return a nil error
	If the user is not found we return ErrNotFound
	If there is another error, we return an error with more
	information about what went wrong. This may not be an
	error generated by the models package.

	As a general rule, any error but ErrNotFound should
	probably result in a 500.
*/
func (us *UserService) ByID(id uint) (*User, error) {
	var user User
	db := us.db.Where("id = ?", id)
	err := first(db, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

/*
	ByEmail looks up a user with the given email address and
	returns that user.
	If the user is found, we will return a nil error
	If the user is not found, we will return ErrNotFound
	If there is another error, we will return an error with
	more information about what went wrong. This may not be
	an error generated by the models package.
*/
func (us *UserService) ByEmail(email string) (*User, error) {
	var user User
	db := us.db.Where("email = ?", email)
	err := first(db, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// NewUserService - creates and returns a new UserService
func NewUserService(connectionInfo string) (*UserService, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return &UserService{
		db: db,
	}, nil
}

// Close - closes UserService database connection
func (us *UserService) Close() error {
	return us.db.Close()
}

// DestructiveReset drops the user table and rebuilds it
func (us *UserService) DestructiveReset() error {
	err := us.db.DropTableIfExists(&User{}).Error
	if err != nil {
		return err
	}
	return us.AutoMigrate()
}

/*
	first will query using the provided gorm.DB and it will
	get the first item returned and place it into dst. If
	nothing is found in the query, it will return ErrNotFound
*/
func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err
}
