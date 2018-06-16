package model

import (
	"errors"
	"github.com/jinzhu/gorm"

	// Import mysql dialect
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	// ErrNotFound is an error when a record is not found in the database
	ErrNotFound = errors.New("user: resource not found")
)

// User is a user in the system
type User struct {
	gorm.Model
	Name    string
	Email   string `gorm:"not null;unique_index"`
	IsAdmin bool
}

// UserService is an ORM abstraction layer
type UserService struct {
	db *gorm.DB
}

// NewUserService creates a new UserService struct
func NewUserService(connInfo string) (*UserService, error) {
	db, err := gorm.Open("mysql", connInfo)
	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	return &UserService{
		db: db,
	}, nil
}

// Close disconnects from the database
func (s *UserService) Close() error {
	return s.db.Close()
}

// Create saves the user in the database, adding in ID
// created at dates, etc
func (s *UserService) Create(user *User) error {
	return s.db.Create(user).Error
}

// GetByID returns the User from the database with the given ID.
// If the user is not found, then ErrNotFound is returned
func (s *UserService) GetByID(id uint) (*User, error) {
	var user User
	q := s.db.Where("id = ?", id)
	if err := getOne(q, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

// Truncate will truncate a table
func (s *UserService) Truncate() {
	s.db.DropTableIfExists(&User{})
	s.db.AutoMigrate(&User{})
}

func getOne(db *gorm.DB, d interface{}) error {
	err := db.First(d).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err
}
