package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	commands []string
	cursor   int
	selected int
}

func initialModel() model {
	return model{
		commands: []string{"add", "remove", "list", "view"},
		cursor:   0,
		selected: -1,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if m.cursor < len(m.commands)-1 {
				m.cursor++
			}
		case "enter", " ":
			//Handle project selection --> Change view to project view here
		}

	}

	return m, nil
}

func (m model) View() string {
	s := "Choose a command \n\n"
	//Read projects here

	for idx, project := range m.commands {
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

func Initialize() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
