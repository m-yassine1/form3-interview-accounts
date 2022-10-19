package main

import (
	"account/service"
	"fmt"
)

func main() {
	fmt.Println("Hello, Modules!")
	s, _ := service.NewAccountService("http://localhost:8080")
	accounts, err := s.GetAccounts()
	if err != nil {
		fmt.Println("URL string mst not be empty")
		return
	}
	fmt.Printf("%#v", len(accounts))
}
