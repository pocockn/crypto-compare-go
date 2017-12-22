package wallet

import "testing"

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
