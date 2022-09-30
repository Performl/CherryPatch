package filesystem

import (
	"fmt"
	"log"
	"os"
)

func CherryInitIfNotExist() {
	// create $HOME/.cherrypatch
	if cherry_root_dir := getCherryRootDir(); !CheckExists(cherry_root_dir) {
		initCherryRoot()
		fmt.Println("Cherry Patch Root initialised")
	}
}

func getCherryRootDir() string {
	usr_home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s/.cherrypatch", usr_home)
}

// initialises cherry root in $HOME
func initCherryRoot() {
	cherry_root_dir := getCherryRootDir()
	if err := os.Mkdir(cherry_root_dir, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}
