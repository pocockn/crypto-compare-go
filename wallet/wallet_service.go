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
		return nil, err
	}

	return &wallet, err
}

// DeleteWallet deletes a wallet from the db
func DeleteWallet(id string) error {
	intID, _ := strconv.Atoi(id)
	wallet := Wallet{ID: intID}
	err := persistance.DB.Delete(&wallet)
	if err != nil {
		return err
	}

	return nil
}

// DepositUnits deposits units to a wallet based on it's ID
func DepositUnits(id string, coin string, units int) error {
	wallet, err := GetWallet(id)
	err = wallet.Deposit(coin, units)
	if err != nil {
		return err
	}
	err = persistance.DB.Update(wallet)
	if err != nil {
		return err
	}
	return nil
}

// WithdrawUnits withdraws units from a wallet based on it's ID
func WithdrawUnits(id string, coin string, units int) error {
	wallet, err := GetWallet(id)
	err = wallet.Withdraw(coin, units)
	if err != nil {
		return err
	}
	err = persistance.DB.Update(wallet)
	if err != nil {
		return err
	}
	return nil
}
