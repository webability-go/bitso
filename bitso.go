package bitso

const VERSION = "0.1.1"

type API struct {
	Key            string
	Secret         string
	Devel          bool
	HaveTransfer   bool
	HaveRemittance bool
}

func NewAPI(key string, secret string) *API {
	return &API{
		Key:    key,
		Secret: secret,
	}
}

/*
func NewWebSocketService() *websocket.Service {
	return &websocket.Service{}
}
*/
