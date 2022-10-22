package service

import (
	"fmt"
	"form3-interview-accounts/api"
	"form3-interview-accounts/model"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var hostname = "http://localhost:8080"

func TestMain(m *testing.M) {
	accountService, _ := getAccountService()
	for accountService.IsHealthy() != nil {
		fmt.Println("service is not yet up. Sleeping for 5 seconds")
		time.Sleep(5 * time.Second)
	}

	accounts, err := accountService.GetAccounts(nil)
	if err != nil {
		fmt.Printf("unable to fetch accounts. Error: %s\n", err)
		os.Exit(2)
	}
	if len(accounts) == 0 {
		for _, value := range accounts {
			err = accountService.DeleteAccount(value.ID, int(*value.Version))
			if err != nil {
				fmt.Printf("unable to delete account %s. Error: %s\n", value.ID, err)
			}
		}
	}
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestServiceCreation(t *testing.T) {
	accountService, err := NewAccountService(nil)
	assert.NotEmpty(t, err, "Error is empty")
	assert.Empty(t, accountService, "Account Service is not empty")

	accountService, err = getAccountService()
	assert.Empty(t, err, "Error is not empty for valid hostname")
	assert.NotEmpty(t, accountService, "Account API is empty for valid hostname")
}

func TestFailedGetAccounts(t *testing.T) {
	accountService, _ := getAccountService()
	accounts, err := accountService.GetAccounts(nil)
	assert.Empty(t, err, "Error is not empty")
	assert.Empty(t, accounts, "Accounts list is empty")
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

func TestCreateSecondAccount(t *testing.T) {
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

func TestGetAccounts(t *testing.T) {
	accountService, _ := getAccountService()
	accounts, err := accountService.GetAccounts(nil)
	assert.Empty(t, err, "Error is not empty")
	assert.NotEmpty(t, accounts, "Accounts list is empty")
	assert.Equal(t, len(accounts), 2)
}

func TestGetAccount(t *testing.T) {
	accountService, _ := getAccountService()
	account, err := getAccountId(accountService)
	assert.Empty(t, err, "Error is not empty")
	accountFetched, err := accountService.GetAccount(account.ID)
	assert.Empty(t, err, "Error is not empty")
	assert.NotEmpty(t, accountFetched, "Account is empty")
	assert.Equal(t, accountFetched.ID, account.ID)
}

func TestGetNotFoundAccount(t *testing.T) {
	accountService, _ := getAccountService()
	account, err := getAccountId(accountService)
	assert.Empty(t, err, "Error is not empty")
	accountFetched, err := accountService.GetAccount(account.ID + "r")
	assert.NotEmpty(t, err, "Error is empty")
	assert.Empty(t, accountFetched, "Account is empty")
}

func TestSuccessfulDeletingAccount(t *testing.T) {
	accountService, _ := getAccountService()
	account, err := getAccountId(accountService)
	assert.Empty(t, err, "Error is not empty")
	version := 0
	err = accountService.DeleteAccount(account.ID, version)
	assert.Empty(t, err, "Error is not empty")
}

func TestErrorDeletingAccount(t *testing.T) {
	accountService, _ := getAccountService()
	account, err := getAccountId(accountService)
	assert.Empty(t, err, "Error is not empty")
	version := 0
	err = accountService.DeleteAccount(account.ID, version+1)
	assert.Empty(t, err, "Error is not empty")
}

func getAccountService() (*AccountService, error) {
	accountApi, _ := api.NewAccountApi(hostname)
	return NewAccountService(accountApi)
}

func getAccountId(accountService *AccountService) (model.Account, error) {
	accounts, err := accountService.GetAccounts(nil)
	return accounts[0], err
}
