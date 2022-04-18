package prssection

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/cian911/gh-gomerge/pkg/api"
	"github.com/cian911/gh-gomerge/pkg/ui/components/section"
	"github.com/cian911/gh-gomerge/pkg/ui/components/table"
)

const sectionType = "prs_section"

type Model struct {
	// List of pull requests
	Prs []api.PullRequest
	// The approrpiate section for the data
	section section.Model
}

func NewModel(id int) Model {
	m := Model{
		Prs: []api.PullRequest{},
		section: section.Model{
			Id:        id,
			Spinner:   spinner.Model{Spinner: spinner.Dot},
			IsLoading: true,
			Type:      sectionType,
		},
	}

	m.section.Table = table.NewModel(m.GetSectionColumns(), m.BuildRows(), "")
}

func (m *Model) GetSectionColumns() []table.Column {
	return []table.Column{
		{
			Title: "age",
			Width: &updatedAtCellWidth,
		},
		{
			Title: "repo",
			Width: &prRepoCellWidth,
		},
		{
			Title: "Title",
			Grow:  BoolPtr(true),
		},
		{
			Title: "Author",
			Width: &prAuthorCellWidth,
		},
	}
}

func BoolPtr(b bool) *bool { return &b }
