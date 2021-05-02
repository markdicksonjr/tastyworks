package tastyworks

import (
	"os"
	"testing"
)

func TestGetOptionChainForTicker(t *testing.T) {
	v, err := Authorize(os.Getenv("TT_USER"), os.Getenv("TT_PASSWORD"))
	if err != nil {
		t.Fatal(err)
	}
	chain, err := GetOptionChainForTicker(v.Data.SessionToken, "SPY")
	if err != nil {
		t.Fatal(err)
	}

	_ = chain
	// TODO: check response
}

func TestGetOptionChainForTicker_Error(t *testing.T) {
	_, err := GetOptionChainForTicker("", "SPY")
	if err == nil {
		t.Fatal("expected error for missing session token but did not get one")
	}
}

func TestGetNestedOptionChainForTicker(t *testing.T) {
	v, err := Authorize(os.Getenv("TT_USER"), os.Getenv("TT_PASSWORD"))
	if err != nil {
		t.Fatal(err)
	}
	chain, err := GetNestedOptionChainForTicker(v.Data.SessionToken, "SPY")
	if err != nil {
		t.Fatal(err)
	}

	_ = chain
	// TODO: check response
}

func TestGetNestedOptionChainForTicker_Error(t *testing.T) {
	_, err := GetNestedOptionChainForTicker("", "SPY")
	if err == nil {
		t.Fatal("expected error for missing session token but did not get one")
	}
}
