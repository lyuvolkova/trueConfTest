package storage

import "refactoring/internal"

func (r *repository) SearchUsers() (internal.UserList, error) {
	return r.s.List, nil
}
