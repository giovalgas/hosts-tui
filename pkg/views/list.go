package views

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
)

type item struct {
	hostname, host, description string
}

func (i item) Title() string {
	return i.hostname
}

func (i item) Description() string {
	return fmt.Sprintf("Host: %s, Description: %s", i.host, i.description)
}

func (i item) FilterValue() string {
	return i.hostname + i.host
}

func newItem(hostname string, host string, description string) item {

	if description == "" {
		description = "No description provided..."
	}

	return item{host: host, hostname: hostname, description: description}
}

func NewHostsList(width int, height int) *list.Model {
	// Get Hosts
	items := []list.Item{
		newItem("127.0.0.1", "localhost", "Port 5432"),
		newItem("127.0.0.1", "secret", ""),
		newItem("111.0.0.1", "vm", ""),
		newItem("122.0.0.1", "test", ""),
		newItem("133.0.0.1", "test-host", ""),
	}

	l := list.New(items, list.NewDefaultDelegate(), width, height)
	l.Title = "Managing Hosts..."

	l.SetWidth(width)
	l.SetHeight(height)

	return &l
}
