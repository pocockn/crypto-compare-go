package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	InitDB("crypto_compare_test")
	DB.Exec("TRUNCATE TABLE wallets;")
}

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

func TestGetWallets(t *testing.T) {
	wallet := bootstrapWallet()
	var wallets []Wallet
	err := DB.Model(&wallets).Select()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, &wallets[0], wallet)
}

func TestGetWallet(t *testing.T) {
	wallet := bootstrapWallet()
	walletToSearch := Wallet{ID: 1234}
	err := DB.Select(&walletToSearch)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, &walletToSearch, wallet)
}

func TestAddSecondCoinToWallet(t *testing.T) {
	_ = bootstrapWallet()
	walletToSearch := &Wallet{ID: 1234}
	err := DB.Select(walletToSearch)
	walletToSearch.CoinsHeld["ETH"] = 400
	err = DB.Update(walletToSearch)
	if err != nil {
		t.Error(err)
	}
	ethUnits := walletToSearch.CoinsHeld["ETH"]
	assert.Equal(t, ethUnits, 400)
}

func bootstrapWallet() *Wallet {
	coinMap := make(map[string]int)
	coinMap["BTC"] = 100
	wallet := &Wallet{
		ID:        1234,
		CoinsHeld: coinMap,
	}
	DB.Insert(wallet)

	return wallet
}
