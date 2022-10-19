package service

import (
	"account/api"
	"account/model"
	"fmt"
)

type AccountService struct {
	form3Api *api.Form3Api
}

func NewAccountService(form3Api *api.Form3Api) (*AccountService, error) {
	if form3Api == nil {
		return nil, fmt.Errorf("error creating account service, fromApi is nil")
	}

	accountService := AccountService{
		form3Api: form3Api,
	}

	return &accountService, nil
}

func (accountService AccountService) GetAccounts() ([]model.Account, error) {
	return accountService.form3Api.GetAccounts()
}

func (accountService AccountService) GetAccount(id string) (*model.Account, error) {
	return accountService.form3Api.GetAccount(id)
}

func (accountService AccountService) DeleteAccount(id string) error {
	return accountService.form3Api.DeleteAccount(id)
}

func (accountService AccountService) CreateAccount(accountBody model.AccountData) (*model.Account, error) {
	return accountService.form3Api.CreateAccount(accountBody)
}
