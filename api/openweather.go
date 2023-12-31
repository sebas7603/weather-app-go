package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/sebas7603/weather-app-go/models"
	"github.com/sebas7603/weather-app-go/utils/helpers"
)

var openWeatherResponse models.OpenWeatherResponse
var openWeatherParams = map[string]string{
	"units": "metric",
}

func RequestOpenWeatherAPI(lat, lon float64) (*models.OpenWeatherResponse, error) {
	openWeatherParams["lat"] = fmt.Sprintf("%v", lat)
	openWeatherParams["lon"] = fmt.Sprintf("%v", lon)
	openWeatherParams["appid"] = os.Getenv("OPENWEATHER_TOKEN")

	response, err := http.Get(helpers.BuildURLWithParams(os.Getenv("OPENWEATHER_URL"), openWeatherParams))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return parseOpenWeatherResponse(response)
}

func parseOpenWeatherResponse(response *http.Response) (*models.OpenWeatherResponse, error) {
	err := json.NewDecoder(response.Body).Decode(&openWeatherResponse)
	if err != nil {
		return nil, err
	}

	return &openWeatherResponse, nil
}
