package bitso

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	BASEURL = "https://api.bitso.com/v3/"
)

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code,string"`
}

func (api *API) ask(service string, params map[string]string) ([]byte, error) {

	hc := http.Client{}
	// only GET for public queries
	req, err := http.NewRequest(http.MethodGet, BASEURL+service, nil)
	if err != nil {
		return nil, err
	}
	resp, err := hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if api.Devel {
		fmt.Println("RAW DATA:", string(data))
	}
	return data, nil
}

func (api *API) askPrivateGet(service string, query string) ([]byte, error) {

	hc := http.Client{}
	req, err := http.NewRequest(http.MethodGet, BASEURL+service+query, nil)
	if err != nil {
		return nil, err
	}
	auth := api.signature(http.MethodGet, "/v3/"+service+query, "")
	if api.Devel {
		fmt.Println("Authorization:", auth)
	}
	req.Header.Add("Authorization", auth)

	resp, err := hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if api.Devel {
		fmt.Println("RAW DATA:", string(data))
	}
	return data, nil
}

func (api *API) askPrivatePost(service string, query string, body string) ([]byte, error) {

	hc := http.Client{}
	req, err := http.NewRequest(http.MethodPost, BASEURL+service+query, bytes.NewBufferString(body))
	if err != nil {
		return nil, err
	}
	auth := api.signature(http.MethodPost, "/v3/"+service+query, body)
	if api.Devel {
		fmt.Println("Auth: ", auth)
	}
	req.Header.Add("Authorization", auth)
	req.Header.Set("Content-Type", "application/json")

	resp, err := hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if api.Devel {
		fmt.Println("RAW DATA:", string(data))
	}
	return data, nil
}

func (api *API) askPrivateDelete(service string, query string, body string) ([]byte, error) {

	hc := http.Client{}
	req, err := http.NewRequest(http.MethodDelete, BASEURL+service+query, bytes.NewBufferString(body))
	if err != nil {
		return nil, err
	}
	auth := api.signature(http.MethodDelete, "/v3/"+service+query, "")
	if api.Devel {
		fmt.Println("Authorization: ", auth)
	}
	req.Header.Add("Authorization", auth)
	req.Header.Set("Content-Type", "application/json")

	resp, err := hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if api.Devel {
		fmt.Println("RAW DATA:", string(data))
	}
	return data, nil
}

func (api *API) signature(method string, url string, body string) string {

	nonce := strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
	signstring := nonce + method + url + body
	signer := hmac.New(sha256.New, []byte(api.Secret))
	signer.Write([]byte(signstring))
	signature := hex.EncodeToString(signer.Sum(nil))
	return fmt.Sprintf("Bitso %s:%s:%s", api.Key, nonce, signature)
}

func (api *API) haveKeys() error {
	if api.Key == "" || api.Secret == "" {
		return errors.New("There is no keys to access to private APIs")
	}
	return nil
}

func (api *API) haveTransfer() error {
	err := api.haveKeys()
	if err != nil {
		return err
	}
	if !api.HaveTransfer {
		return errors.New("There is no transfer flag set on the API")
	}
	return nil
}

func (api *API) haveRemittance() error {
	err := api.haveKeys()
	if err != nil {
		return err
	}
	if !api.HaveRemittance {
		return errors.New("There is no remittance flag set on the API")
	}
	return nil
}
