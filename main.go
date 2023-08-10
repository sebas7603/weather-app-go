package main

import (
	"fmt"

	"github.com/sebas7603/weather-app-go/cmd"
)

func main() {
	err := cmd.Start()
	if err != nil {
		fmt.Println("Ops! There was an unexpected error:", err)
	}

	return
}
