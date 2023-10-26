package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"refactoring/internal"
)

const store = `users.json`

func ReadStore() (*internal.UserStore, error) {
	f, err := os.ReadFile(store)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}
	s := internal.UserStore{}
	err = json.Unmarshal(f, &s)
	if err != nil {
		return nil, fmt.Errorf("unmarshal file: %w", err)
	}
	return &s, err
}
