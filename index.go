// Package CIPIndex implements easy calculation of the CIP index in combination
// with currency conversions using the Open Exchange Rates API

package CIPIndex

// Value is the combination of a currency and the amount of that
// Currency
type Value struct {
	Price    float64
	Currency Currency
}

// toCurrency converts the Value to a specific currency.
func (v *Value) toCurrency(ticker string) float64 {
	return v.Price * v.Currency.Value(ticker)
}

// ExchangeCoin is the information about a cryptocurrency derived from
// an exchange.
type ExchangeCoin struct {
	Value  Value
	Ticker string
	Volume float64
}

// Coin is the summarized value of a cryptocurrency.
type Coin struct {
	// Value in USD, EURO or BTC depending on what you want
	Value Value

	// Ticker, not required
	Ticker string

	// Total Effective supply: what is not locked.
	TES float64

	// Theoretical total supply, not required.
	TotalSupply int
	Marketcap   float64
}

// Init readies an nil Coin for loading.
func (c *Coin) Init(currency Currency, ticker string) {
	c.Value.Currency = currency
	c.Ticker = ticker
}

//Load generates a complete Coin struct from exchange coins and the coins supply
func (c *Coin) Load(coins []ExchangeCoin, TES float64) {
	var totalVolume float64
	for _, coin := range coins {
		totalVolume += coin.Volume
	}

	for _, coin := range coins {
		weight := coin.Volume / totalVolume
		c.Value.Price += coin.Value.toCurrency(c.Value.Currency.Ticker) * weight
	}
	c.TES = TES
}

// CalculateMarketcap calculates the marketcap of a coin and returns the Value
func (c *Coin) CalculateMarketcap() Value {
	c.Marketcap = c.TES * c.Value.Price
	return Value{c.Marketcap, c.Value.Currency}
}

// CIPIndex is a representation of the index.
type CIPIndex struct {
	Coins    []Coin
	TotalCap float64
	// Main currency you use for the index
	Currency Currency
}

// Value returns the current point value of the CIPIndex.
func (c *CIPIndex) Value() float64 {
	c.TotalCap = 0
	var index float64

	for _, coin := range c.Coins {
		marketcap := coin.CalculateMarketcap()
		c.TotalCap += marketcap.toCurrency(c.Currency.Ticker)
	}

	for _, coin := range c.Coins {
		marketcap := coin.CalculateMarketcap()
		index += (marketcap.toCurrency(c.Currency.Ticker) / c.TotalCap) *
			coin.Value.toCurrency(c.Currency.Ticker)
	}
	return index
}
