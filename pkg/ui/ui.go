package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cian911/gh-gomerge/pkg/api"
	"github.com/spf13/cast"
)

type Model struct {
	PrList   list.Model
	Viewport viewport.Model
	Spinner  spinner.Model

	PrDetails string
	Ready     bool

	Width  int
	Height int

	ghClient *api.GhClient
}

func NewModel() Model {
	l := list.New(nil, list.NewDefaultDelegate(), 0, 0)
	l.Title = "gh-gomerge - Pull Requests"
	l.SetShowHelp(false)
	l.SetShowStatusBar(false)
	l.SetShowFilter(false)

	s := spinner.New()
	s.Spinner = spinner.Dot

	gc := api.New()

	return Model{
		PrList:    l,
		Spinner:   s,
		PrDetails: "",
		Ready:     false,
		ghClient:  gc,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(tea.EnterAltScreen, m.Spinner.Tick, m.ListPrsCmd())
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width, m.Height = msg.Width, msg.Height
		statusBarHeight := lipgloss.Height(StatusView(m))
		height := m.Height - statusBarHeight

		prListViewWidth := cast.ToInt(0.7 * float64(m.Width))
		prListWidth := prListViewWidth - listViewStyle.GetHorizontalFrameSize()
		m.PrList.SetSize(prListWidth, height)

		prDetailViewWidth := m.Width - prListViewWidth
		m.Viewport = viewport.New(prDetailViewWidth, height)
		m.Viewport = viewport.New(m.Width, m.Height)
		m.Viewport.MouseWheelEnabled = true
		m.Viewport.SetContent(m.ViewportContent(m.Width - prDetailViewWidth))
	case prsMsg:
		m.PrList.SetItems(msg.items)
		m.PrList, cmd = m.PrList.Update(msg)
		cmds = append(cmds, cmd)
		m.Viewport.GotoTop()
		m.Viewport.SetContent(m.ViewportContent(m.Viewport.Width))
		m.Ready = true
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			cmd = tea.Quit
			cmds = append(cmds, cmd)
		case tea.KeyUp, tea.KeyDown, tea.KeyLeft, tea.KeyRight:
			m.PrList, cmd = m.PrList.Update(msg)
			cmds = append(cmds, cmd)
			m.Viewport.GotoTop()
			m.Viewport.SetContent(m.ViewportContent(m.Viewport.Width))
		}
	default:
		m.PrList, cmd = m.PrList.Update(msg)
		cmds = append(cmds, cmd)

		m.Viewport, cmd = m.Viewport.Update(msg)
		cmds = append(cmds, cmd)
	}

	m.Spinner, cmd = m.Spinner.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return lipgloss.JoinVertical(lipgloss.Right, lipgloss.JoinHorizontal(lipgloss.Top, ListView(m), DetailView(m)))
}
