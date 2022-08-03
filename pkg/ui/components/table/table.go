package table

type Model struct {
	Columns []Column
	Rows    []Row
}

type Column struct {
	Title string
	Width *int
	Grow  *bool
}

type Row []string

func NewModel(columns []Column, rows []Row) Model {
	return Model{
		Columns: columns,
		Rows:    rows,
	}
}

func (m Model) View() string {
	return "Table"
}
