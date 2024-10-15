package db

import (
	"os"
)

func InitDB() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	if _, err := os.Stat(getTrackrPath(homeDir)); os.IsExist(err) {
		return nil
	}

	err = createTrackrDir(homeDir)
	if err != nil {
		return err
	}

	return nil
}

func Write() error {
  return nil
}

func Read() error {
  return nil
}
