package bitso

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type UserTradesResponse struct {
	Success    bool        `json:"success"`
	Error      *Error      `json:"error,omitempty"`
	UserTrades []UserTrade `json:"payload,omitempty"`
}

type UserTrade struct {
	Tid          int       `json:"tid,string"`
	OID          string    `json:"iod"`
	Book         string    `json:"book"`
	Side         string    `json:"side"`
	Price        float64   `json:"price,string"`
	Major        float64   `json:"major,string"`
	Minor        float64   `json:"minor,string"`
	FeesAmount   float64   `json:"fees_amount,string"`
	FeesCurrency string    `json:"fees_currency"`
	CreatedAt    time.Time `json:"created_at"`
	ClientID     string    `json:"client_id"`
}

func (api *API) UserTrades(tid []int) ([]UserTrade, error) {
	err := api.haveKeys()
	if err != nil {
		return nil, err
	}

	data, err := api.askPrivateGet("user_trades/", "")
	if err != nil {
		return nil, err
	}
	ab := UserTradesResponse{}
	// There is a bug on trades: dates come as "+0000" and not "+00:00"
	data = []byte(strings.Replace(string(data), "+0000", "+00:00", -1))
	err = json.Unmarshal(data, &ab)
	if err != nil {
		return nil, err
	}

	if api.Devel {
		fmt.Printf("DATA RECEIVED: %+v\n", ab)
	}
	return ab.UserTrades, nil
}
