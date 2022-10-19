package service

import (
	"account/api"
	"account/model"
	"fmt"
)

type AccountService struct {
	form3Api *api.Form3Api
}

var accountService *AccountService = nil
var form3Api *api.Form3Api = nil

func NewAccountService(form3Api *api.Form3Api) (*AccountService, error) {
	if accountService != nil {
		return accountService, nil
	}

	if form3Api == nil {
		return nil, fmt.Errorf("form3Api service is not initialized")
	}

	accountService = &AccountService{
		form3Api: form3Api,
	}

	return accountService, nil
}

func (accountService AccountService) GetAccounts() ([]model.Account, error) {
	accounts, err := form3Api.GetAccounts()
	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
		return []model.Account{}, err
	}
	return accounts, nil
}

func (accountService AccountService) GetAccount(id string) (*model.Account, error) {
	account, err := form3Api.GetAccount(id)
	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
		return nil, err
	}

	return account, nil
}

func (accountService AccountService) DeleteAccount(id string) error {
	err := form3Api.DeleteAccount(id)
	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
		return err
	}

	return nil
}

func (accountService AccountService) CreateAccount(accountBody model.AccountData) (*model.Account, error) {
	account, err := form3Api.CreateAccount(accountBody)
	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
		return nil, err
	}

	return account, nil
}
