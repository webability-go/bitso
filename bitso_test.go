package bitso

import (
	"fmt"
	"testing"
)

const (
	// Add your key here
	KEY = "iHofeixtvF"
	// Add your secret key here
	SECRET = "57f5edb89e2afc4b06ad64b2475f0317"
)

func TestBitso(t *testing.T) {
	api := NewAPI(KEY, SECRET)
	books, err := api.AvailableBooks()
	fmt.Println(err)
	fmt.Printf("%+v\n", books)

	tickers, err := api.Tickers()
	fmt.Println(err)
	fmt.Printf("%+v\n", tickers)

	ticker1, err := api.Ticker("")
	fmt.Println(err)
	fmt.Printf("%+v\n", ticker1)

	ticker2, err := api.Ticker("btc_mxn")
	fmt.Println(err)
	fmt.Printf("%+v\n", ticker2)

	orderbook, err := api.OrderBook("btc_mxn", false)
	fmt.Println(err)
	fmt.Printf("%+v\n", orderbook)

	trades, err := api.Trades("btc_mxn", "", "", 10)
	fmt.Println(err)
	fmt.Printf("%+v\n", trades)

	api.Key = KEY
	api.Secret = SECRET
	api.Devel = true

	accountstatus, err := api.AccountStatus()
	fmt.Println(err)
	fmt.Printf("%+v\n", accountstatus)

	balance, err := api.Balance()
	fmt.Println(err)
	fmt.Printf("%+v\n", balance)

	fees, err := api.Fees()
	fmt.Println(err)
	fmt.Printf("%+v\n", fees)

	bankcodes, err := api.BankCodes()
	fmt.Println(err)
	fmt.Printf("%+v\n", bankcodes)

	usertrades, err := api.UserTrades(nil)
	fmt.Println(err)
	fmt.Printf("%+v\n", usertrades)

	//	fmt.Println(api)
}
