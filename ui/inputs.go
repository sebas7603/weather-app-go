package ui

import (
	"github.com/AlecAivazis/survey/v2"
)

var searchInput = &survey.Input{Message: "Place to search:"}

func ShowPlaceInput() (searchValue string) {
	survey.AskOne(searchInput, &searchValue, survey.WithValidator(survey.Required))
	return
}

func ShowContinue() {
	var rContinue bool
	survey.AskOne(qContinue, &rContinue)
	return
}
