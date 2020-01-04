package bitso

import (
	"encoding/json"
	"fmt"
)

type BalanceResponse struct {
	Success bool  `json:"success"`
	Error   Error `json:"error,omitempty"`
	Payload struct {
		Balances Balances `json:"balances,omitempty"`
	} `json:"payload,omitempty"`
}

type Balances []Balance

type Balance struct {
	Currency  string  `json:"currency"`
	Total     float64 `json:"total,string"`
	Locked    float64 `json:"locked,string"`
	Available float64 `json:"available,string"`
}

func (api *API) Balance() (Balances, error) {
	err := api.haveKeys()
	if err != nil {
		return nil, err
	}

	data, err := api.askPrivateGet("balance/", "")
	if err != nil {
		return nil, err
	}
	ab := BalanceResponse{}
	err = json.Unmarshal(data, &ab)
	if err != nil {
		return nil, err
	}

	if api.Devel {
		fmt.Printf("DATA RECEIVED: %+v\n", ab)
	}
	return ab.Payload.Balances, nil
}
