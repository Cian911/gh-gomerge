package table

import "github.com/charmbracelet/lipgloss"

type Model struct {
	Columns    []Column
	Rows       []Row
	EmptyState *string
}

type Column struct {
	Title string
	Width *int
	Grow  *bool
}

type Row []string

func NewModel(columns []Column, rows []Row, emptyState *string) Model {
	return Model{
		Columns:    columns,
		Rows:       rows,
		EmptyState: emptyState,
	}
}

func (m Model) View() string {
	header := m.renderHeader()

	return lipgloss.JoinVertical(lipgloss.Left, header, "")
}

func (m Model) renderHeader() string {
	headerColumns := m.renderHeaderColumns()
	header := lipgloss.JoinHorizontal(lipgloss.Top, headerColumns...)

	return headerStyle.Copy().Width(5).MaxWidth(8).Render(header)
}

func (m Model) renderHeaderColumns() []string {
	renderedColumns := make([]string, len(m.Columns))
	takenWidth := 0
	numGrowingColumns := 0

	for i, column := range m.Columns {
		if column.Grow != nil && *column.Grow {
			numGrowingColumns += 1
			continue
		}

		if column.Width != nil {
			renderedColumns[i] = titleCellStyle.Copy().
				Width(*column.Width).
				MaxWidth(*column.Width).
				Render(column.Title)

			takenWidth += *column.Width
		}

		cell := titleCellStyle.Copy().Render(column.Title)
		renderedColumns[i] = cell
		takenWidth += lipgloss.Width(cell)
	}

	leftoverWidth := 1 - takenWidth //m.dimensions.Width - takenWidth
	if numGrowingColumns == 0 {
		return renderedColumns
	}

	growCellWidth := leftoverWidth / numGrowingColumns
	for i, column := range m.Columns {
		if column.Grow == nil || !*column.Grow {
			continue
		}

		renderedColumns[i] = titleCellStyle.Copy().Width(growCellWidth).MaxWidth(growCellWidth).Render(column.Title)
	}

	return renderedColumns
}
