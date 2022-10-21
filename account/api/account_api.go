package api

import (
	"account/model"
	"account/util"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type AccountApi struct {
	url string
}

const healthyPath = "/v1/health"
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
		return false, fmt.Errorf("URL %s is invalid. Error: %s", urlString, err)
	}

	return true, nil
}

func removeSlashEndOfHostname(url string) string {
	return strings.TrimSuffix(url, "/")
}

func (form3Api AccountApi) IsHealthy() error {
	resp, err := http.Get(form3Api.getUrl(healthyPath, nil))
	if err != nil {
		return fmt.Errorf("error checking healthy status. Error: %s", err)
	}
	defer resp.Body.Close()

	var data model.HealhtyData
	err = util.FromJsonToModel(resp.Body, &data)
	if err != nil {
		return fmt.Errorf("unable to parse json response for healthy status. Error: %s", err)
	} else if data.Status != "up" {
		return fmt.Errorf("healhthy status is not up, but %s", data.Status)
	}
	return nil
}

func (form3Api AccountApi) GetAccounts(filters map[string]string) ([]model.Account, error) {
	resp, err := http.Get(form3Api.getUrl(getAllAccountsPath, filters))
	if err != nil {
		return []model.Account{}, fmt.Errorf("error fetching list of accounts. Error: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("error fetching accounts with status %d response: %s", resp.StatusCode, string(bytes))
	}

	var data model.AccountsData
	err = util.FromJsonToModel(resp.Body, &data)
	if err != nil {
		return []model.Account{}, fmt.Errorf("unable to parse json response for list of accounts. Error: %s", err)
	}

	return data.Data, nil
}

func (form3Api AccountApi) GetAccount(id string) (*model.Account, error) {
	resp, err := http.Get(form3Api.getUrl(fmt.Sprintf(getAccountPath, id), nil))
	if err != nil {
		return nil, fmt.Errorf("error fetching account %s. Error: %s", id, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("error fetching account %s with status %d response: %s", id, resp.StatusCode, string(bytes))
	}

	var data model.AccountData
	err = util.FromJsonToModel(resp.Body, &data)
	if err != nil {
		return nil, fmt.Errorf("unable to parse json response for account. Error: %s", err)
	}

	return &data.Data, nil
}

func (form3Api AccountApi) DeleteAccount(id string, version int) error {
	req, err := http.NewRequest(http.MethodDelete, form3Api.getUrl(fmt.Sprintf(deleteAccountPath, id, version), nil), nil)
	if err != nil {
		return fmt.Errorf("error creating delete account %s. Error: %s", id, err)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return fmt.Errorf("failed to delete account %s error: %s", id, err)
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		bytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to delete account %s with status code %d response: %s", id, resp.StatusCode, string(bytes))
	}

	return nil
}

func (form3Api AccountApi) CreateAccount(account model.AccountData) (*model.Account, error) {
	marshaledData, err := json.Marshal(account)
	if err != nil {
		return nil, fmt.Errorf("error parsing request account body %#v. Error: %s", account, err)
	}

	resp, err := http.Post(form3Api.getUrl(createAccountPath, nil), applicationJsonContentType, bytes.NewBuffer(marshaledData))
	if err != nil {
		return nil, fmt.Errorf("error creating account %#v. Error: %s", account, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		bytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("faile to create account with status %d response: %s", resp.StatusCode, string(bytes))
	}

	var data model.AccountData
	err = util.FromJsonToModel(resp.Body, &data)
	if err != nil {
		return nil, fmt.Errorf("unable to parse json response for creating account. Error: %s", err)
	}

	return &data.Data, nil
}

func (form3Api AccountApi) getUrl(path string, filters map[string]string) string {
	return util.BuildUrl(form3Api.url, path, filters)
}
