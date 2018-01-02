package persistance

import (
	"crypto-compare-go/models"
	"strconv"
)

// AllWallets returns a list of wallets from the database
func AllWallets() ([]models.Wallet, error) {
	var wallets []models.Wallet
	err := DB.Model(&wallets).Select()
	if err != nil {
		return nil, err
	}
	return wallets, nil
}

// GetWallet returns us a wallet from the DB from an ID
func GetWallet(id string) (*models.Wallet, error) {
	intID, _ := strconv.Atoi(id)
	wallet := models.Wallet{ID: intID}
	err := DB.Select(&wallet)
	if err != nil {
		panic(err)
	}

	return &wallet, err
}
