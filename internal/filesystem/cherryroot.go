package filesystem

import (
	"fmt"
	"log"
	"os"
)

func CherryInitIfNotExist() error {
	// create $HOME/.cherrypatch
	if !CherryRootInitialized() {
		initCherryRoot()
		fmt.Println("Cherry Patch Root initialised")
	}
	return nil
}

func GetCherryRootDir() string {
	usr_home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s/.cherrypatch", usr_home)
}

func CherryRootInitialized() bool {
	return CheckExists(GetCherryRootDir())
}

// initialises cherry root in $HOME/.cherrypatch
func initCherryRoot() {
	CreateDir(GetCherryRootDir())
}
