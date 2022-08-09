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
	ageCellWidth        = lipgloss.Width("6mo ago")
	prStatusCellWidth   = lipgloss.Width(" CI ")
	prRepoNameCellWidth = 15
	prAuthorCellWidth   = 15
	WaitingGlyph        = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "007", Dark: "249"}).Render("")
	FailureGlyph        = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "001", Dark: "001"}).Render("")
	SuccessGlyph        = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "002", Dark: "002"}).Render("")

	headerHeight = 2

	cellStyle = lipgloss.NewStyle().
			PaddingLeft(1).
			PaddingRight(1).
			MaxHeight(1)

	selectedCellStyle = cellStyle.Copy().
				Background(lipgloss.AdaptiveColor{Light: "006", Dark: "008"})

	titleCellStyle = cellStyle.Copy().
			Bold(true).
			Foreground(lipgloss.AdaptiveColor{Light: "000", Dark: "015"})

	singleRuneTitleCellStyle = titleCellStyle.Copy().Width(3)

	headerStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.AdaptiveColor{Light: "254", Dark: "000"}).
			BorderBottom(true)

	rowStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.AdaptiveColor{Light: "254", Dark: "000"}).
			BorderBottom(true)
)

type Column struct {
	Title string
	Width *int
	Grow  *bool
}

type Row []string

func DetailView(m Model) string {
	return m.Viewport.View()
}

func (m Model) ViewportContent(width int) string {
	var builder strings.Builder
	hc := lipgloss.JoinHorizontal(lipgloss.Top, m.renderHeaders()...)
	builder.WriteString(headerStyle.Copy().Width(m.Viewport.Width).MaxWidth(m.Viewport.Width).Render(hc))

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

func (m Model) renderHeaders() []string {
	headerColumns := []Column{
		{
			// Title
			Title: "Title",
			Width: &ageCellWidth,
		},
		{
			// Repo
			Title: "Repository",
			Width: &prRepoNameCellWidth,
		},
		{
			// PR Status
			Title: "Status",
			Width: &prStatusCellWidth,
		},
	}

	renderedColumns := make([]string, len(headerColumns))
	takenWidth := 0
	numGrowingColumns := 0

	for i, col := range headerColumns {
		if col.Grow != nil && *col.Grow {
			numGrowingColumns += 1
			continue
		}

		if col.Width != nil {
			renderedColumns[i] = titleCellStyle.Copy().
				Width(*col.Width).
				MaxWidth(*col.Width).
				Render(col.Title)
			takenWidth += *col.Width
			continue
		}

		// Cell is a glyph
		if len(col.Title) == 1 {
			takenWidth += 3
			renderedColumns[i] = singleRuneTitleCellStyle.Copy().
				Width(3).
				MaxWidth(3).
				Render(col.Title)
			continue
		}

		cell := titleCellStyle.Copy().Render(col.Title)
		renderedColumns[i] = cell
		takenWidth += lipgloss.Width(cell)
	}

	leftoverWidth := m.Viewport.Width - takenWidth
	if numGrowingColumns == 0 {
		return renderedColumns
	}

	growCellWidth := leftoverWidth / numGrowingColumns
	for i, col := range headerColumns {
		if col.Grow == nil || !*col.Grow {
			continue
		}

		renderedColumns[i] = titleCellStyle.Copy().Width(growCellWidth).MaxWidth(growCellWidth).Render(col.Title)
	}

	return renderedColumns
}

func (m Model) renderTitle(title, number string) string {
	return lipgloss.NewStyle().
		Inline(true).
		Foreground(lipgloss.Color("#c7c467")).
		Render(
			fmt.Sprintf("#%s %s", number, title),
		)
}

func (m Model) renderRepo(repo string) string {
	return lipgloss.NewStyle().
		Render(
			fmt.Sprintf("%s", repo),
		)
}

func (m Model) renderStatus(status string) string {
	ciCellStyle := lipgloss.NewStyle()
	if status == "SUCCESS" {
		return ciCellStyle.Render(SuccessGlyph)
	}

	if status == "PENDING" {
		return ciCellStyle.Render(WaitingGlyph)
	}

	return ciCellStyle.Render(FailureGlyph)
}
