package bitso

import (
	"encoding/json"
	"fmt"
)

type BankCodesResponse struct {
	Success   bool       `json:"success"`
	Error     *Error     `json:"error,omitempty"`
	BankCodes []BankCode `json:"payload,omitempty"`
}

type BankCode struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func (api *API) BankCodes() ([]BankCode, error) {
	err := api.haveKeys()
	if err != nil {
		return nil, err
	}

	data, err := api.askPrivateGet("mx_bank_codes/", "")
	if err != nil {
		return nil, err
	}
	ab := BankCodesResponse{}
	err = json.Unmarshal(data, &ab)
	if err != nil {
		return nil, err
	}

	if api.Devel {
		fmt.Printf("DATA RECEIVED: %+v\n", ab)
	}
	return ab.BankCodes, nil
}
