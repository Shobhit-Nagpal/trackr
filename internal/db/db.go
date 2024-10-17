package db

import (
	"fmt"
	"log"
	"os"
)

func InitDB() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	if _, err := os.Stat(getTrackrPath(homeDir)); !os.IsNotExist(err) {
		return nil
	}

	err = createTrackrDir(homeDir)
	if err != nil {
		return err
	}

	return nil
}

func CreateProject(name, link string) {
	_, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("DB ERROR: %s", err.Error())
	}

  //Check if project already exists
  if projectExists(name) {
    log.Fatalf("%s already exists\n", name)
  }

  err = createProject(name, link)
  if err != nil {
    log.Fatalf("Project creation error: %s", err.Error())
  }
}

func GetProjects() []string {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("DB ERROR: %s", err.Error())
	}

	contents, err := os.ReadDir(getTrackrPath(homeDir))
	if err != nil {
		log.Fatalf("DB ERROR: %s", err.Error())
	}

	projects := []string{}

	for _, content := range contents {
		projects = append(projects, content.Name())
	}

	return projects
}
