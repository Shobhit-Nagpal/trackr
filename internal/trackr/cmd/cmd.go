package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/Shobhit-Nagpal/trackr/internal/trackr/add"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	cmdView = iota
	addView
	listView
	removeView
)

type CmdModel struct {
	sessionState int
	commands     []string
	add          tea.Model
	remove       tea.Model
	list         tea.Model
	view         tea.Model
	cursor       int
	selected     int
}

func initialModel() CmdModel {
	return CmdModel{
		commands: []string{"add", "remove", "list", "view"},
		add:      add.InitialAddModel(),
		cursor:   0,
		selected: 0,
	}
}

func (m CmdModel) Init() tea.Cmd {
	return nil
}

func (m CmdModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch m.sessionState {
	case addView:
		newAdd, newCmd := m.add.Update(msg)
		addModel, ok := newAdd.(add.AddModel)
		if !ok {
			log.Fatalf("Error from cmd model during assertion: add")
			return m, tea.Quit
		}
		m.add = addModel
		cmd = newCmd
	default:
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
				//Handle sessionState selection --> Change view to cmd view here
				m.sessionState = addView
			}

		}

	}
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m CmdModel) View() string {
	switch m.sessionState {
	case addView:
		viewString := m.add.View()
		return viewString
	default:
    //For the choosing of cmd
		s := "\n\nChoose a command \n\n"
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

			s += fmt.Sprintf("%s [%s] %s\n\n", cursor, checked, project)
		}
		return s
	}

}

func Render() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
