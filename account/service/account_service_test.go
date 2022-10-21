package service

import (
	"account/api"
	"account/model"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var hostname = "http://localhost:8080"

func TestServiceCreation(t *testing.T) {
	accountService, err := NewAccountService(nil)
	assert.NotEmpty(t, err, "Error is empty")
	assert.Empty(t, accountService, "Account Service is not empty")

	accountService, err = getAccountService()
	assert.Empty(t, err, "Error is not empty for valid hostname")
	assert.NotEmpty(t, accountService, "Account API is empty for valid hostname")
}

func TestGetAccounts(t *testing.T) {
	accountService, _ := getAccountService()
	accounts, err := accountService.GetAccounts(nil)
	assert.Empty(t, err, "Error is not empty")
	assert.NotEmpty(t, accounts, "Accounts list is empty")
	assert.Equal(t, len(accounts), 2)
}

func TestFailedGetAccounts(t *testing.T) {
	accountService, _ := getAccountService()
	accounts, err := accountService.GetAccounts(nil)
	assert.NotEmpty(t, err, "Error is not empty")
	assert.NotEmpty(t, accounts, "Accounts list is empty")
	assert.Equal(t, len(accounts), 0)
}

func TestGetAccount(t *testing.T) {
	accountService, _ := getAccountService()
	id := "0d209d7f-d07a-4542-947f-5885fddddae7"
	account, err := accountService.GetAccount(id)
	assert.Empty(t, err, "Error is not empty")
	assert.NotEmpty(t, account, "Account is empty")
	assert.Equal(t, account.ID, "0d209d7f-d07a-4542-947f-5885fddddae7")
}

func TestGetNotFoundAccount(t *testing.T) {
	accountService, _ := getAccountService()
	id := "0d209d7f-d07a-4542-947f-5885fddddae7"
	account, err := accountService.GetAccount(id)
	assert.NotEmpty(t, err, "Error is empty")
	assert.Empty(t, account, "Account is empty")
}

func TestCreateAccount(t *testing.T) {
	accountService, _ := getAccountService()
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

	account, err := accountService.CreateAccount(createAccount)
	assert.Empty(t, err, "Error is not empty")
	assert.NotEmpty(t, account, "Account is empty")
	assert.Equal(t, account.ID, uuidString)
}

func TestSuccessfulDeletingAccount(t *testing.T) {
	accountService, _ := getAccountService()
	id := "0d209d7f-d07a-4542-947f-5885fddddae7"
	version := 0
	err := accountService.DeleteAccount(id, version)
	assert.Empty(t, err, "Error is not empty")
}

func TestErrorDeletingAccount(t *testing.T) {
	accountService, _ := getAccountService()
	id := "0d209d7f-d07a-4542-947f-5885fddddae7"
	version := 0
	err := accountService.DeleteAccount(id, version)
	assert.NotEmpty(t, err, "Error is empty")
}

func getAccountService() (*AccountService, error) {
	accountApi, _ := api.NewAccountApi(hostname)
	return NewAccountService(accountApi)
}
