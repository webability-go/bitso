package bitso

/*
type WithdrawResponse struct {
	Success bool `json:"success"`
	Payload struct {
		Wid       string    `json:"wid"`
		Status    string    `json:"status"`
		CreatedAt time.Time `json:"created_at"`
		Currency  string    `json:"currency"`
		Method    string    `json:"method"`
		Amount    string    `json:"amount"`
		Details   struct {
			WithdrawalAddress string      `json:"withdrawal_address"`
			TxHash            interface{} `json:"tx_hash"`
		} `json:"details"`
	} `json:"payload"`
}

type BooksResponse struct {
	Success bool           `json:"success"`
	Payload []BookResponse `json:"payload"`
}

type BookResponse struct {
	Book          string `json:"book"`
	MinimumAmount string `json:"minimum_amount"`
	MaximumAmount string `json:"maximum_amount"`
	MinimumPrice  string `json:"minimum_price"`
	MaximumPrice  string `json:"maximum_price"`
	MinimumValue  string `json:"minimum_value"`
	MaximumValue  string `json:"maximum_value"`
}

type AccountInfoResponse struct {
	Success bool    `json:"success"`
	Payload Account `json:"payload"`
}

type Account struct {
	ClientID              string `json:"client_id"`
	FirstName             string `json:"first_name"`
	LastName              string `json:"last_name"`
	Status                string `json:"status"`
	DailyLimit            string `json:"daily_limit"`
	MonthlyLimit          string `json:"monthly_limit"`
	DailyRemaining        string `json:"daily_remaining"`
	MonthlyRemaining      string `json:"monthly_remaining"`
	CashDepositAllowance  string `json:"cash_deposit_allowance"`
	CellphoneNumber       string `json:"cellphone_number"`
	CellphoneNumberStored string `json:"cellphone_number_stored"`
	EmailStored           string `json:"email_stored"`
	OfficialID            string `json:"official_id"`
	ProofOfResidency      string `json:"proof_of_residency"`
	SignedContract        string `json:"signed_contract"`
	OriginOfFunds         string `json:"origin_of_funds"`
}

type PlacedOrderResponse struct {
	Success bool `json:"success"`
	Payload struct {
		Oid string `json:"oid"`
	} `json:"payload"`
}

type DestinationResponse struct {
	Success bool `json:"success"`
	Payload struct {
		AccountIdentifierName string `json:"account_identifier_name"`
		AccountIdentifier     string `json:"account_identifier"`
	} `json:"payload"`
}

type LookUpOrdersResponse struct {
	Success bool `json:"success"`
	Payload []LookUpOrderResponse `json:"payload"`
}
type LookUpOrderResponse struct{
	Book           string    `json:"book"`
	OriginalAmount string    `json:"original_amount"`
	UnfilledAmount float64    `json:"unfilled_amount"`
	OriginalValue  string    `json:"original_value"`
	CreatedAt      string 	 `json:"created_at"`
	UpdatedAt      string 	 `json:"updated_at"`
	Price          string    `json:"price"`
	Oid            string    `json:"oid"`
	Side           string    `json:"side"`
	Status         string    `json:"status"`
	Type           string    `json:"type"`
}


type UserTradesResponse struct {
	Success bool `json:"success"`
	Payload []UserTradeResponse`json:"payload"`
}

type UserTradeResponse struct{
	Book         string    	`json:"book"`
	Major        string    	`json:"major"`
	CreatedAt    string 	`json:"created_at"`
	Minor        string    	`json:"minor"`
	FeesAmount   string    	`json:"fees_amount"`
	FeesCurrency string    	`json:"fees_currency"`
	Price        string    	`json:"price"`
	Tid          string		`json:"tid"`
	Oid          string    	`json:"oid"`
	Side         string    	`json:"side"`
}
*/
