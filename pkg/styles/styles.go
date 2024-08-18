package styles

import (
	"github.com/charmbracelet/lipgloss"
)

const MAUVE string = "#cba6f7"

type Styles struct {
	Base lipgloss.Style
}

func NewStyles(lg *lipgloss.Renderer) *Styles {
	s := Styles{}

	s.Base = lg.NewStyle().
		Padding(1, 2).
		Border(lipgloss.RoundedBorder(), true).
		BorderForeground(lipgloss.Color(MAUVE))

	return &s
}
