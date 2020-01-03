package public

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	BASEURL = "https://api.bitso.com/v3/"
)

// API structure
type API struct {
	DEVEL bool
}

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code,string"`
}

// NewAPI will creates a container to access the public REST API
func NewAPI(devel bool) *API {
	return &API{
		DEVEL: devel,
	}
}

func (api *API) ask(service string, params map[string]string) ([]byte, error) {

	hc := http.Client{}
	req, err := http.NewRequest("GET", BASEURL+service, nil)
	if err != nil {
		return nil, err
	}
	resp, err := hc.Do(req)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	if api.DEVEL {
		fmt.Println("RAW DATA:", string(data))
	}
	return data, nil
}
