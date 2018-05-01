package CIPIndex

// Worked out definition for calculating the CIP index.
type Value struct {
	Price    float64
	Currency Currency
}

func (v *Value) toCurrency(ticker string) float64 {
	return v.Price * v.Currency.Value(ticker)
}

// Values derived from a single exchange
type ExchangeCoin struct {
	Value  Value
	Ticker string
	Volume float64
}

// A coin is a traded cryptocurrency.
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

//Initialization ready for loading
func (c *Coin) Init(currency Currency, ticker string) {
	c.Value.Currency = currency
	c.Ticker = ticker
}

//Load a Coin from exchange + supply data
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

// Returns the marketcap for a coin
func (c *Coin) CalculateMarketcap() Value {
	c.Marketcap = c.TES * c.Value.Price
	return Value{c.Marketcap, c.Value.Currency}
}

// The Index as defined by Karel L. Kubat
type CIPIndex struct {
	Coins    []Coin
	TotalCap float64
	// Main currency you use for the index
	Currency Currency
}

// Returns the value for the index
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
