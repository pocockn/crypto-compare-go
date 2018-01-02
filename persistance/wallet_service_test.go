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
	var wallets []models.Wallet
	err := DB.Model(&wallets).Select()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, &wallets[0], wallet)
}

func TestGetWallet(t *testing.T) {
	wallet := BootstrapWallet()
	walletToSearch := models.Wallet{ID: 1234}
	err := DB.Select(&walletToSearch)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, &walletToSearch, wallet)
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
