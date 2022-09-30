package filesystem

import (
	"os"
)

// checks if folder/file exists
func CheckExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
