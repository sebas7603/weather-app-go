package ui

import (
	"fmt"

	"github.com/sebas7603/weather-app-go/models"
)

var Black = "\x1b[30m"
var Red = "\x1b[31m"
var Green = "\x1b[32m"
var Yellow = "\x1b[33m"
var Blue = "\x1b[34m"
var Magenta = "\x1b[35m"
var Cyan = "\x1b[36m"
var White = "\x1b[37m"
var HiBlack = "\x1b[90m"
var HiRed = "\x1b[91m"
var HiGreen = "\x1b[92m"
var HiYellow = "\x1b[93m"
var HiBlue = "\x1b[94m"
var HiMagenta = "\x1b[95m"
var HiCyan = "\x1b[96m"
var HiWhite = "\x1b[97m"

func PrintDataResults(place models.Feature, weather models.OpenWeatherResponse) {
	fmt.Printf("\nCity Info: %s%s\n", HiRed, place.Text)
	fmt.Printf("%sCity: %s%s\n", HiYellow, HiBlue, place.PlaceName)
	fmt.Printf("%sLon:  %s%v\n", HiYellow, HiBlue, place.Center[0])
	fmt.Printf("%sLat:  %s%v\n", HiYellow, HiBlue, place.Center[1])

	fmt.Printf("\n%sWeather Info: %s%s\n", White, HiRed, weather.Weather[0].Description)
	fmt.Printf("%sTemp: %s%v °C\n", HiGreen, HiCyan, weather.Main.Temp)
	fmt.Printf("%sMin:  %s%v °C\n", HiGreen, HiCyan, weather.Main.TempMin)
	fmt.Printf("%sMax:  %s%v °C\n", HiGreen, HiCyan, weather.Main.TempMax)
	fmt.Printf("%sWind: %s%v km/h\n", HiGreen, HiCyan, weather.Wind.Speed)
	fmt.Printf("%s\n", White)

	return
}
