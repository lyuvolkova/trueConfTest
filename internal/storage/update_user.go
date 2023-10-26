package storage

func (r *repository) UpdateUser(id string, displayName string) error {
	if !r.sMutex.TryLock() {
		return ErrStorageBusy
	}
	defer r.sMutex.Unlock()

	if _, ok := r.s.List[id]; !ok {
		return ErrUserNotFound
	}

	u := r.s.List[id]
	u.DisplayName = displayName
	r.s.List[id] = u

	return r.Flush()
}
