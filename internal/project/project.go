package project

import (
	"CherryPatch/internal/filesystem"
	"fmt"
	"log"
	"os"
)

// type Project struct {
// 	name string
// }

func CreateProject(name string) error {
	cherryRootPath := filesystem.GetCherryRootDir()
	// check if .cherrypatch exists
	if !filesystem.CherryRootInitialized() {
		os.Exit(1)
	}
	// check if project already exists in same name
	// TODO ignore case for name
	projectPath := fmt.Sprintf("%s/%s", cherryRootPath, name)
	if !filesystem.CheckExists(projectPath) {
		// create new folder
		fmt.Printf("project %s does not exist, creating...\n", name)
		filesystem.CreateDir(projectPath)
		fmt.Println("created")
	}

	// switch to project
	// open .cherrypatch/project.state
	projectStateFilePath := fmt.Sprintf("%s/project.state", cherryRootPath)
	projectStateFile := filesystem.OpenFile(projectStateFilePath)
	defer filesystem.CloseFile(projectStateFile)
	// write new project state to file
	_, err := projectStateFile.WriteString(fmt.Sprintf("%s\n", name))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("switched to project", name)

	return nil
}
