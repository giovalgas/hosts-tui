package views

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/giovalgas/hosts/pkg/keys"
	"github.com/giovalgas/hosts/pkg/styles"
)

type Views struct {
	List list.Model
}

func NewViews(width int, height int, keys keys.Keys, styles styles.Styles) *Views {
	v := Views{}
	v.List = *NewHostsList(width, height, keys, styles)

	return &v
}
