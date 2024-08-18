package views

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/giovalgas/hosts/pkg/keys"
)

type Views struct {
	List list.Model
}

func NewViews(width int, height int, keys keys.Keys) *Views {
	v := Views{}
	v.List = *NewHostsList(width, height, keys)

	return &v
}
