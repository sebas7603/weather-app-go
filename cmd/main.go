package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sebas7603/weather-app-go/api"
	"github.com/sebas7603/weather-app-go/config"
	"github.com/sebas7603/weather-app-go/ui"
	"github.com/sebas7603/weather-app-go/utils"
	"github.com/sebas7603/weather-app-go/utils/helpers"
)

var err error
var dbPath string
var searchHistory []string

func Start() error {
	dbPath, err = config.InitialConfig()
	if err != nil {
		return err
	}

	// Read search history from file
	err = utils.ReadFromFile(dbPath, &searchHistory)
	if err != nil {
		return err
	}

	mainOption := ""
	for mainOption != "Exit" {
		mainOption = ui.ShowMainMenu()
		switch mainOption {

		case "New Search":
			err = searchAndDisplayPlaceData()
			if err != nil {
				return err
			}

			ui.ShowContinue()
			break

		case "Search History":
			for _, search := range searchHistory {
				fmt.Printf("%s\n", strings.Title(search))
			}
			fmt.Printf("\n\n")

			ui.ShowContinue()
			break
		}
	}

	return nil
}

func searchAndDisplayPlaceData() error {
	searchValue := ui.ShowPlaceInput()

	// Request to Mapbox API
	mapboxData, err := api.RequestMapboxAPI(searchValue)
	if err != nil {
		fmt.Println("Error in Mapbox request")
		return err
	}

	// Creating menu options slice
	placeIndexes := []string{"0"}
	for i := 0; i < len(mapboxData.Features); i++ {
		placeIndexes = append(placeIndexes, strconv.Itoa(i+1))
	}

	// Menu for select a place
	selectedPlace := ui.ShowPlacesMenu(placeIndexes, mapboxData)
	if selectedPlace == nil {
		return nil
	}

	// Save search history in file
	searchHistory = helpers.PrependSliceWithLimitAvoidDuplicates(searchHistory, selectedPlace.PlaceName, 6)
	err = utils.WriteToFileReplacingData(dbPath, searchHistory)
	if err != nil {
		fmt.Println("Error writing in file")
		return err
	}

	// Request to Open Weather API
	openweatherData, err := api.RequestOpenWeatherAPI(selectedPlace.Center[1], selectedPlace.Center[0])
	if err != nil {
		fmt.Println("Error in Open Weather request")
		return err
	}

	// Show the result data
	ui.PrintDataResults(*selectedPlace, *openweatherData)

	return nil
}
