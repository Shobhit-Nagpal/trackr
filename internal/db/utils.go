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

func createProjectDir(name string) error {
	return nil
}
