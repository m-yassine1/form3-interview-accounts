package main

import (
	"account/api"
	"account/model"
	"account/service"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	form3Api, err := api.NewAccountApi("http://localhost:8080")
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	s, err := service.NewAccountService(form3Api)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	/*
		accounts, err := s.GetAccounts()
		if err != nil {
			fmt.Printf("%s", err)
			return
		}

		fmt.Printf("%#v\n", len(accounts))

		id := accounts[0].ID
		account, err := s.GetAccount(id)
		if err != nil {
			fmt.Printf("%s", err)
			return
		}

		fmt.Printf("%#v\n\n\n", account)

		err = s.DeleteAccount(id, 0)
		if err != nil {
			fmt.Printf("%s", err)
			return
		}

		accounts, err = s.GetAccounts()
		if err != nil {
			fmt.Printf("%s", err)
			return
		}
		fmt.Printf("%#v\n\n\n", len(accounts))
	*/
	var version int64 = 0
	names := make([]string, 0)
	country := "GB"
	uuidString := uuid.New()
	createAccount := model.AccountData{
		Data: model.Account{
			ID:             uuidString.String(),
			OrganisationID: uuidString.String(),
			Type:           "accounts",
			Version:        &version,
			Attributes: &model.AccountAttributes{
				Name:    names,
				Country: &country,
			},
		},
	}

	account, err := s.CreateAccount(createAccount)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	fmt.Printf("%#v\n\n\n", account)

	accounts, err := s.GetAccounts()
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	fmt.Printf("%#v\n", len(accounts))
}
