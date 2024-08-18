package model

import (
	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/giovalgas/hosts/pkg/keys"
	"github.com/giovalgas/hosts/pkg/styles"
	"github.com/giovalgas/hosts/pkg/views"
	"golang.org/x/term"
	"os"
	"strings"
)

type Model struct {
	styles *styles.Styles
	views  *views.Views
	lg     *lipgloss.Renderer
	keys   *keys.Keys
	width  int
	height int
}

func NewModel() *Model {
	width, height, _ := getScreenSize()
	lg := lipgloss.DefaultRenderer()
	k := keys.NewKeys()

	return &Model{
		styles: styles.NewStyles(lg),
		views:  views.NewViews(width, height, k),
		keys:   &k,
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
		if !m.views.List.FilteringEnabled() {
			switch {
			case msg.String() == "ctrl+c" || msg.String() == "q":
				return m, tea.Quit
			case key.Matches(msg, m.keys.Copy):

				err := clipboard.WriteAll(strings.Split(m.views.List.SelectedItem().FilterValue(), ",")[0])

				if err != nil {
					return m, tea.Quit
				}

				return m, nil
			}
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
