package main

import (
	"errors"
	"fmt"
)

var (
	ErrUserNotFound      = errors.New("couldn't find user")
	ErrUserAlreadyExists = errors.New("user already exists")
)

type User struct {
	ID   int
	Name string
}

type Storage struct {
	users map[string]User
}

// NewStorage initializes and returns a new Storage instance that utilizes the provided
// user slice for storing its users.
func NewStorage(users map[string]User) Storage {
	return Storage{
		users: users,
	}
}

// FindUser returns the user with the provided ID. If the user cannot be found, FindUser
// returns an ErrUserNotFound.
func (s *Storage) FindUser(id int) (User, error) {
	for _, u := range s.users {
		if u.ID == id {
			return u, nil
		}
	}

	return User{}, ErrUserNotFound
}

// AddUser adds a user with the provided id and name. If a user with the provided name already exists, AddUser
// returns an ErrUserAlreadyExists.
func (s *Storage) AddUser(id int, name string) error {
	if _, found := s.users[name]; found {
		return fmt.Errorf("%w: %q", ErrUserAlreadyExists, name)
	}
	s.users[name] = User{ID: id, Name: name}
	return nil
}
