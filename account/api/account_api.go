package api

import (
	"account/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type AccountApi struct {
	url string
}

const createAccountPath = "/v1/organisation/accounts"
const getAllAccountsPath = "/v1/organisation/accounts"
const getAccountPath = "/v1/organisation/accounts/%s"
const deleteAccountPath = "/v1/organisation/accounts/%s?version=%d"
const applicationJsonContentType = "application/json"

func NewAccountApi(url string) (*AccountApi, error) {
	_, err := validUrl(url)
	if err != nil {
		return nil, err
	}

	form3Api := &AccountApi{
		url: removeSlashEndOfHostname(url),
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

func removeSlashEndOfHostname(url string) string {
	return strings.TrimSuffix(url, "/")
}

func (form3Api AccountApi) GetAccounts() ([]model.Account, error) {
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

func (form3Api AccountApi) GetAccount(id string) (*model.Account, error) {
	resp, err := http.Get(form3Api.getUrl(fmt.Sprintf(getAccountPath, id)))
	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
		return nil, fmt.Errorf("error fetching account %s", id)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bytes, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("error fetching account %s with status %d response: %s", id, resp.StatusCode, string(bytes))
	}

	var data model.AccountData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
		return nil, fmt.Errorf("unable to parse json response for account")
	}

	return &data.Data, nil
}

func (form3Api AccountApi) DeleteAccount(id string, version int) error {
	req, err := http.NewRequest(http.MethodDelete, form3Api.getUrl(fmt.Sprintf(deleteAccountPath, id, version)), nil)
	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
		return fmt.Errorf("error creating delete account %s", id)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil && (resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK) {
		bytes, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("failed to delete account %s with status code %d response: %s", id, resp.StatusCode, string(bytes))
	}

	return nil
}

func (form3Api AccountApi) CreateAccount(account model.AccountData) (*model.Account, error) {
	marshaledData, err := json.Marshal(account)
	if err != nil {
		fmt.Printf("\n\n\n%T\n%s\n%#v\n", err, err, err)
		return nil, fmt.Errorf("error parsing request account body %#v", account)
	}

	resp, err := http.Post(form3Api.getUrl(createAccountPath), applicationJsonContentType, bytes.NewBuffer(marshaledData))
	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
		return nil, fmt.Errorf("error creating account %#v", account)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		bytes, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("faile to create account with status %d response: %s", resp.StatusCode, string(bytes))
	}

	var data model.AccountData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
		return nil, fmt.Errorf("unable to parse json response for creating account")
	}

	return &data.Data, nil
}

func (form3Api AccountApi) getUrl(path string) string {
	return form3Api.url + path
}
