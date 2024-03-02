package makedirectory

import (
	"os"
	"path/filepath"
)

func Create(path string) error {
	// Check if the directory exists
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	// If the directory does not exist, create its parent
	if os.IsNotExist(err) {
		err = Create(filepath.Dir(path))
		if err != nil {
			return err
		}
		// Create the directory
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
