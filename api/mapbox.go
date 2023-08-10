package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/sebas7603/weather-app-go/models"
)

var mapboxResponse models.MapboxResponse
var mapboxParams = map[string]string{
	"limit": "6",
	"types": "place",
}

func RequestMapboxAPI(searchValue string) (*models.MapboxResponse, error) {
	requestMapboxURL := fmt.Sprintf("%s/%s.json?access_token=%s", os.Getenv("MAPBOX_URL"), searchValue, os.Getenv("MAPBOX_TOKEN"))
	response, err := http.Get(buildUrlWithParams(requestMapboxURL, mapboxParams))
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

func buildUrlWithParams(baseURL string, params map[string]string) (paramsURL string) {
	paramsURL = baseURL
	for index, param := range params {
		paramsURL = fmt.Sprintf("%s&%s=%s", paramsURL, index, param)
	}
	return
}
