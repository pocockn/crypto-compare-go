package persistance

import (
	"crypto-compare-go/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	InitDB("crypto_compare_test")
	DB.Exec("TRUNCATE TABLE wallets;")
}

func TestGetWallets(t *testing.T) {
	wallet := BootstrapWallet()
	allWallets, err := AllWallets()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, &allWallets[0], wallet)
}

func TestGetWallet(t *testing.T) {
	wallet := BootstrapWallet()
	walletFromDb, error := GetWallet("1234")
	if error != nil {
		t.Error(error)
	}
	assert.Equal(t, walletFromDb, wallet)
}

func TestAddSecondCoinToWallet(t *testing.T) {
	_ = BootstrapWallet()
	walletToSearch := &models.Wallet{ID: 1234}
	err := DB.Select(walletToSearch)
	walletToSearch.CoinsHeld["ETH"] = 400
	err = DB.Update(walletToSearch)
	if err != nil {
		t.Error(err)
	}
	ethUnits := walletToSearch.CoinsHeld["ETH"]
	assert.Equal(t, ethUnits, 400)
}

func TestDeleteWallet(t *testing.T) {
	wallet := BootstrapWallet()
	err := DeleteWallet("1234")
	if err != nil {
		t.Error(err)
	}
	err = DB.Select(wallet)
	assert.Equal(t, "pg: no rows in result set", err.Error())
}
