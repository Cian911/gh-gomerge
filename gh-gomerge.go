package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cian911/gh-gomerge/pkg/ui"
)

func main() {
	p := tea.NewProgram(ui.NewModel(), tea.WithAltScreen())

	if err := p.Start(); err != nil {
		log.Fatalf("Failed to start program: %v", err)
	}
}
