package main

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
var searches []string

func main() {
	dbPath, err = config.InitialConfig()
	if err != nil {
		fmt.Println("Error applying intial config:", err)
		return
	}

	// Read search history from file
	err = utils.ReadFromFile(dbPath, &searches)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	mainOption := ""
	for mainOption != "Exit" {
		mainOption = ui.ShowMainMenu()
		switch mainOption {

		case "New Search":
			searchValue := ui.ShowPlaceInput()

			// Request to Mapbox API
			mapboxData, err := api.RequestMapboxAPI(searchValue)
			if err != nil {
				fmt.Println("Error in Mapbox request:", err)
				return
			}

			// Creating menu options slice
			placeIndexes := []string{"0"}
			for i := 0; i < len(mapboxData.Features); i++ {
				placeIndexes = append(placeIndexes, strconv.Itoa(i+1))
			}

			// Menu for select a place
			selectedPlace := ui.ShowPlacesMenu(placeIndexes, mapboxData)
			if selectedPlace == nil {
				break
			}

			// Avoid duplicate entries
			for i, search := range searches {
				if strings.ToLower(search) == strings.ToLower(selectedPlace.PlaceName) {
					searches = helpers.RemoveFromSliceByIndex(searches, i)
					break
				}
			}
			searches = helpers.PrependSliceWithLimit(searches, selectedPlace.PlaceName, 6)

			// Save search history in file
			err = utils.WriteToFileReplacingData(dbPath, searches)
			if err != nil {
				fmt.Println("Error writing in file:", err)
				return
			}

			// Request to Open Weather API
			openweatherData, err := api.RequestOpenWeatherAPI(selectedPlace.Center[1], selectedPlace.Center[0])
			if err != nil {
				fmt.Println("Error in Open Weather request:", err)
				return
			}

			// Show the result data
			ui.PrintDataResults(*selectedPlace, *openweatherData)
			ui.ShowContinue()
			break

		case "Search History":
			for _, search := range searches {
				fmt.Printf("%s\n", strings.Title(search))
			}
			fmt.Printf("\n\n")

			ui.ShowContinue()
			break
		}
	}
}
