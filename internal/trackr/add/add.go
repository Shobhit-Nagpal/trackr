package add

import (
	"fmt"
	"log"

	"github.com/Shobhit-Nagpal/trackr/internal/db"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type AddModel struct {
	inputs  []textinput.Model
	focused int
	err     error
}

type (
	errMsg error
)

const (
	name = iota
	link
)

const (
	hotBlue  = lipgloss.Color("#0092F8")
	darkGray = lipgloss.Color("#767676")
)

var (
	inputStyle    = lipgloss.NewStyle().Foreground(hotBlue)
	continueStyle = lipgloss.NewStyle().Foreground(darkGray)
)

func nameValidator(s string) error {
	if len(s) > 0 {
		return nil
	}

	return fmt.Errorf("Project name cannot be empty")
}

func linkValidator(s string) error {
	if len(s) > 0 {
		return nil
	}

	return fmt.Errorf("GitHub repo should be created for the project")
}

func InitialAddModel() AddModel {
	var inputs []textinput.Model = make([]textinput.Model, 2)
	inputs[name] = textinput.New()
	inputs[name].Placeholder = "Trackr"
	inputs[name].Focus()
	inputs[name].CharLimit = 20
	inputs[name].Width = 30
	inputs[name].Prompt = ""
	inputs[name].Validate = nameValidator

	inputs[link] = textinput.New()
  inputs[link].Placeholder = "https://github.com/Shobhit-Nagpal/..."
	inputs[link].CharLimit = 100
	inputs[link].Width = 80
	inputs[link].Prompt = ""
	inputs[link].Validate = linkValidator

	return AddModel{
		inputs:  inputs,
		focused: 0,
		err:     nil,
	}
}

func (m AddModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m AddModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.focused == len(m.inputs)-1 {
        db.CreateProject(m.inputs[name].Value(), m.inputs[link].Value())
				return m, tea.Quit
			}
			m.nextInput()
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyShiftTab, tea.KeyCtrlP:
			m.prevInput()
		case tea.KeyTab, tea.KeyCtrlN:
			m.nextInput()
		}
		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focused].Focus()

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func (m AddModel) View() string {
	return fmt.Sprintf(
		`Create a new project:

 %s
 %s

 %s  
 %s

 %s
`,
		inputStyle.Width(30).Render("Project Name"),
		m.inputs[name].View(),
		inputStyle.Width(30).Render("Repo Link"),
		m.inputs[link].View(),
		continueStyle.Render("Continue ->"),
	) + "\n"
}

// nextInput focuses the next input field
func (m *AddModel) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

// prevInput focuses the previous input field
func (m *AddModel) prevInput() {
	m.focused--
	// Wrap around
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}

func Render() {
	p := tea.NewProgram(InitialAddModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
