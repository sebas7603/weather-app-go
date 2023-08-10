package models

type MapboxResponse struct {
	Type        string    `json:"type"`
	Query       []string  `json:"query"`
	Features    []Feature `json:"features"`
	Attribution string    `json:"attribution"`
}

type Feature struct {
	ID         string        `json:"id"`
	Type       string        `json:"type"`
	PlaceType  []string      `json:"place_type"`
	Relevance  float64       `json:"relevance"`
	Properties Properties    `json:"properties"`
	Text       string        `json:"text"`
	Language   string        `json:"language"`
	PlaceName  string        `json:"place_name"`
	Bbox       []float64     `json:"bbox"`
	Center     []float64     `json:"center"`
	Geometry   Geometry      `json:"geometry"`
	Context    []ContextData `json:"context"`
}

type Properties struct {
	MapboxID string `json:"mapbox_id"`
	Wikidata string `json:"wikidata"`
}

type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type ContextData struct {
	ID        string `json:"id"`
	MapboxID  string `json:"mapbox_id"`
	Wikidata  string `json:"wikidata"`
	ShortCode string `json:"short_code"`
	Text      string `json:"text"`
	Language  string `json:"language"`
}
