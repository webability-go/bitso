package bitso

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type FeesResponse struct {
	Success bool   `json:"success"`
	Error   *Error `json:"error,omitempty"`
	Fees    Fees   `json:"payload,omitempty"`
}

type Fees struct {
	Fees           []Fee       `json:"fees"`
	WithdrawalFees interface{} `json:"withdrawal_fees"`
}

type Fee struct {
	Book            string  `json:"book"`
	TakerFeeDecimal float64 `json:"taker_fee_decimal,string"`
	TakerFeePercent float64 `json:"taker_fee_percent,string"`
	MakerFeeDecimal float64 `json:"maker_fee_decimal,string"`
	MakerFeePercent float64 `json:"maker_fee_percent,string"`
}

func (api *API) Fees() (*Fees, error) {
	err := api.haveKeys()
	if err != nil {
		return nil, err
	}

	data, err := api.askPrivateGet("fees/", "")
	if err != nil {
		return nil, err
	}
	ab := FeesResponse{}
	err = json.Unmarshal(data, &ab)
	if err != nil {
		return nil, err
	}
	// turn fees_withdrawal to float64
	fees := map[string]float64{}
	for id, val := range ab.Fees.WithdrawalFees.(map[string]interface{}) {
		fees[id], _ = strconv.ParseFloat(val.(string), 64)
	}
	ab.Fees.WithdrawalFees = fees

	if api.Devel {
		fmt.Printf("DATA RECEIVED: %+v\n", ab)
	}
	return &ab.Fees, nil
}
