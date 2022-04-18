package constants

import "github.com/charmbracelet/lipgloss"

var (
	WaitingGlyph = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#2b2b40"}).Render("")
	FailureGlyph = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#F23D5C", Dark: "#F23D5C"}).Render("")
	SuccessGlyph = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#3DF294", Dark: "#3DF294"}).Render("")
)
