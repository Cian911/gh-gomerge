package main

import (
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	data "github.com/cian911/gh-gomerge/pkg/api"
	"github.com/cian911/gh-gomerge/pkg/ui/components/section"
)

type model struct {
	choices  []data.PullRequest
	prs      []section.Section
	cursor   int              // which to-do list item our cursor is pointing at
	selected map[int]struct{} // which to-do items are selected
}

func initModel(c data.GhClient) model {
	return model{
		choices: c.ListPullRequests(),

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Check for key press
	case tea.KeyMsg:
		// Yes, what key?
		switch msg.String() {

		// Exit app
		case "ctrl+c", "q":
			return m, tea.Quit

		// Move keys up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// Move keys down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// Select item
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m model) getCurrSection() section.Section {
	sections := m.getCurrentViewSections()
	if len(sections) == 0 {
		return nil
	}
	return sections[0]
}

func (m model) setCurrentViewSections(newSections []section.Section) {
	m.prs = newSections
}

func (m model) getCurrentViewSections() []section.Section {
	return m.prs
}

func (m model) View() string {
	// s := "Pull Requests\n\n\n"
	s := strings.Builder{}

	content := lipgloss.JoinHorizontal(lipgloss.Top, m.getCurrSection().View())

	s.WriteString(content)
	s.WriteString("\n")
	return s.String()

	/* for i, choice := range m.choices { */
	/* cursor := " " */
	/* if m.cursor == i { */
	/*   cursor = ">" */
	/* } */
	/*  */
	/* checked := " " */
	/* if _, ok := m.selected[i]; ok { */
	/*   checked = "x" */
	/* } */
	/*  */
	/* s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice.Title) */
	/* } */

	// s += "\nPress q to quit.\n"
}

func main() {
	c := data.New()
	p := tea.NewProgram(
		initModel(*c),
		tea.WithAltScreen())
	if err := p.Start(); err != nil {
		log.Fatalf("Opps, failed: %v", err)
	}
}
