package storage

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"

	"refactoring/internal"
)

func WriteStore(store string, s *internal.UserStore) error {
	b, err := json.Marshal(s)
	if err != nil {
		return fmt.Errorf("marshal file: %w", err)
	}
	err = os.WriteFile(store, b, fs.ModePerm)
	if err != nil {
		return fmt.Errorf("write file: %w", err)
	}
	return nil
}
