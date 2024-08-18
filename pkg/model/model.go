package model

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/giovalgas/hosts/pkg/styles"
)

type Model struct {
	styles *styles.Styles
}

func NewModel() *Model {
	return &Model{
		styles: styles.NewStyles(lipgloss.DefaultRenderer()),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m Model) View() string {

	s := "Hello, World!"
	s += "\nPress q or ctrl+c to quit.\n"

	return m.styles.Base.Render(s)
}
