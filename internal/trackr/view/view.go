package view

import (
	"fmt"
	"log"
	"os"

	"github.com/Shobhit-Nagpal/trackr/internal/db"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
)

type ViewModel struct {
	projects []string
	cursor   int
	selected int
}

func InitialViewModel() ViewModel {
	projects := db.GetProjects()
	return ViewModel{
		projects: projects,
		cursor:   0,
		selected: 0,
	}
}

func (m ViewModel) Init() tea.Cmd {
	return nil
}

func (m ViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "j", "down":
			if m.cursor < len(m.projects)-1 {
				m.cursor++
			}
		case "enter", " ":
			RenderProject(m.projects[m.selected])
		}

	}

	return m, nil
}

func (m ViewModel) View() string {

	if len(m.projects) == 0 {
		s := "\n\nNo projects to view\n\n"
		return s
	}

	s := "Choose a project to view\n\n"

	for idx, project := range m.projects {
		cursor := " "
		if m.cursor == idx {
			cursor = ">"
		}

		checked := " "
		if m.selected == idx {
			checked = "X"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, project)
	}

	return s
}

func Render() {
	p := tea.NewProgram(InitialViewModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func RenderProject(name string) {
	project := db.GetProject(name)
	out, err := glamour.Render(project, "dark")
	if err != nil {
		log.Fatalf("Error rendering project: %s", err.Error())
	}
	fmt.Print(out)
}
