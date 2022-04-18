package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	// sections *[]section
	viewport viewport.Model
	err      error
	isReady  bool
}

type initScreen struct{}

type errMsg struct{ error }

func NewModel() Model {
	return Model{
		isReady: false,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(tea.EnterAltScreen)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is this a key prress?
	case tea.KeyMsg:

		// Yes, check the key
		switch msg.String() {

		// Close app
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			// Move up

		case "down", "j":
			// Move down

		case "m":
			// Select PR to me merged

		case "a":
			// Select PR to be approved

		case "enter", " ":
			// Open confirmation section

		}

	case initScreen:
		// Show initScreen with loading message

	case errMsg:
		m.err = msg
		return m, nil
	}

	return m, nil
}

func (m Model) View() string {
	s := strings.Builder{}

	return s.String()
}
