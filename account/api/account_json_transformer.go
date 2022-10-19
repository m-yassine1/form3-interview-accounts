package api

import (
	"encoding/json"
	"io"
)

func FromJsonToModel(reader io.ReadCloser, data interface{}) error {
	return json.NewDecoder(reader).Decode(data)
}
