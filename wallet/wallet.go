package wallet

import (
	"github.com/pocockn/crypto-compare-go/uuid"
)

// Wallet will hold the users coins
type Wallet struct {
	ID string
	// Map of Coins the key being the coin name
	// value being the units held
	CoinsHeld map[string]int
}

// NewWallet Create a new wallet with an intial coin and unit amount
func NewWallet(initialCoinAndUnit map[string]int) *Wallet {
	return &Wallet{
		ID:        uuid.NewUUID(),
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
