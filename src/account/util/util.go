package util

import (
	"encoding/json"
	"fmt"
	"io"
)

func FromJsonToModel(reader io.ReadCloser, data interface{}) error {
	return json.NewDecoder(reader).Decode(data)
}

func BuildUrl(hostname string, path string, queryParameters map[string]string) string {
	return hostname + path + buildQueryParams(queryParameters)
}

func buildQueryParams(queryParameters map[string]string) string {
	if queryParameters == nil || len(queryParameters) == 0 {
		return ""
	}

	queryParamString := "?"

	for key, value := range queryParameters {
		queryParamString += fmt.Sprintf("%s=%s&", key, value)
	}

	return queryParamString
}
