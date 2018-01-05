package wallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWalletCreation(t *testing.T) {
	expectedUnits := 100
	btcWallet := createWallet()
	if expectedUnits != btcWallet.CoinsHeld["BTC"] {
		t.Error("Unexpected value, should be ", expectedUnits)
	}
}

func TestWalletHasCorrectBalance(t *testing.T) {
	expectedUnits := 100
	btcWallet := createWallet()
	actualUnits := btcWallet.SpecificBalance("BTC")
	if expectedUnits != actualUnits {
		t.Error("Unexpected value, should be ", expectedUnits)
	}
}

func TestUnitsCanBeWithDrawn(t *testing.T) {
	expectedUnits := 50
	btcWallet := createWallet()
	btcWallet.Withdraw("BTC", 50)
	actualUnits := btcWallet.SpecificBalance("BTC")
	if expectedUnits != actualUnits {
		t.Error("Unexpected value, should be ", expectedUnits)
	}
}

func TestBalanceCannotBecomeNegative(t *testing.T) {
	btcWallet := createWallet()
	err := btcWallet.Withdraw("BTC", 200)
	if err == nil {
		t.Error("Balance shouldn't be able to become negative")
	}
}

func TestCoinIsDeletedOnceBalanceReachesZero(t *testing.T) {
	btcWallet := createWallet()
	btcWallet.Withdraw("BTC", 100)
	emptyMap := make(map[string]int)
	assert.Equal(t, btcWallet.CoinsHeld, emptyMap)
}

func TestUnitsCanBeDeposited(t *testing.T) {
	btcWallet := createWallet()
	expectedUnits := 150
	btcWallet.Deposit("BTC", 50)
	actualUnits := btcWallet.SpecificBalance("BTC")
	if expectedUnits != actualUnits {
		t.Error("Unexpected value, should be ", expectedUnits)
	}
}

func createWallet() *Wallet {
	coinMap := make(map[string]int)
	coinMap["BTC"] = 100
	return NewWallet(coinMap)
}
