package test

import (
	"account/api"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var hostname = "http://localhost:8080"

func TestApiCreation(t *testing.T) {
	accountApi, err := api.NewAccountApi("")
	assert.NotEmpty(t, err, "Error is empty")
	assert.Empty(t, accountApi, "Account API is not empty")

	accountApi, err = api.NewAccountApi("test")
	assert.NotEmpty(t, err, "Error is empty")
	assert.Empty(t, accountApi, "Account API is not empty")

	accountApi, err = api.NewAccountApi(hostname)
	assert.Empty(t, err, "Error is not empty")
	assert.NotEmpty(t, accountApi, "Account API is empty")
}

func TestGetAccounts(t *testing.T) {
	accountApi, _ := api.NewAccountApi(hostname)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/v1/organisation/accounts", hostname),
		httpmock.NewStringResponder(200, `{     "data": [         {             "attributes": {                 "alternative_names": null,                 "country": "GB",                 "name": []             },             "created_on": "2022-10-19T09:03:08.334Z",             "id": "0d209d7f-d07a-4542-947f-5885fddddae7",             "modified_on": "2022-10-19T09:03:08.334Z",             "organisation_id": "ba61483c-d5c5-4f50-ae81-6b8c039bea43",             "type": "accounts",             "version": 0         },         {             "attributes": {                 "alternative_names": null,                 "country": "GB",                 "name": []             },             "created_on": "2022-10-19T09:05:40.235Z",             "id": "0d209d7f-d07a-4542-947f-5885fddddae8",             "modified_on": "2022-10-19T09:05:40.235Z",             "organisation_id": "ba61483c-d5c5-4f50-ae81-6b8c039bea43",             "type": "accounts",             "version": 0         }] }]`))
	accounts, err := accountApi.GetAccounts()
	assert.Empty(t, err, "Error is not empty")
	assert.NotEmpty(t, accounts, "Accounts list is empty")
	assert.Equal(t, len(accounts), 2)
}
