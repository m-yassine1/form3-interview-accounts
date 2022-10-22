package service

import (
	"fmt"
	"form3-interview-accounts/internal/validation"
	"form3-interview-accounts/model"
)

type AccountOperations interface {
	GetAccounts(filters map[string]string) ([]model.Account, error)
	GetAccount(id string) (*model.Account, error)
	DeleteAccount(id string, version int) error
	CreateAccount(accountBody model.AccountData) (*model.Account, error)
	IsHealthy() error
}

type AccountService struct {
	accountOperations AccountOperations
}

func NewAccountService(accountOperations AccountOperations) (*AccountService, error) {
	if accountOperations == nil {
		return nil, fmt.Errorf("error creating account service, fromApi is nil")
	}

	accountService := AccountService{
		accountOperations: accountOperations,
	}

	return &accountService, nil
}

func (accountService AccountService) GetAccounts(filters map[string]string) ([]model.Account, error) {
	return accountService.accountOperations.GetAccounts(filters)
}

func (accountService AccountService) GetAccount(id string) (*model.Account, error) {
	return accountService.accountOperations.GetAccount(id)
}

func (accountService AccountService) DeleteAccount(id string, version int) error {
	return accountService.accountOperations.DeleteAccount(id, version)
}

func (accountService AccountService) CreateAccount(accountData model.AccountData) (*model.Account, error) {
	err := validation.ValidateAccount(accountData)
	if err != nil {
		return nil, err
	}
	return accountService.accountOperations.CreateAccount(accountData)
}

func (accountService AccountService) IsHealthy() error {
	return accountService.accountOperations.IsHealthy()
}
