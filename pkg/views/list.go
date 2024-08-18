package views

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/giovalgas/hosts/pkg/keys"
	"github.com/giovalgas/hosts/pkg/reader"
	"github.com/giovalgas/hosts/pkg/styles"
	"time"
)

func newDelegate(styles styles.Styles) *list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.Styles.SelectedTitle = styles.SelectedItem
	d.Styles.SelectedDesc = styles.SelectedItem

	return &d
}

func NewHostsList(width int, height int, keys keys.Keys, styles styles.Styles) *list.Model {
	// Get Hosts
	items := reader.ReadHostsFile()

	l := list.New(items, newDelegate(styles), width, height)

	l.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{keys.Copy}
	}

	l.Title = "Managing Hosts..."

	l.StatusMessageLifetime = 5 * time.Second
	l.SetWidth(width)
	l.SetHeight(height)

	return &l
}
