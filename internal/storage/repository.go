package storage

import (
	"sync"

	"refactoring/internal"
)

type repository struct {
	s      *internal.UserStore
	sMutex *sync.Mutex
	store  string
}

func NewRepository(store string) *repository {
	return &repository{
		sMutex: new(sync.Mutex),
		store:  store,
	}
}

func (r *repository) Load() (err error) {
	r.s, err = ReadStore(r.store)
	return
}

func (r *repository) Flush() (err error) {
	return WriteStore(r.store, r.s)
}
