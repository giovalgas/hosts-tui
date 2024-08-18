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
		item{desc: "127.0.0.1", title: "localhost"},
		item{desc: "127.0.0.1", title: "test-host"},
		item{desc: "172.0.0.1", title: "ssh-machine-host"},
		item{desc: "172.0.0.1", title: "super-host"},
	}

	l := list.New(items, list.NewDefaultDelegate(), width, height)
	l.Title = "Managing Hosts..."

	l.SetWidth(width)
	l.SetHeight(height)

	return &l
}
