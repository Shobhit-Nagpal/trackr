package db

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var style = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#0092F8")).
	PaddingLeft(2)

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

func CreateProject(name, link string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("DB ERROR: %s", err.Error())
	}

	projectName := strings.ToLower(name)

	//Check if project already exists
	if projectExists(homeDir, projectName) {
		log.Fatalf("%s already exists\n", name)
	}

	err = createProject(homeDir, projectName, link)
	if err != nil {
		log.Fatalf("Project creation error: %s", err.Error())
	}

	fmt.Println(style.Render("Project created!"))

}

func RemoveProject(name string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("DB ERROR: %s", err.Error())
	}

	projectName := strings.ToLower(name)

	//Check if project exists
	if !projectExists(homeDir, projectName) {
		log.Fatalf("%s does not exist\n", name)
	}

	err = removeProject(homeDir, projectName)
	if err != nil {
		log.Fatalf("Project deletion error: %s", err.Error())
	}

	fmt.Println(style.Render("Project deleted!"))
}
