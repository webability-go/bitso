package public

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type OrderBookResponse struct {
	Success   bool      `json:"success"`
	Error     Error     `json:"error,omitempty"`
	OrderBook OrderBook `json:"payload,omitempty"`
}

type OrderBook struct {
	Asks      []Order   `json:"asks"`
	Bids      []Order   `json:"bids"`
	UpdatedAt time.Time `json:"updated_at"`
	Sequence  int       `json:"sequence,string"`
}

type Order struct {
	Book   string  `json:"book"`
	Price  float64 `json:"price,string"`
	Amount float64 `json:"amount,string"`
	OID    string  `json:"oid,omitempty"`
}

func (api *API) OrderBook(book string, aggregate bool) (*OrderBook, error) {
	if book == "" {
		return nil, errors.New("There is no book specified")
	}
	extra := ""
	if !aggregate {
		extra = "&aggregate=false"
	}
	data, err := api.ask("order_book/?book="+book+extra, nil)
	if err != nil {
		return nil, err
	}
	ab := OrderBookResponse{}
	err = json.Unmarshal(data, &ab)
	if err != nil {
		return nil, err
	}

	if api.DEVEL {
		fmt.Printf("DATA RECEIVED: %+v\n", ab)
	}
	return &ab.OrderBook, nil
}
