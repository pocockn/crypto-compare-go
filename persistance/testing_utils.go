package persistance

import "crypto-compare-go/models"

// BootstrapWallet adds a wallet to the database for testing purposes
func BootstrapWallet() *models.Wallet {
	coinMap := make(map[string]int)
	coinMap["BTC"] = 100
	wallet := &models.Wallet{
		ID:        1234,
		CoinsHeld: coinMap,
	}
	DB.Insert(wallet)

	return wallet
}
