package update

import (
	"fmt"
	"os"

	"github.com/Shobhit-Nagpal/trackr/internal/trackr/view"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type UpdateModel struct {
	textarea textarea.Model
  markdown  string
	err      error
}

type errMsg error

func InitialUpdateModel(name string) UpdateModel {
	md := view.GetRenderedMarkdown(name)
	ti := textarea.New()
	ti.Placeholder = "Once upon a time..."
	ti.Focus()

	return UpdateModel{
		textarea: ti,
    markdown: md,
		err:      nil,
	}
}

func (m UpdateModel) Init() tea.Cmd {
	return nil
}

func (m UpdateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			if m.textarea.Focused() {
				m.textarea.Blur()
			}
		case tea.KeyCtrlC:
			return m, tea.Quit
		default:
			if !m.textarea.Focused() {
				cmd = m.textarea.Focus()
				cmds = append(cmds, cmd)
			}
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m UpdateModel) View() string {
	return fmt.Sprintf(
		"Tell me a story.\n\n%s\n\n%s",
		m.textarea.View(),
		"(ctrl+c to quit)",
	) + "\n\n"
}

func Render(name string) {
	p := tea.NewProgram(InitialUpdateModel(name))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
