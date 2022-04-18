package section

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/cian911/gh-gomerge/pkg/ui/components/table"
)

type Section interface {
	Id() int
	Update(msg tea.Msg) (Section, tea.Cmd)
	View() string
	NumRows() int
	NextRow() int
	PrevRow() int
	GetSectionRows() tea.Cmd
	GetSectionColumns() []table.Column
	IsLoading() bool
	ConstructRows() []table.Row
}

type Model struct {
	Id        int
	Spinner   spinner.Model
	IsLoading bool
	Table     table.Model
	Type      string
}
