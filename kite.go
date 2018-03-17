package kite

import (
	"net/http"
)

const (
	version         = "3"
	rootAPIEndpoint = "https://api.kite.trade"
	loginPath       = "/connect/login"
	accessTokenPath = "/session/token"
	marginsPath     = "/user/margins"
)

// Client provides the functionality required for the consumer to access all of the Kite Connect Endpoints
type Client struct {
	RootURL     string
	APIKey      string
	AccessToken string
	HTTPClient  *http.Client
}

func (c Client) getRootURL() string {
	if c.RootURL != "" {
		return c.RootURL
	}
	return rootAPIEndpoint
}
