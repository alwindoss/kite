package kite_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alwindoss/kite"
)

func TestMargins(t *testing.T) {
	// echoHandler, passes back form parameter p
	echoHandler := func(w http.ResponseWriter, r *http.Request) {
		margins := &kite.Margins{}
		margins.Status = "Success"
		json.NewEncoder(w).Encode(margins)
	}

	// create test server with handler
	ts := httptest.NewServer(http.HandlerFunc(echoHandler))
	defer ts.Close()

	c := kite.Client{
		HTTPClient: &http.Client{},
		RootURL:    ts.URL,
	}
	m, err := c.GetMargins()
	if m.Status != "Success" {
		t.Errorf("expected status %s but found %s", "Succcess", m.Status)
		t.Fail()
	}
	if err != nil {
		t.Errorf("retuned error: %v\n", err)
		t.Fail()
	}
}
