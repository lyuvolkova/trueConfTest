package storage

import "refactoring/internal"

func (r *repository) GetUser(id string) (*internal.User, error) {
	if _, ok := r.s.List[id]; !ok {
		return nil, ErrUserNotFound
	}
	tmp := r.s.List[id]
	return &tmp, nil

}
