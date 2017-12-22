package wallet

// Wallet will hold the users coins
type Wallet struct {
	// Map of Coins the key being the coin name
	// value being the units held
	CoinsHeld map[string]int
}

// NewWallet Create a new wallet with an intial coin and unit amount
func NewWallet(initialCoinAndUnit map[string]int) *Wallet {
	return &Wallet{
		CoinsHeld: initialCoinAndUnit,
	}
}

// SpecificBalance returns the units held for a specific coin
func (wallet *Wallet) SpecificBalance(coin string) int {
	return wallet.CoinsHeld[coin]
}
