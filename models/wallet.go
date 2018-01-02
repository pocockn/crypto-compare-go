package models

import (
	"math/rand"
	"strconv"
)

// Wallet will hold the users coins
type Wallet struct {
	ID int
	// Map of Coins the key being the coin name
	// value being the units held
	CoinsHeld map[string]int
}

// NewWallet Create a new wallet with an intial coin and unit amount
func NewWallet(initialCoinAndUnit map[string]int) *Wallet {
	return &Wallet{
		ID:        rand.Int(),
		CoinsHeld: initialCoinAndUnit,
	}
}

// SpecificBalance returns the units held for a specific coin
func (wallet *Wallet) SpecificBalance(coin string) int {
	return wallet.CoinsHeld[coin]
}

// Withdraw will take out some units from the specific coin
func (wallet *Wallet) Withdraw(coin string, amount int) {
	wallet.CoinsHeld[coin] -= amount
}

// Deposit will deposit some units to a specific coin
func (wallet *Wallet) Deposit(coin string, amount int) {
	wallet.CoinsHeld[coin] += amount
}

// AllWallets returns a list of wallets from the database
func AllWallets() ([]Wallet, error) {
	var wallets []Wallet
	err := DB.Model(&wallets).Select()
	if err != nil {
		return nil, err
	}
	return wallets, nil
}

// GetWallet returns us a wallet from the DB from an ID
func GetWallet(id string) (*Wallet, error) {
	intID, _ := strconv.Atoi(id)
	wallet := Wallet{ID: intID}
	err := DB.Select(&wallet)
	if err != nil {
		panic(err)
	}

	return &wallet, err
}
