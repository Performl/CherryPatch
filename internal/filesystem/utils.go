package filesystem

import (
	"log"
	"os"
)

// checks if folder/file exists
func CheckExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// tries to create directory
func CreateDir(path string) {
	if err := os.Mkdir(path, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

func OpenFile(path string) *os.File {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func CloseFile(f *os.File) {
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
