package bitso

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type TradesResponse struct {
	Success bool   `json:"success"`
	Error   Error  `json:"error,omitempty"`
	Trades  Trades `json:"payload,omitempty"`
}

type Trades []Trade

type Trade struct {
	Book       string    `json:"book"`
	CreatedAt  time.Time `json:"created_at"`
	Amount     float64   `json:"amount,string"`
	MarkerSite string    `json:"marker_side"`
	Price      float64   `json:"price,string"`
	TID        int       `json:"tid"`
}

func (api *API) Trades(book string, marker string, sort string, limit int) (*Trades, error) {
	if book == "" {
		return nil, errors.New("There is no book specified")
	}
	extra := ""
	if marker != "" {
		extra += "&marker=" + marker
	}
	if sort == "" || sort == "asc" || sort == "desc" {
		if sort != "" {
			extra += "&sort=" + sort
		}
	} else {
		return nil, errors.New("Sort has a wrong value (asc or desc): " + sort)
	}
	slimit := strconv.Itoa(limit)
	if limit >= 0 && limit < 101 {
		if limit > 0 {
			extra += "&limit=" + slimit
		}
	} else {
		return nil, errors.New("Limit has a wrong value (1 to 100): " + slimit)
	}
	data, err := api.ask("trades/?book="+book+extra, nil)
	if err != nil {
		return nil, err
	}
	ab := TradesResponse{}
	// There is a but on trades: dates comes as "+0000" and not "+00:00"
	data = []byte(strings.Replace(string(data), "+0000", "+00:00", -1))
	err = json.Unmarshal(data, &ab)
	if err != nil {
		return nil, err
	}

	if api.Devel {
		fmt.Printf("DATA RECEIVED: %+v\n", ab)
	}
	return &ab.Trades, nil
}
