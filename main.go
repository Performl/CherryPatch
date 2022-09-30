package main

import (
	"CherryPatch/internal/filesystem"
	"fmt"
)

func main() {
	// create $HOME/.cherrypatch
	if cherry_root_dir := filesystem.GetCherryRootDir(); !filesystem.CheckExists(cherry_root_dir) {
		filesystem.InitCherryRoot()
		fmt.Println("Cherry Patch Root initialised")
	}
}
