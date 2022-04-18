package table

import "github.com/charmbracelet/lipgloss"

var (
	headerHeight = 2

	cellStyle = lipgloss.NewStyle().PaddingLeft(1).PaddingRight(1).MaxHeight(1)

	titleCellStyle = cellStyle.Copy().Background(lipgloss.Color("#383838"))

	headerStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#ddd")).
			BorderBottom(true)
)
