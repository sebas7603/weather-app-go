package models

type OpenWeatherResponse struct {
	Coord      Coord         `json:"coord"`
	Weather    []WeatherItem `json:"weather"`
	Base       string        `json:"base"`
	Main       Main          `json:"main"`
	Visibility int           `json:"visibility"`
	Wind       Wind          `json:"wind"`
	Clouds     Clouds        `json:"clouds"`
	Dt         int           `json:"dt"`
	Sys        Sys           `json:"sys"`
	Timezone   int           `json:"timezone"`
	ID         int           `json:"id"`
	Name       string        `json:"name"`
	Cod        int           `json:"cod"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type WeatherItem struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  float64 `json:"pressure"`
	Humidity  float64 `json:"humidity"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
}

type Clouds struct {
	All int `json:"all"`
}

type Sys struct {
	Type    int    `json:"type"`
	ID      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
}
