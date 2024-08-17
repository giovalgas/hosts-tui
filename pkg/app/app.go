package app

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
	"os"
	"os/exec"
)

type Model struct{}

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
	width, height, _ := term.GetSize(int(os.Stdin.Fd()))

	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		PaddingTop(2).
		PaddingLeft(4).
		Width(width).
		Height(height)

	s := "Hello, World!"

	s += "\nPress q or ctrl+c to quit.\n"
	return style.Render(s)
}

func RunApp() {
	_, err := tea.NewProgram(ApplicationModel()).Run()

	defer func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		_ = cmd.Run()
	}()

	if err != nil {
		_ = fmt.Errorf("error ocurred while initilializing the app: %s", err)
	}
}
