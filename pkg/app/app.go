package app

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
}

func ApplicationModel() Model {
	return Model{}
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
	return s
}

func RunApp() {
	_, err := tea.NewProgram(ApplicationModel()).Run()

	if err != nil {
		_ = fmt.Errorf("error ocurred while initilializing the app: %s", err)
	}
}
