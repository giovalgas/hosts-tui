package styles

import (
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
	"os"
)

type Styles struct {
	Base lipgloss.Style
}

func NewStyles(lg *lipgloss.Renderer) *Styles {
	width, height, _ := term.GetSize(int(os.Stdin.Fd()))
	s := Styles{}

	s.Base = lg.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		PaddingTop(2).
		PaddingLeft(4).
		Width(width).
		Height(height)

	return &s
}
