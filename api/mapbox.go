package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/sebas7603/weather-app-go/models"
	"github.com/sebas7603/weather-app-go/utils/helpers"
)

var mapboxResponse models.MapboxResponse
var mapboxParams = map[string]string{
	"limit": "6",
	"types": "place",
}

func RequestMapboxAPI(searchValue string) (*models.MapboxResponse, error) {
	mapboxParams["access_token"] = os.Getenv("MAPBOX_TOKEN")
	requestMapboxURL := fmt.Sprintf("%s/%s.json", os.Getenv("MAPBOX_URL"), searchValue)

	response, err := http.Get(helpers.BuildURLWithParams(requestMapboxURL, mapboxParams))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return parseMapboxResponse(response)
}

func parseMapboxResponse(response *http.Response) (*models.MapboxResponse, error) {
	err := json.NewDecoder(response.Body).Decode(&mapboxResponse)
	if err != nil {
		return nil, err
	}

	return &mapboxResponse, nil
}
