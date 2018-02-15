package kite

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	version         = "3"
	rootAPIEndpoint = "https://api.kite.trade"
	loginPath       = "/connect/login"
)

// Client provides the functionality required for the consumer to access all of the Kite Connect Endpoints
type Client interface {
	Login() (string, error)
}

type kite struct {
	httpClient http.Client
	apiKey     string
}

func (k kite) Login() (string, error) {
	loginEndpoint := rootAPIEndpoint + loginPath
	req, err := http.NewRequest("GET", loginEndpoint, nil)
	if err != nil {
		return "", err
	}
	q := req.URL.Query()
	q.Add("v", "3")
	q.Add("api_key", k.apiKey)
	fmt.Println("Query Params: " + q.Encode())
	resp, err := k.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respReader := resp.Body
	body, err := ioutil.ReadAll(respReader)
	if err != nil {
		return "", err
	}
	fmt.Println(string(body))
	return "ACCESS TOKEN", nil
}

// NewClient creates a Kite Connect Client that can be used by the consumer to consume the Kite Connect Endpoints
func NewClient(apiKey string) Client {
	var client = kite{apiKey: apiKey}
	return client
}
