package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/sebas7603/weather-app-go/utils"
)

var dbFolder = "db"
var dbPath = fmt.Sprintf("%s/database.json", dbFolder)

func InitialConfig() (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}

	err = utils.CheckDatabasePath(dbPath)
	if err != nil {
		return "", err
	}

	return dbPath, nil
}
