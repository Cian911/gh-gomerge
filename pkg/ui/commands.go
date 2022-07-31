package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type prsMsg struct{ items []list.Item }

func (m Model) ListPrsCmd() tea.Cmd {
	return func() tea.Msg {
		var (
			items []list.Item
		)

		prs := m.ghClient.ListPullRequests()
		for _, v := range prs {
			items = append(items, Item{
				Title: v.Title,
				Body:  v.Body,
				State: v.State,
			})
		}

		return prsMsg{items: items}
	}
}
