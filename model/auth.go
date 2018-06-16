package model

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// UserAuth is an authentication system using the user model
type UserAuth struct {
	User *UserService
}

var (
	// ErrAuthFailed is an error when an authentication attempt was unsuccessful
	ErrAuthFailed = errors.New("Username or password was invalid")
)

// NewUserAuth returns a new UserAuth model
func NewUserAuth(connInfo string) (*UserAuth, error) {
	us, err := NewUserService(connInfo)
	if err != nil {
		return nil, err
	}

	return &UserAuth{
		User: us,
	}, nil
}

// Create will hash a user's password
func (u *UserAuth) Create(user *User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashed)
	return u.User.Create(user)
}

// Update will update a users password if it is set
func (u *UserAuth) Update(user *User) error {
	// Test to see if the user password is already hashed
	_, err := bcrypt.Cost([]byte(user.Password))
	if err != nil {
		// If there was an error when attempting to decode the hash cost, then
		// the password is not encrypted - so update it
		hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashed)
	}
	return u.User.Update(user)
}

// Authenticate will authenticate a user's email and password
func (u *UserAuth) Authenticate(email, password string) (*User, error) {
	user, err := u.User.GetByEmail(email)
	if err != nil {
		if err == ErrNotFound {
			// Return ErrAuthFailed so we don't expose whether the username or the
			// password is incorrect
			return nil, ErrAuthFailed
		}

		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			// Return ErrAuthFailed so we don't expose whether the username or the
			// password is incorrect
			return nil, ErrAuthFailed
		}

		return nil, err
	}

	return user, nil
}
