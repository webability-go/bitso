package bitso

import (
	"fmt"
	"testing"
)

func TestBitso(t *testing.T) {
	b1 := NewPublicAPI()
	books, err := b1.AvailableBooks()
	fmt.Println(err)
	fmt.Printf("%+v\n", books)

	tickers, err := b1.Tickers()
	fmt.Println(err)
	fmt.Printf("%+v\n", tickers)

	ticker1, err := b1.Ticker("")
	fmt.Println(err)
	fmt.Printf("%+v\n", ticker1)

	ticker2, err := b1.Ticker("btc_mxn")
	fmt.Println(err)
	fmt.Printf("%+v\n", ticker2)

	orderbook, err := b1.OrderBook("btc_mxn", false)
	fmt.Println(err)
	fmt.Printf("%+v\n", orderbook)

	trades, err := b1.Trades("btc_mxn", "", "", 10)
	fmt.Println(err)
	fmt.Printf("%+v\n", trades)

	b2 := NewPrivateAPI()
	b3 := NewRemittanceAPI()
	b4 := NewTransferAPI()
	b5 := NewWebHookAPI()
	b6 := NewWebSocketService()

	fmt.Println(b1, b2, b3, b4, b5, b6)
}
