package storage

func (r *repository) DeleteUser(id string) error {
	if !r.sMutex.TryLock() {
		return ErrStorageBusy
	}
	defer r.sMutex.Unlock()

	if _, ok := r.s.List[id]; !ok {
		return ErrUserNotFound
	}
	delete(r.s.List, id)

	return r.Flush()
}
