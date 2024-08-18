package views

import (
	"github.com/charmbracelet/bubbles/list"
)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func NewHostsList(width int, height int) *list.Model {
	// Get Hosts
	items := []list.Item{
		item{title: "Hello", desc: "World1"},
		item{title: "Hello", desc: "World2"},
		item{title: "Hello", desc: "World3"},
		item{title: "Hello", desc: "World4"},
	}

	l := list.New(items, list.NewDefaultDelegate(), width, height)
	l.Title = "Managing Hosts..."

	l.SetWidth(width)
	l.SetHeight(height)

	return &l
}
