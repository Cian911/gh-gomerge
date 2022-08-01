package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	statusBarStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
			Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"})

	statusStyle = statusBarStyle.Copy().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#FF5F87")).
			Padding(0, 1).
			MarginRight(1)
)

func StatusView(m Model) string {
	var status string
	var statusDesc string

	switch m.Ready {
	case true:
		status = "Ready"
		statusDesc = "Ready to merge."
	case false:
		status = m.Spinner.View()
		statusDesc = "Loading..."
	}

	statusKey := statusStyle.Render(status)
	statusValue := statusBarStyle.Copy().
		Width(m.Width - lipgloss.Width(statusKey)).
		Render(statusDesc)

	statusBar := lipgloss.JoinHorizontal(lipgloss.Top, statusKey, statusValue)

	return statusStyle.Width(m.Width).Render(statusBar)
}
