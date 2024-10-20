package remove

import (
	"fmt"
	"os"

	"github.com/Shobhit-Nagpal/trackr/internal/db"
	tea "github.com/charmbracelet/bubbletea"
)

type RemoveModel struct {
	projects []string
	cursor   int
	selected int
}

func InitialRemoveModel() RemoveModel {
	projects := db.GetProjects()
	return RemoveModel{
		projects: projects,
		cursor:   0,
		selected: 0,
	}
}

func (m RemoveModel) Init() tea.Cmd {
	return nil
}

func (m RemoveModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		case "enter":
			db.RemoveProject(m.projects[m.cursor])
			return m, tea.Quit
		}

	}

	return m, nil
}

func (m RemoveModel) View() string {
	s := "Choose a project to delete\n\n"
	//Read projects here

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
	p := tea.NewProgram(InitialRemoveModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func Remove(name string) {
	db.RemoveProject(name)
}
