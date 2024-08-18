package model

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/giovalgas/hosts/pkg/styles"
	"github.com/giovalgas/hosts/pkg/views"
	"golang.org/x/term"
	"os"
)

type Model struct {
	styles *styles.Styles
	views  *views.Views
	lg     *lipgloss.Renderer
	width  int
	height int
}

func NewModel() *Model {
	width, height, _ := getScreenSize()
	lg := lipgloss.DefaultRenderer()

	return &Model{
		styles: styles.NewStyles(lg),
		views:  views.NewViews(width, height),
		lg:     lg,
		width:  width,
		height: height,
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
	case tea.WindowSizeMsg:
		m.width = msg.Width - 2
		m.height = msg.Height - 4

		m.views.List.SetSize(m.width, m.height)
	}

	var cmd tea.Cmd
	m.views.List, cmd = m.views.List.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return m.styles.Base.
		Width(m.width).
		Height(m.height).
		Render(m.views.List.View())
}

func getScreenSize() (width int, height int, error error) {
	width, height, error = term.GetSize(int(os.Stdout.Fd()))

	return width - 2, height - 2, error
}
