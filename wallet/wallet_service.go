package wallet

import (
	"strconv"

	"github.com/pocockn/crypto-compare-go/persistance"
)

// AllWallets returns a list of wallets from the database
func AllWallets() ([]Wallet, error) {
	var wallets []Wallet
	err := persistance.DB.Model(&wallets).Select()
	if err != nil {
		return nil, err
	}
	return wallets, nil
}

// GetWallet returns us a wallet from the DB from an ID
func GetWallet(id string) (*Wallet, error) {
	intID, _ := strconv.Atoi(id)
	wallet := Wallet{ID: intID}
	err := persistance.DB.Select(&wallet)
	if err != nil {
		panic(err)
	}

	return &wallet, err
}

// DeleteWallet deletes a wallet from the db
func DeleteWallet(id string) error {
	intID, _ := strconv.Atoi(id)
	wallet := Wallet{ID: intID}
	err := persistance.DB.Delete(&wallet)

	if err != nil {
		panic(err)
	}

	return nil
}
