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
	// ErrInvalidID is an error when an invalid ID is specified
	ErrInvalidID = errors.New("user: specified ID was invalid")
)

// User is a user in the system
type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"not null;unique_index"`
	IsAdmin  bool
	Password string `gorm:"not null"`
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

// Delete will remove the user from the database
func (s *UserService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidID
	}
	user := User{Model: gorm.Model{ID: id}}
	return s.db.Delete(&user).Error
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

// GetByEmail returns the User from the database with the given email.
// If the user is not found, then ErrNotFound is returned
func (s *UserService) GetByEmail(email string) (*User, error) {
	var user User
	q := s.db.Where("email = ?", email)
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

// Update updates an existing user in the database
func (s *UserService) Update(user *User) error {
	return s.db.Save(user).Error
}

func getOne(db *gorm.DB, d interface{}) error {
	err := db.First(d).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err
}
