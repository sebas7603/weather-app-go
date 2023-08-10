package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/joho/godotenv"

	"github.com/sebas7603/weather-app-go/api"
	"github.com/sebas7603/weather-app-go/utils"
)

var err error

var file *os.File
var dbFolder = "db"
var dbPath = fmt.Sprintf("%s/database.json", dbFolder)

var searches []string

var mainOption string
var mainMenu = &survey.Select{
	Message: "Choose an option:",
	Options: []string{"New Search", "Search History", "Exit"},
}

var rContinue bool
var qContinue = &survey.Confirm{
	Message: "Continue ?",
	Default: true,
}

var searchValue string
var searchInput = &survey.Input{Message: "Place to search:"}

func main() {
	err = godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	err = utils.CheckDatabasePath(dbPath)
	if err != nil {
		fmt.Println("Error in Database path:", err)
		return
	}

	err = utils.ReadFromFile(dbPath, &searches)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	for mainOption != "Exit" {
		// Clear console
		fmt.Print("\033[H\033[2J")

		// First menu
		survey.AskOne(mainMenu, &mainOption)

		switch mainOption {

		case "New Search":
			survey.AskOne(searchInput, &searchValue, survey.WithValidator(survey.Required))

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
			var placeIndex string
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

			// Get the selected place and add to history
			survey.AskOne(placesMenu, &placeIndex)
			if placeIndex == "0" {
				break
			}
			i, _ := strconv.Atoi(placeIndex)
			selectedPlace := mapboxData.Features[i-1]

			// Avoid duplicate entries
			for i, search := range searches {
				if strings.ToLower(search) == strings.ToLower(selectedPlace.PlaceName) {
					searches = removeFromSlice(searches, i)
					break
				}
			}
			if len(searches) < 6 {
				searches = append([]string{strings.ToLower(selectedPlace.PlaceName)}, searches...)
			} else {
				searches = append([]string{strings.ToLower(selectedPlace.PlaceName)}, searches[:5]...)
			}

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
			fmt.Printf("\nCity Info:\n")
			fmt.Printf("City:        %s\n", selectedPlace.PlaceName)
			fmt.Printf("Lon:         %v\n", selectedPlace.Center[0])
			fmt.Printf("Lat:         %v\n", selectedPlace.Center[1])
			fmt.Printf("Temperature: %v\n", openweatherData.Main.Temp)
			fmt.Printf("Min:         %v\n", openweatherData.Main.TempMin)
			fmt.Printf("Max:         %v\n", openweatherData.Main.TempMax)
			fmt.Println("")

			survey.AskOne(qContinue, &rContinue)
			break

		case "Search History":
			for _, search := range searches {
				fmt.Printf("%s\n", strings.Title(search))
			}
			fmt.Printf("\n\n")

			survey.AskOne(qContinue, &rContinue)
			break
		}
	}
}

func buildUrlWithParams(baseURL string, params map[string]string) (paramsURL string) {
	paramsURL = baseURL
	for index, param := range params {
		paramsURL = fmt.Sprintf("%s&%s=%s", paramsURL, index, param)
	}
	return
}

// Remove a value into a slice at a given index, preserving the order of existing elements
func removeFromSlice(array []string, index int) []string {
	return append(array[:index], array[index+1:]...)
}
