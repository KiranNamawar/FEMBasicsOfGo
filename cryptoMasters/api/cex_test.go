package api_test

import (
	"cryptomaster/api"
	"testing"
)

func TestAPICall(t *testing.T) {
	_, err := api.GetRate("")
	if err == nil {
		t.Error("Error not found")
	}
}
