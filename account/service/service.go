package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type AccountData struct {
	Data []Account `json:"data,omitempty"`
}

type Account struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	ID             string             `json:"id,omitempty"`
	OrganisationID string             `json:"organisation_id,omitempty"`
	Type           string             `json:"type,omitempty"`
	Version        *int64             `json:"version,omitempty"`
}

type AccountAttributes struct {
	AccountClassification   *string  `json:"account_classification,omitempty"`
	AccountMatchingOptOut   *bool    `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	AlternativeNames        []string `json:"alternative_names,omitempty"`
	BankID                  string   `json:"bank_id,omitempty"`
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Country                 *string  `json:"country,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	JointAccount            *bool    `json:"joint_account,omitempty"`
	Name                    []string `json:"name,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	Status                  *string  `json:"status,omitempty"`
	Switched                *bool    `json:"switched,omitempty"`
}

type AccountService struct {
	url string
}

var accountService *AccountService = nil

func NewAccountService(url string) (*AccountService, error) {
	if accountService != nil {
		return accountService, nil
	}

	_, err := validUrl(url)
	if err != nil {
		return nil, err
	}

	accountService = &AccountService{
		url: fixUrlString(url),
	}

	return accountService, nil
}

func (accountService *AccountService) UpdateUrl(url string) error {
	_, err := validUrl(url)
	if err != nil {
		return err
	}

	accountService.url = fixUrlString(url)
	return nil
}

func validUrl(urlString string) (bool, error) {
	if urlString == "" {
		return false, fmt.Errorf("URL string mst not be empty")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
		return false, fmt.Errorf("URL " + urlString + " is invalid")
	}

	return true, nil
}

func fixUrlString(url string) string {
	if url[len(url)-1] != '/' {
		url += "/"
	}
	return url
}

func (accountService AccountService) GetAccounts() ([]Account, error) {
	url := accountService.url
	resp, err := http.Get(url + "/v1/organisation/accounts")
	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
		return []Account{}, fmt.Errorf("error fetching list of accounts")
	}
	defer resp.Body.Close()

	var data AccountData
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&data)
	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
		return []Account{}, fmt.Errorf("unable to parse json response")
	}

	return data.Data, nil
}
