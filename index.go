package CIPIndex_go


// Worked out definition for calculating the CIP index.


// A coin is a traded cryptocurrency.
type Coin struct {
	Value 					float64
	Ticker 					string
	TotalEffectiveSupply 	float64
	TotalSupply 			int
	Marketcap 				float64
}

// Returns the marketcap for a coin
func (c *Coin) CalculateMarketcap() float64 {
	c.Marketcap = c.TotalEffectiveSupply * c.Value
	return c.Marketcap
}

// The Index as defined by Karel L. Kubat
type CIPIndex struct {
	Coins 		[]Coin
	TotalCap 	float64
}

// Returns the value for the index
func (c *CIPIndex) Value () float64 {
	c.TotalCap = 0
	var index float64 = 0

	for _, coin := range c.Coins{
		c.TotalCap += coin.CalculateMarketcap()
	}

	for _, coin := range c.Coins{
		index += (coin.CalculateMarketcap()/c.TotalCap) * coin.Value
	}
	return index
}
