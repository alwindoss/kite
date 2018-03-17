package kite

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

// Margins is the response retuned from the kite server
type Margins struct {
	Status string `json:"status"`
	Data   struct {
		Equity struct {
			Enabled   bool    `json:"enabled"`
			Net       float64 `json:"net"`
			Available struct {
				AdhocMargin   int `json:"adhoc_margin"`
				Cash          int `json:"cash"`
				Collateral    int `json:"collateral"`
				IntradayPayin int `json:"intraday_payin"`
			} `json:"available"`
			Utilised struct {
				Debits        float64 `json:"debits"`
				Exposure      int     `json:"exposure"`
				M2MRealised   float64 `json:"m2m_realised"`
				M2MUnrealised int     `json:"m2m_unrealised"`
				OptionPremium int     `json:"option_premium"`
				Payout        int     `json:"payout"`
				Span          int     `json:"span"`
				HoldingSales  int     `json:"holding_sales"`
				Turnover      int     `json:"turnover"`
			} `json:"utilised"`
		} `json:"equity"`
		Commodity struct {
			Enabled   bool `json:"enabled"`
			Net       int  `json:"net"`
			Available struct {
				AdhocMargin   int `json:"adhoc_margin"`
				Cash          int `json:"cash"`
				Collateral    int `json:"collateral"`
				IntradayPayin int `json:"intraday_payin"`
			} `json:"available"`
			Utilised struct {
				Debits        int `json:"debits"`
				Exposure      int `json:"exposure"`
				M2MRealised   int `json:"m2m_realised"`
				M2MUnrealised int `json:"m2m_unrealised"`
				OptionPremium int `json:"option_premium"`
				Payout        int `json:"payout"`
				Span          int `json:"span"`
				HoldingSales  int `json:"holding_sales"`
				Turnover      int `json:"turnover"`
			} `json:"utilised"`
		} `json:"commodity"`
	} `json:"data"`
}

// GetMargins returns funds, cash, and margin information for the user for equity and commodity segments
func (c Client) GetMargins() (*Margins, error) {
	marginsURL := c.getRootURL() + marginsPath
	req, err := http.NewRequest("GET", marginsURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create a request for Margin")
	}
	token := c.APIKey + ":" + c.AccessToken
	req.Header.Add("Authorization", "token "+token)
	req.Header.Add("X-Kite-Version", version)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get the margins from the server")
	}
	defer resp.Body.Close()
	margins := Margins{}
	err = json.NewDecoder(resp.Body).Decode(&margins)
	if err != nil {
		return nil, errors.Wrap(err, "unable to  decode the response")
	}
	return &margins, err
}
