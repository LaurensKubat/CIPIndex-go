package CIPIndex_go


// Worked out definition for calculating the CIP index.
type Value struct {
	Base 		float64
	Currency 	Currency
}

func (v *Value) toCurrency (ticker string) float64{
	return v.Base * v.Currency.Value(ticker)
}

// A coin is a traded cryptocurrency.
type Coin struct {
	// Value in USD, EURO or BTC depending on what you want
	Value 					Value

	// Ticker, not required
	Ticker 					string

	// Value used in calculation
	TES 	float64

	// Theoretical total supply, not required.
	TotalSupply 			int
	Marketcap 				float64
}

// Returns the marketcap for a coin
func (c *Coin) CalculateMarketcap() Value {
	c.Marketcap = c.TES * c.Value.Base
	return Value{c.Marketcap, c.Value.Currency}
}

// The Index as defined by Karel L. Kubat
type CIPIndex struct {
	Coins 		[]Coin
	TotalCap 	float64
	// Main currency you use for the index
	Currency 	Currency
}

// Returns the value for the index
func (c *CIPIndex) Value () float64 {
	c.TotalCap = 0
	var index float64 = 0

	for _, coin := range c.Coins{
		marketcap := coin.CalculateMarketcap()
		c.TotalCap += marketcap.toCurrency(c.Currency.Ticker)
	}

	for _, coin := range c.Coins{
		marketcap := coin.CalculateMarketcap()
		index += (marketcap.toCurrency(c.Currency.Ticker)/c.TotalCap) *
			coin.Value.toCurrency(c.Currency.Ticker)
	}
	return index
}
