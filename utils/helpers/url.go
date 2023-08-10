package helpers

import (
	"fmt"
	"strings"
)

func BuildURLWithParams(baseURL string, params map[string]string) string {
	paramsURL := baseURL
	hasFirstParam := strings.Contains(paramsURL, "?")
	for index, param := range params {
		if !hasFirstParam {
			paramsURL = fmt.Sprintf("%s?%s=%s", paramsURL, index, param)
			hasFirstParam = true
			continue
		}
		paramsURL = fmt.Sprintf("%s&%s=%s", paramsURL, index, param)
	}

	return paramsURL
}
