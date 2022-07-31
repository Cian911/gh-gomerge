package ui

import "github.com/charmbracelet/lipgloss"

var (
	listViewStyle = lipgloss.NewStyle().
		PaddingRight(1).
		MarginRight(1).
		Border(lipgloss.RoundedBorder(), true, true, false, false)
)

func ListView(m Model) string {
	return listViewStyle.Render(m.PrList.View())
}
