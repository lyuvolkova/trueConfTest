package storage

import (
	"sync"

	"refactoring/internal"
)

type repository struct {
	s      *internal.UserStore
	sMutex *sync.Mutex
}

func NewRepository() *repository {
	return &repository{
		sMutex: new(sync.Mutex),
	}
}

func (r *repository) Load() (err error) {
	r.s, err = ReadStore()
	return
}

func (r *repository) Flush() (err error) {
	return WriteStore(r.s)
}
