package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

func CheckDatabasePath(filePath string) error {
	slicePath := strings.Split(filePath, "/")
	for i := 0; i < len(slicePath)-1; i++ {
		if slicePath[i] != "." {
			err := createFolderIfNotExists(slicePath[i])
			if err != nil {
				return err
			}
		}
	}

	err := createFileIfNotExists(filePath)
	if err != nil {
		return err
	}

	return nil
}

func createFileIfNotExists(filePath string) error {
	_, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			file, err := os.Create(filePath)
			if err != nil {
				return err
			}
			defer file.Close()
		} else {
			return err
		}
	}

	return nil
}

func createFolderIfNotExists(folderPath string) error {
	_, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		err := os.Mkdir(folderPath, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

func ReadFromFile(filePath string, target interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	if len(fileBytes) > 0 {
		err = json.Unmarshal(fileBytes, &target)
		if err != nil {
			return err
		}
	}

	return nil
}

func WriteToFileReplacingData(filePath string, data interface{}) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return err
	}

	return nil
}
