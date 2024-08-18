package app

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/giovalgas/hosts/pkg/model"
	"os"
	"os/exec"
)

func RunApp() {
	_, err := tea.NewProgram(model.NewModel(), tea.WithAltScreen()).Run()

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
