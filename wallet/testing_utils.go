package wallet

import "github.com/pocockn/crypto-compare-go/persistance"

// BootstrapWallet adds a wallet to the database for testing purposes
func BootstrapWallet() *Wallet {
	coinMap := make(map[string]int)
	coinMap["BTC"] = 100
	wallet := &Wallet{
		ID:        1234,
		CoinsHeld: coinMap,
	}
	persistance.DB.Insert(wallet)

	return wallet
}
