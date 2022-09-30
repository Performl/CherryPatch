package filesystem

import (
	"fmt"
	"log"
	"os"
)

func GetCherryRootDir() string {
	usr_home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s/.cherrypatch", usr_home)
}

// initialises cherry root in $HOME
func InitCherryRoot() {
	cherry_root_dir := GetCherryRootDir()
	if err := os.Mkdir(cherry_root_dir, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}
