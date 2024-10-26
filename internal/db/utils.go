package db

import (
	"fmt"
	"os"
)

const TRACKR_DIR = ".trackr"

func getTrackrPath(homeDir string) string {
	return fmt.Sprintf("%s/%s", homeDir, TRACKR_DIR)
}

func createTrackrDir(homeDir string) error {
	err := os.Chdir(homeDir)
	if err != nil {
		return err
	}

	err = os.Mkdir(getTrackrPath(homeDir), 0755)
	if err != nil {
		return err
	}

	return nil
}

func projectExists(homeDir, name string) bool {
	trackrDir := getTrackrPath(homeDir)
	if _, err := os.Stat(getProjectPath(trackrDir, name)); !os.IsNotExist(err) {
		return true
	}

	return false
}

func getProjectPath(trackrDir, name string) string {
	return fmt.Sprintf("%s/%s", trackrDir, name)
}

//Creation

func createProject(homeDir, name, link string) error {
	trackrDir := getTrackrPath(homeDir)
	err := createProjectDir(trackrDir, name)
	if err != nil {
		return err
	}

	projectPath := getProjectPath(trackrDir, name)
	err = createMarkdownFile(projectPath, name, link)
	if err != nil {
		return err
	}
	return nil
}

func createProjectDir(trackrDir, name string) error {
	projectPath := getProjectPath(trackrDir, name)
	err := os.Mkdir(projectPath, 0755)
	if err != nil {
		return err
	}

	return nil
}

func createMarkdownFile(projectPath, name, link string) error {
	fileName := fmt.Sprintf("%s/TODO.md", projectPath)
  initialContent := fmt.Sprintf("# %s\n\n- Link: %s\n\n### To do:\n\n[] Step 1", name, link)
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write([]byte(initialContent))
	if err != nil {
		return err
	}

	return nil
}

// Removal
func removeProject(homeDir, name string) error {
	trackrDir := getTrackrPath(homeDir)
	err := removeProjectDir(trackrDir, name)
	if err != nil {
		return err
	}

	return nil
}

func removeProjectDir(trackrDir, name string) error {
	projectPath := getProjectPath(trackrDir, name)
	err := os.RemoveAll(projectPath)
	if err != nil {
		return err
	}

	return nil
}

// Read
func getProject(homeDir, name string) (string, error) {
	trackrDir := getTrackrPath(homeDir)
	projectPath := getProjectPath(trackrDir, name)
  todoFile := fmt.Sprintf("%s/%s", projectPath, "TODO.md")

	contentBytes, err := os.ReadFile(todoFile)
	if err != nil {
		return "", err
	}

	return string(contentBytes), nil
}
