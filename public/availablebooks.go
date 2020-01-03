package public

import (
	"encoding/json"
	"fmt"
)

type AvailableBookResponse struct {
	Success bool  `json:"success"`
	Error   Error `json:"error,omitempty"`
	Books   Books `json:"payload,omitempty"`
}

type Books []Book

type Book struct {
	Book          string  `json:"book"`
	MinimumAmount float64 `json:"minimum_amount,string"`
	MaximumAmount float64 `json:"maximum_amount,string"`
	MinimumPrice  float64 `json:"minimum_price,string"`
	MaximumPrice  float64 `json:"maximum_price,string"`
	MinimumValue  float64 `json:"minimum_value,string"`
	MaximumValue  float64 `json:"maximum_value,string"`
}

func (api *API) AvailableBooks() (Books, error) {
	data, err := api.ask("available_books/", nil)
	if err != nil {
		return nil, err
	}
	ab := AvailableBookResponse{}
	err = json.Unmarshal(data, &ab)
	if err != nil {
		return nil, err
	}

	if api.DEVEL {
		fmt.Printf("DATA RECEIVED: %+v\n", ab)
	}
	return ab.Books, nil
}
