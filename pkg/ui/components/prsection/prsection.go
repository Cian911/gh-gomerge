package prsection

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/cian911/gh-gomerge/pkg/api"
	"github.com/cian911/gh-gomerge/pkg/ui/components/section"
	"github.com/cian911/gh-gomerge/pkg/ui/components/table"
)

var (
	ciCellWidth        = lipgloss.Width(" CI ")
	linesCellWidth     = lipgloss.Width(" 123450 / -123450 ")
	updatedAtCellWidth = lipgloss.Width("2mo ago")
	prRepoCellWidth    = 15
	prAuthorCellWidth  = 15
	ContainerPadding   = 1

	containerStyle = lipgloss.NewStyle().
			Padding(0, ContainerPadding)

	spinnerStyle = lipgloss.NewStyle().Padding(0, 1)
)

type Model struct {
	Prs     []api.PullRequest
	section section.Model
}

func NewModel(id int) Model {
	m := Model{
		Prs: []api.PullRequest{},
		section: section.Model{
			Id:        id,
			IsLoading: true,
			Type:      "prs_section",
		},
	}

	// m.section.Table =
	return m
}

func (m Model) SectionColumns() []table.Column {
	return []table.Column{
		{
			Title: "Title",
			Width: &prAuthorCellWidth,
		},
		{
			Title: "Author",
			Width: &prAuthorCellWidth,
		},
	}
}
