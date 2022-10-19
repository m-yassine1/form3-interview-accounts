package api

import (
	"account/model"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Form3Api struct {
	url string
}

const createAccountPath = "/v1/organisation/accounts"
const getAllAccountsPath = "/v1/organisation/accounts"
const getAccountPath = "/v1/organisation/accounts/%s"
const deleteAccountPath = "/v1/organisation/accounts/%s"
const applicationJsonContentType = "application/json"

func NewForm3Api(url string) (*Form3Api, error) {
	_, err := validUrl(url)
	if err != nil {
		return nil, err
	}

	form3Api := &Form3Api{
		url: fixUrlString(url),
	}

	return form3Api, nil
}

func validUrl(urlString string) (bool, error) {
	if urlString == "" {
		return false, fmt.Errorf("URL string mst not be empty")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
		return false, fmt.Errorf("URL %s is invalid", urlString)
	}

	return true, nil
}

func fixUrlString(url string) string {
	if url[len(url)-1] != '/' {
		url += "/"
	}
	return url
}

func (form3Api Form3Api) GetAccounts() ([]model.Account, error) {
	resp, err := http.Get(form3Api.getUrl(getAllAccountsPath))
	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
		return []model.Account{}, fmt.Errorf("error fetching list of accounts")
	}
	defer resp.Body.Close()

	var data model.AccountsData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
		return []model.Account{}, fmt.Errorf("unable to parse json response for list of accounts")
	}

	return data.Data, nil
}

func (form3Api Form3Api) GetAccount(id string) (*model.Account, error) {
	resp, err := http.Get(form3Api.getUrl(fmt.Sprintf(getAccountPath, id)))
	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
		return nil, fmt.Errorf("error fetching account %s", id)
	}
	defer resp.Body.Close()

	var data model.AccountData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
		return nil, fmt.Errorf("unable to parse json response for account")
	}

	return &data.Data, nil
}

func (form3Api Form3Api) DeleteAccount(id string) error {
	resp, err := http.NewRequest(http.MethodDelete, form3Api.getUrl(fmt.Sprintf(deleteAccountPath, id)), nil)
	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
		return fmt.Errorf("error deleting account %s", id)
	}
	defer resp.Body.Close()

	if resp.Response.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to delete account %s with status code %d", id, resp.Response.StatusCode)
	}

	return nil
}

func (form3Api Form3Api) CreateAccount(account model.AccountData) (*model.Account, error) {
	marshaledData, err := json.Marshal(account)
	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
		return nil, fmt.Errorf("error parsing request account body %#v", account)
	}

	resp, err := http.Post(form3Api.getUrl(createAccountPath), applicationJsonContentType, bytes.NewBuffer(marshaledData))
	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
		return nil, fmt.Errorf("error creating account %#v", account)
	}
	defer resp.Body.Close()

	var data model.AccountData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
		return nil, fmt.Errorf("unable to parse json response for creating account")
	}

	return &data.Data, nil
}

func (form3Api Form3Api) getUrl(path string) string {
	return form3Api.url + path
}
