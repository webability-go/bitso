package bitso

import (
	"fmt"
	"testing"
)

func TestBitso(t *testing.T) {
	b1 := NewPublicAPI()
	b2 := NewPrivateAPI()
	b3 := NewRemittanceAPI()
	b4 := NewTransferAPI()
	b5 := NewWebHookAPI()
	b6 := NewWebSocketService()

	fmt.Println(b1, b2, b3, b4, b5, b6)
}
