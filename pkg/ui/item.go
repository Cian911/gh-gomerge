package ui

import "time"

type Item struct {
	Number int
	Title  string
	Body   string
	Author struct {
		Login string
	}
	Mergeable      string
	State          string
	IsDraft        bool
	ReviewDecision string
	Additions      int
	Deletions      int
	HeadRepository struct {
		Name string
	}
	UpdatedAt time.Time
}

func (i Item) Description() string { return i.Body }

func (i Item) FilterValue() string { return i.Title }
