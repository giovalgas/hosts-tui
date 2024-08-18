package keys

import "github.com/charmbracelet/bubbles/key"

type Keys struct {
	Copy key.Binding
}

func NewKeys() Keys {
	return Keys{
		Copy: key.NewBinding(
			key.WithKeys("c"),
			key.WithHelp("c", "Copy hostname"),
		),
	}
}
