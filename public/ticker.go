package public

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type TickersResponse struct {
	Success bool    `json:"success"`
	Error   Error   `json:"error,omitempty"`
	Tickers Tickers `json:"payload,omitempty"`
}

type TickerResponse struct {
	Success bool   `json:"success"`
	Error   Error  `json:"error,omitempty"`
	Ticker  Ticker `json:"payload,omitempty"`
}

type Tickers []Ticker

type Ticker struct {
	Book      string    `json:"book"`
	Volume    float64   `json:"volume,string"`
	High      float64   `json:"high,string"`
	Last      float64   `json:"last,string"`
	Low       float64   `json:"low,string"`
	Vwap      float64   `json:"vwap,string"`
	Ask       float64   `json:"ask,string"`
	Bid       float64   `json:"bid,string"`
	CreatedAt time.Time `json:"created_at"`
}

func (api *API) Tickers() (Tickers, error) {
	data, err := api.ask("ticker/", nil)
	if err != nil {
		return nil, err
	}
	ab := TickersResponse{}
	err = json.Unmarshal(data, &ab)
	if err != nil {
		return nil, err
	}

	if api.DEVEL {
		fmt.Printf("DATA RECEIVED: %+v\n", ab)
	}
	return ab.Tickers, nil
}

func (api *API) Ticker(book string) (*Ticker, error) {
	if book == "" {
		return nil, errors.New("There is no book specified")
	}
	data, err := api.ask("ticker/?book="+book, nil)
	if err != nil {
		return nil, err
	}
	ab := TickerResponse{}
	err = json.Unmarshal(data, &ab)
	if err != nil {
		return nil, err
	}

	if api.DEVEL {
		fmt.Printf("DATA RECEIVED: %+v\n", ab)
	}
	return &ab.Ticker, nil
}
