package kite_test

import (
	"testing"

	"github.com/alwindoss/kite"
)

func TestLogin(t *testing.T) {
	client := kite.NewClient("APIKey")
	accessToken := client.Login()
	if accessToken != "ACCESS TOKEN" {
		t.Errorf("Unpected Result")
	}
}
