package styles

import (
	"github.com/charmbracelet/lipgloss"
)

const Mauve string = "#cba6f7"
const Green string = "#a6da95"
const Red string = "#ed8796"
const Lavender string = "#b4befe"

type Styles struct {
	Base               lipgloss.Style
	StatusMessage      lipgloss.Style
	ErrorStatusMessage lipgloss.Style
	SelectedItem       lipgloss.Style
}

func NewStyles(lg *lipgloss.Renderer) *Styles {
	s := Styles{}

	s.Base = lg.NewStyle().
		Padding(1, 2).
		Border(lipgloss.RoundedBorder(), true).
		BorderForeground(lipgloss.Color(Lavender))

	s.StatusMessage = lipgloss.NewStyle().
		Foreground(lipgloss.Color(Green))

	s.ErrorStatusMessage = lipgloss.NewStyle().
		Foreground(lipgloss.Color(Red))

	s.SelectedItem = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(lipgloss.Color(Mauve)).
		Foreground(lipgloss.Color(Mauve)).
		Padding(0, 0, 0, 1)

	return &s
}
