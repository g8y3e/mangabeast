package viewmodel

import (
	"github.com/g8y3e/mangabeast/api/model"
)

type Home struct {
	Title string

	ViewModel
}

func NewHome(home *model.Home) Home {
	result := Home{
		Title:     home.Title,
		ViewModel: newViewModel("home"),
	}

	return result
}
