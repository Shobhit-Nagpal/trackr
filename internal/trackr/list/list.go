package list

import (
	"fmt"
	"log"
	"os"

	"github.com/Shobhit-Nagpal/trackr/internal/db"
	"github.com/Shobhit-Nagpal/trackr/internal/trackr/view"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	listView = iota
	viewView
)

type ListModel struct {
	sessionState int
	projects     []string
	cursor       int
	selected     int
	view         tea.Model
}

func InitialListModel() ListModel {
	projects := db.GetProjects()
	return ListModel{
		projects: projects,
		cursor:   0,
		selected: 0,
		view:     view.InitialViewModel(),
	}
}

func (m ListModel) Init() tea.Cmd {
	return nil
}

func (m ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch m.sessionState {
	case viewView:
		newView, newCmd := m.view.Update(msg)
		viewModel, ok := newView.(view.ViewModel)
		if !ok {
			log.Fatalf("Error from list model during assertion: view")
			return m, tea.Quit
		}
		m.view = viewModel
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
				if m.cursor < len(m.projects)-1 {
					m.cursor++
				}
			case "enter", " ":
				m.sessionState = viewView
			}

		}

	}
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m ListModel) View() string {

	switch m.sessionState {
	case viewView:
		s := view.GetRenderedMarkdown(m.projects[m.selected])
		return s
	default:

		if len(m.projects) == 0 {
			s := "\n\nNo projects to view\n\n"
			return s
		}

		s := "\n\nChoose a project to view\n\n"

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
}

func Render() {
	p := tea.NewProgram(InitialListModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
