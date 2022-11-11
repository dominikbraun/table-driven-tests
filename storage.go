package main

import "errors"

var ErrUserNotFound = errors.New("couldn't find user")

type User struct {
	ID   int
	Name string
}

type Storage struct {
	users []User
}

// NewStorage initializes and returns a new Storage instance that utilizes the provided
// user slice for storing its users.
func NewStorage(users []User) Storage {
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
