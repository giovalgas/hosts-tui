package views

import "github.com/charmbracelet/bubbles/list"

type Views struct {
	List list.Model
}

func NewViews(width int, height int) *Views {
	v := Views{}
	v.List = *NewHostsList(width, height)

	return &v
}
