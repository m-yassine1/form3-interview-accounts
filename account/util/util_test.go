package util

import (
	"account/model"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonParsing(t *testing.T) {
	s := io.NopCloser(strings.NewReader(`{"data":{"attributes":{"country":"GB"},"id":"6fc6ffaf-caa5-4f9f-a2ec-5c0aec46319e","organisation_id":"6fc6ffaf-caa5-4f9f-a2ec-5c0aec46319e","type":"accounts","version":0}}`))
	var data model.AccountData
	err := FromJsonToModel(s, &data)
	assert.Empty(t, err)
	assert.Equal(t, data.Data.ID, "6fc6ffaf-caa5-4f9f-a2ec-5c0aec46319e")
}

func TestUrlBuilding(t *testing.T) {
	m := make(map[string]string)
	m["k1"] = "test"
	url := BuildUrl("test", "/hello", m)
	assert.NotEmpty(t, url)
	assert.Equal(t, url, "test/hello?k1=test&")
}
