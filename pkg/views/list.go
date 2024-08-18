package views

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/giovalgas/hosts/pkg/keys"
	"github.com/giovalgas/hosts/pkg/reader"
)

func NewHostsList(width int, height int, keys keys.Keys) *list.Model {
	// Get Hosts
	items := reader.ReadHostsFile()

	l := list.New(items, list.NewDefaultDelegate(), width, height)

	l.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{keys.Copy}
	}

	l.Title = "Managing Hosts..."

	l.SetWidth(width)
	l.SetHeight(height)

	return &l
}
