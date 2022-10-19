package service

import (
	"account/api"
	"account/model"
	"fmt"
)

type AccountService struct {
	accountApi *api.AccountApi
}

func NewAccountService(form3Api *api.AccountApi) (*AccountService, error) {
	if form3Api == nil {
		return nil, fmt.Errorf("error creating account service, fromApi is nil")
	}

	accountService := AccountService{
		accountApi: form3Api,
	}

	return &accountService, nil
}

func (accountService AccountService) GetAccounts() ([]model.Account, error) {
	return accountService.accountApi.GetAccounts()
}

func (accountService AccountService) GetAccount(id string) (*model.Account, error) {
	return accountService.accountApi.GetAccount(id)
}

func (accountService AccountService) DeleteAccount(id string, version int) error {
	return accountService.accountApi.DeleteAccount(id, version)
}

func (accountService AccountService) CreateAccount(accountBody model.AccountData) (*model.Account, error) {
	return accountService.accountApi.CreateAccount(accountBody)
}
