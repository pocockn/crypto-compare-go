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
