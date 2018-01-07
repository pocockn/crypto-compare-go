package wallet

import (
	"errors"
	"math/rand"
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
func (wallet *Wallet) SpecificBalance(coin string) (int, error) {
	err := wallet.isCoinInWallet(coin)
	if err != nil {
		return -1, err
	}
	return wallet.CoinsHeld[coin], nil
}

// Withdraw will take out some units from the specific coin
func (wallet *Wallet) Withdraw(coin string, amount int) error {
	if wallet.CoinsHeld[coin]-amount < 0 {
		return errors.New("Coin value cannot be below 0")
	}
	wallet.CoinsHeld[coin] -= amount

	// If balance is 0, remove the coin from the wallet
	if wallet.CoinsHeld[coin] == 0 {
		delete(wallet.CoinsHeld, coin)
	}
	return nil
}

// Deposit will deposit some units to a specific coin
func (wallet *Wallet) Deposit(coin string, amount int) error {
	err := wallet.isCoinInWallet(coin)
	if err != nil {
		return err
	}
	wallet.CoinsHeld[coin] += amount
	return nil
}

func (wallet *Wallet) isCoinInWallet(coin string) error {
	_, ok := wallet.CoinsHeld[coin]
	if !ok {
		return errors.New(coin + " not found in wallet")
	}
	return nil
}
