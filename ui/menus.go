package ui

import (
	"fmt"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/sebas7603/weather-app-go/models"
)

var mapboxData models.MapboxResponse

var mainMenu = &survey.Select{
	Message: "Choose an option:",
	Options: []string{"New Search", "Search History", "Exit"},
}

var qContinue = &survey.Confirm{
	Message: "Continue ?",
	Default: true,
}

func ShowMainMenu() (mainOption string) {
	fmt.Print("\033[H\033[2J")
	survey.AskOne(mainMenu, &mainOption)
	return
}

func ShowPlacesMenu(placeIndexes []string, mapboxData *models.MapboxResponse) *models.Feature {
	placeIndex := "0"
	placesMenu := &survey.Select{
		Message: "Select a place:",
		Options: placeIndexes,
		Description: func(value string, index int) string {
			if value == "0" {
				return "Cancel"
			}
			i, _ := strconv.Atoi(value)
			return mapboxData.Features[i-1].PlaceName
		},
	}
	survey.AskOne(placesMenu, &placeIndex)

	if placeIndex == "0" {
		return nil
	}
	i, _ := strconv.Atoi(placeIndex)
	selectedPlace := mapboxData.Features[i-1]

	return &selectedPlace
}
