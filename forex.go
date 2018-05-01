package CIPIndex_go

import (
	"github.com/mattevans/dinero"
	"time"
)

// A forex client
type ForexClient struct {
	client  *dinero.Client
}

func (f *ForexClient) Init(OPEN_EXCHANGE_APP_ID string)	{
	f.client = dinero.NewClient(OPEN_EXCHANGE_APP_ID)
}

func (f *ForexClient) WatchRates(base string, refresh int) *Rates {
	f.client.Rates.SetBaseCurrency(base)
	RatesDict := make(map[string]Currency)
	rate
	go func() {
		response, err := f.client.Rates.All()
		if err != nil {
			panic(err)
		}
		rates.load(response.Rates)
		time.Sleep(time.Duration(refresh) * time.Second)
	}()
	return &rates
}

// A forex traded currency compared to a base currency
type Currency struct {
	Ticker 		string
	// Value of 1 currency in Base
	Base	float64
}

// Used to convert between fiat currencies using USD as a base currency
type Rates struct {
	rates 	map[string]Currency
}

func (r *Rates) load (openRates map[string]float64){
	for ticker, rate := range openRates{
		currency := Currency{
			Ticker:	ticker,
			Base:	rate,
		}
		r.rates[ticker] = currency
	}
}


// Used to easily convert a coins value.
type Value struct {

}