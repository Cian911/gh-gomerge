package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

var (
	dividerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#9B9B9B", Dark: "#5C5C5C"})
	selectedStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#FFFFFF")).
			Foreground(lipgloss.Color("#000000"))
)

func DetailView(m Model) string {
	return m.Viewport.View()
}

func (m Model) ViewportContent(width int) string {
	var builder strings.Builder
	divider := dividerStyle.Render(strings.Repeat("-", width)) + "\n"
	for _, it := range m.PrList.Items() {
		value := fmt.Sprintf("%s\n", it.(Item).Title)
		if it == m.PrList.SelectedItem() {
			value = fmt.Sprintf("%s\n", selectedStyle.Render(it.(Item).Title))
		}
		builder.WriteString(divider)
		builder.WriteString(value)
	}

	return wordwrap.String(builder.String(), width)
}
