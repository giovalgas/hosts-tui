package views

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/giovalgas/hosts/pkg/keys"
	"strings"
)

type Item struct {
	host        string
	hostname    []string
	description string
}

func (i Item) Title() string {
	return strings.Join(i.hostname, " ")
}

func (i Item) Description() string {
	return fmt.Sprintf("Host: %s, Description: %s", i.host, i.description)
}

func (i Item) FilterValue() string {
	return strings.Join(i.hostname, ",") + "," + i.host
}

func NewItem(host string, hostname []string, description string) Item {

	if description == "" {
		description = "No description provided..."
	}

	return Item{host: host, hostname: hostname, description: description}
}

func NewHostsList(width int, height int, keys keys.Keys) *list.Model {
	// Get Hosts
	items := []list.Item{
		NewItem("127.0.0.1", []string{"localhost"}, "Port 5432"),
		NewItem("127.0.0.1", []string{"secret", "testeeeeeee"}, ""),
		NewItem("111.0.0.1", []string{"vm"}, ""),
		NewItem("122.0.0.1", []string{"test"}, ""),
		NewItem("133.0.0.1", []string{"test-host"}, ""),
	}

	l := list.New(items, list.NewDefaultDelegate(), width, height)

	l.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{keys.Copy}
	}

	l.Title = "Managing Hosts..."

	l.SetWidth(width)
	l.SetHeight(height)

	return &l
}
