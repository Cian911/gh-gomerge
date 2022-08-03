package sidebar

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	IsOpen     bool
	viewport   viewport.Model
	stateEmpty string
}

var (
	sidebarStyle = lipgloss.NewStyle().Padding(0, 2).BorderLeft(true).BorderStyle(lipgloss.Border{
		Top:         "",
		Bottom:      "",
		Left:        "",
		Right:       "",
		TopLeft:     "",
		TopRight:    "",
		BottomRight: "",
		BottomLeft:  "",
	}).BorderForeground(lipgloss.Color("#ccc"))
)

func NewModel() Model {
	return Model{
		IsOpen: true,
		viewport: viewport.Model{
			Width:  100,
			Height: 1000,
		},
		stateEmpty: "",
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		// Jump halfway down the page
		case key.Matches(msg, key.NewBinding(key.WithKeys("ctrl+u"), key.WithHelp("Ctrl+u", "preview page up"))):
			m.viewport.HalfViewDown()

		// Jump halfway up the page
		case key.Matches(msg, key.NewBinding(key.WithKeys("ctrl+d"), key.WithHelp("Ctrl+d", "preview page down"))):
			m.viewport.HalfViewUp()
		}
	}

	return m, nil
}

func (m Model) View() string {
	if !m.IsOpen {
		return ""
	}

	style := sidebarStyle.Copy().Height(100).MaxHeight(200).Width(300).MaxWidth(500)

	if m.viewport.View() == "" {
		return style.Copy().Align(lipgloss.Center).Render(
			lipgloss.PlaceVertical(100, lipgloss.Center, m.stateEmpty),
		)
	}

	return style.Copy().Render(lipgloss.JoinVertical(
		lipgloss.Top,
		m.viewport.View(),
		style.Copy().Render(fmt.Sprintf("%d%%", int(m.viewport.ScrollPercent()*100))),
	))
}

func (m Model) SetContent(data string) {
	m.viewport.SetContent(data)
}
