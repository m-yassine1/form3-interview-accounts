package api

import (
	"account/model"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var hostname = "http://localhost:8080"

func TestApiCreation(t *testing.T) {
	accountApi, err := NewAccountApi("")
	assert.NotEmpty(t, err, "Error is empty")
	assert.Empty(t, accountApi, "Account API is not empty")

	accountApi, err = NewAccountApi("test")
	assert.NotEmpty(t, err, "Error is empty for invalid hostname")
	assert.Empty(t, accountApi, "Account API is not empty for invalid hostname")

	accountApi, err = getAccountApi()
	assert.Empty(t, err, "Error is not empty for valid hostname")
	assert.NotEmpty(t, accountApi, "Account API is empty for valid hostname")
}

func TestAccountApiHealthy(t *testing.T) {
	accountApi, _ := getAccountApi()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/v1/health", hostname),
		httpmock.NewStringResponder(200, `{"status": "up"}`))
	err := accountApi.IsHealthy()
	assert.Empty(t, err, "Error is not empty")
}

func TestAccountApiNotHealthy(t *testing.T) {
	accountApi, _ := getAccountApi()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/v1/health", hostname),
		httpmock.NewStringResponder(200, `{"status": "down"}`))
	err := accountApi.IsHealthy()
	assert.NotEmpty(t, err, "Error is empty")
}

func TestGetAccounts(t *testing.T) {
	accountApi, _ := getAccountApi()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/v1/organisation/accounts", hostname),
		httpmock.NewStringResponder(200, `{     "data": [         {             "attributes": {                 "alternative_names": null,                 "country": "GB",                 "name": []             },             "created_on": "2022-10-19T09:03:08.334Z",             "id": "0d209d7f-d07a-4542-947f-5885fddddae7",             "modified_on": "2022-10-19T09:03:08.334Z",             "organisation_id": "ba61483c-d5c5-4f50-ae81-6b8c039bea43",             "type": "accounts",             "version": 0         },         {             "attributes": {                 "alternative_names": null,                 "country": "GB",                 "name": []             },             "created_on": "2022-10-19T09:05:40.235Z",             "id": "0d209d7f-d07a-4542-947f-5885fddddae8",             "modified_on": "2022-10-19T09:05:40.235Z",             "organisation_id": "ba61483c-d5c5-4f50-ae81-6b8c039bea43",             "type": "accounts",             "version": 0         }] }]`))
	accounts, err := accountApi.GetAccounts(nil)
	assert.Empty(t, err, "Error is not empty")
	assert.NotEmpty(t, accounts, "Accounts list is empty")
	assert.Equal(t, len(accounts), 2)
}

func TestFailedGetAccounts(t *testing.T) {
	accountApi, _ := getAccountApi()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/v1/organisation/accounts", hostname),
		httpmock.NewStringResponder(400, ``))
	accounts, err := accountApi.GetAccounts(nil)
	assert.NotEmpty(t, err, "Error is not empty")
	assert.Empty(t, accounts, "Accounts list is empty")
	assert.Equal(t, len(accounts), 0)
}

func TestGetAccount(t *testing.T) {
	accountApi, _ := getAccountApi()
	id := "0d209d7f-d07a-4542-947f-5885fddddae7"
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/v1/organisation/accounts/%s", hostname, id),
		httpmock.NewStringResponder(200, `{     "data":  {             "attributes": {                 "alternative_names": null,                 "country": "GB",                 "name": []             },   "created_on": "2022-10-19T09:03:08.334Z",             "id": "0d209d7f-d07a-4542-947f-5885fddddae7",             "modified_on": "2022-10-19T09:03:08.334Z",             "organisation_id": "ba61483c-d5c5-4f50-ae81-6b8c039bea43",             "type": "accounts",             "version": 0  }}`))
	account, err := accountApi.GetAccount(id)
	assert.Empty(t, err, "Error is not empty")
	assert.NotEmpty(t, account, "Account is empty")
	assert.Equal(t, account.ID, "0d209d7f-d07a-4542-947f-5885fddddae7")
}

func TestGetNotFoundAccount(t *testing.T) {
	accountApi, _ := getAccountApi()
	id := "0d209d7f-d07a-4542-947f-5885fddddae7"
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/v1/organisation/accounts/%s", hostname, id),
		httpmock.NewStringResponder(404, ``))
	account, err := accountApi.GetAccount(id)
	assert.NotEmpty(t, err, "Error is empty")
	assert.Empty(t, account, "Account is empty")
}

func TestCreateAccount(t *testing.T) {
	accountApi, _ := getAccountApi()
	var version int64 = 0
	names := make([]string, 0)
	country := "GB"
	uuidString := uuid.New().String()
	createAccount := model.AccountData{
		Data: model.Account{
			ID:             uuidString,
			OrganisationID: uuidString,
			Type:           "accounts",
			Version:        &version,
			Attributes: &model.AccountAttributes{
				Name:    names,
				Country: &country,
			},
		},
	}
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/v1/organisation/accounts", hostname),
		httpmock.NewStringResponder(201, `{     "data":  {             "attributes": {                 "alternative_names": null,                 "country": "GB",                 "name": []             },   "created_on": "2022-10-19T09:03:08.334Z",             "id": "`+uuidString+`",             "modified_on": "2022-10-19T09:03:08.334Z",             "organisation_id": "`+uuidString+`",             "type": "accounts",             "version": 0  }}`))
	account, err := accountApi.CreateAccount(createAccount)
	assert.Empty(t, err, "Error is not empty")
	assert.NotEmpty(t, account, "Account is empty")
	assert.Equal(t, account.ID, uuidString)
}

func TestSuccessfulDeletingAccount(t *testing.T) {
	accountApi, _ := getAccountApi()
	id := "0d209d7f-d07a-4542-947f-5885fddddae7"
	version := 0
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", fmt.Sprintf("%s/v1/organisation/accounts/%s?version=%d", hostname, id, version),
		httpmock.NewStringResponder(204, ``))
	err := accountApi.DeleteAccount(id, version)
	assert.Empty(t, err, "Error is not empty")
}

func TestErrorDeletingAccount(t *testing.T) {
	accountApi, _ := getAccountApi()
	id := "0d209d7f-d07a-4542-947f-5885fddddae7"
	version := 0
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", fmt.Sprintf("%s/v1/organisation/accounts/%s?version=%d", hostname, id, version),
		httpmock.NewStringResponder(400, ``))
	err := accountApi.DeleteAccount(id, version)
	assert.NotEmpty(t, err, "Error is empty")
}

func getAccountApi() (*AccountApi, error) {
	return NewAccountApi(hostname)
}
