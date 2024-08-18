package styles

import (
	"github.com/charmbracelet/lipgloss"
)

const Mauve string = "#cba6f7"
const Green string = "#a6da95"

type Styles struct {
	Base          lipgloss.Style
	StatusMessage lipgloss.Style
}

func NewStyles(lg *lipgloss.Renderer) *Styles {
	s := Styles{}

	s.Base = lg.NewStyle().
		Padding(1, 2).
		Border(lipgloss.RoundedBorder(), true).
		BorderForeground(lipgloss.Color(Mauve))

	s.StatusMessage = lipgloss.NewStyle().
		Foreground(lipgloss.Color(Green))

	return &s
}
