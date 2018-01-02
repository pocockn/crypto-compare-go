package models

import (
	"testing"
)

func TestWalletCreation(t *testing.T) {
	coinMap := make(map[string]int)
	expectedUnits := 100
	coinMap["BTC"] = 100
	btcWallet := NewWallet(coinMap)
	if expectedUnits != btcWallet.CoinsHeld["BTC"] {
		t.Error("Unexpected value, should be ", expectedUnits)
	}
}

func TestWalletHasCorrectBalance(t *testing.T) {
	coinMap := make(map[string]int)
	expectedUnits := 100
	coinMap["BTC"] = 100
	btcWallet := NewWallet(coinMap)
	actualUnits := btcWallet.SpecificBalance("BTC")
	if expectedUnits != actualUnits {
		t.Error("Unexpected value, should be ", expectedUnits)
	}
}

func TestUnitsCanBeWithDrawn(t *testing.T) {
	coinMap := make(map[string]int)
	expectedUnits := 50
	coinMap["BTC"] = 100
	btcWallet := NewWallet(coinMap)
	btcWallet.Withdraw("BTC", 50)
	actualUnits := btcWallet.SpecificBalance("BTC")
	if expectedUnits != actualUnits {
		t.Error("Unexpected value, should be ", expectedUnits)
	}
}

func TestUnitsCanBeDeposited(t *testing.T) {
	coinMap := make(map[string]int)
	expectedUnits := 150
	coinMap["BTC"] = 100
	btcWallet := NewWallet(coinMap)
	btcWallet.Deposit("BTC", 50)
	actualUnits := btcWallet.SpecificBalance("BTC")
	if expectedUnits != actualUnits {
		t.Error("Unexpected value, should be ", expectedUnits)
	}
}
