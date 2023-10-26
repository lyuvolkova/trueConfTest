package storage

import "errors"

var (
	ErrStorageBusy  = errors.New("file storage busy")
	ErrUserNotFound = errors.New("user not found")
)
