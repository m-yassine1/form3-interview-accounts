package internal

import (
	"fmt"
	"form3-interview-accounts/model"
	"form3-interview-accounts/util"
)

func ValidateAccount(body model.AccountData) error {
	if body.Data.Attributes == nil {
		return fmt.Errorf("invalid body, attributes is missing")
	}

	if body.Data.Attributes.Country == nil {
		return fmt.Errorf("invalid country, country is missing")
	}

	for _, val := range util.GetSupportedCountries() {
		if val == *body.Data.Attributes.Country {
			return nil
		}
	}

	return fmt.Errorf("invalid country %s", *body.Data.Attributes.Country)
}
