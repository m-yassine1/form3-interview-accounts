package main

import (
	"account/api"
	"account/service"
	"fmt"
)

func main() {
	fmt.Println("Hello, Modules!")
	form3Api, err := api.NewForm3Api("http://localhost:8080")
	if err != nil {
		fmt.Println("Error creating form3 api")
		return
	}
	s, err := service.NewAccountService(form3Api)
	if err != nil {
		fmt.Println("Error creating account service")
		return
	}
	accounts, err := s.GetAccounts()
	if err != nil {
		fmt.Println("URL string mst not be empty")
		return
	}
	fmt.Printf("%#v", len(accounts))
}
