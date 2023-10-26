package storage

import (
	"strconv"
	"time"

	"refactoring/internal"
)

func (r *repository) CreateUser(displayName string, email string) (string, error) {
	if !r.sMutex.TryLock() {
		return "", ErrStorageBusy
	}
	defer r.sMutex.Unlock()

	r.s.Increment++
	u := internal.User{
		CreatedAt:   time.Now(),
		DisplayName: displayName,
		Email:       email,
	}

	id := strconv.Itoa(r.s.Increment)
	r.s.List[id] = u

	err := r.Flush()
	if err != nil {
		return "", err
	}

	return id, nil
}
