package bitso

import (
	"github.com/webability-go/bitso/private"
	"github.com/webability-go/bitso/public"
	"github.com/webability-go/bitso/remittance"
	"github.com/webability-go/bitso/transfer"
	"github.com/webability-go/bitso/webhook"
	"github.com/webability-go/bitso/websocket"
)

func NewPublicAPI() *public.API {
	return &public.API{}
}

func NewPrivateAPI() *private.API {
	return &private.API{}
}

func NewRemittanceAPI() *remittance.API {
	return &remittance.API{}
}

func NewTransferAPI() *transfer.API {
	return &transfer.API{}
}

func NewWebHookAPI() *webhook.API {
	return &webhook.API{}
}

func NewWebSocketService() *websocket.Service {
	return &websocket.Service{}
}