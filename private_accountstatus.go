package bitso

import (
	"encoding/json"
	"fmt"
)

type AccountStatusResponse struct {
	Success     bool        `json:"success"`
	Error       Error       `json:"error,omitempty"`
	AccountData AccountData `json:"payload,omitempty"`
}

type AccountData struct {
	ClientID              string  `json:"client_id"`
	FirstName             string  `json:"first_name"`
	LastName              string  `json:"last_name"`
	Status                string  `json:"status"`
	DailyLimit            float64 `json:"daily_limit,string"`
	MonthlyLimit          float64 `json:"monthly_limit,string"`
	DailyRemaining        float64 `json:"daily_remaining,string"`
	MonthlyRemaining      float64 `json:"monthly_remaining,string"`
	CashDepositAllowance  float64 `json:"cash_deposit_allowance,string"`
	CellphoneNumber       string  `json:"cellphone_number"`
	CellphoneNumberStored string  `json:"cellphone_number_stored"`
	EmailStored           string  `json:"email_stored"`
	OfficialID            string  `json:"official_id"`
	ProofOfResidency      string  `json:"proof_of_residency"`
	SignedContract        string  `json:"signed_contract"`
	OriginOfFunds         string  `json:"origin_of_funds"`
}

func (api *API) AccountStatus() (*AccountData, error) {
	err := api.haveKeys()
	if err != nil {
		return nil, err
	}

	data, err := api.askPrivateGet("account_status/", "")
	if err != nil {
		return nil, err
	}
	ab := AccountStatusResponse{}
	err = json.Unmarshal(data, &ab)
	if err != nil {
		return nil, err
	}

	if api.Devel {
		fmt.Printf("DATA RECEIVED: %+v\n", ab)
	}
	return &ab.AccountData, nil
}
