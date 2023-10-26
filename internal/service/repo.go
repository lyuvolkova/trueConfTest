package service

import "refactoring/internal"

type repo interface {
	CreateUser(displayName string, email string) (string, error)
	DeleteUser(id string) error
	UpdateUser(id string, displayName string) error
	SearchUsers() (internal.UserList, error)
	GetUser(id string) (*internal.User, error)
	Load() (err error)
}
