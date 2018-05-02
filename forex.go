package CIPIndex

import (
	"github.com/mattevans/dinero"
	"time"
	"encoding/json"
)

// ForexClient is simply a wrapper around dinero.Client
type ForexClient struct {
	client *dinero.Client
}

// Init opens a connection to open exchange.
func (f *ForexClient) Init(OPEN_EXCHANGE_APP_ID string) {
	f.client = dinero.NewClient(OPEN_EXCHANGE_APP_ID)
}

// NewRateService returns a rates object which gets periodically updates and invalidates the cache
func (f *ForexClient) NewRateService(base string, refresh int) *Rates {
	f.client.Rates.SetBaseCurrency(base)
	RatesDict := make(map[string]Currency)
	rate := Rates{Rates: RatesDict, Base: base}
	go func() {
		response, err := f.client.Rates.All()
		if err != nil {
			panic(err)
		}
		rate.load(response.Rates)
		if refresh == 0 {
			return
		}
		time.Sleep(time.Duration(refresh) * time.Second)
		f.client.Cache.Expire(base)
	}()
	return &rate
}

// Currency is a currency with internal conversions to different currencies
type Currency struct {
	Ticker string
	// Value of 1 currency in Base
	Base float64
	// Rates object storing the conversions for this currency
	Rates *Rates
}

func (c *Currency) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}

func (c *Currency) Load(in []byte) error {
	if err := json.Unmarshal(in, &c); err != nil {
		return err
	}
	return nil
}

// Convert converts the currency
func (c *Currency) Convert(ticker string) Currency {
	//Requested price in base currency
	requestedBase := c.Rates.Rates[ticker].Base
	//Base currency in requested currency
	return Currency{
		Ticker: ticker,
		Base:   requestedBase,
		Rates:  c.Rates,
	}
}


// Rates is a wrapper around the open exchange api json response.
type Rates struct {
	Rates map[string]Currency
	Base  string
}

func (r *Rates) MarshalBinary() ([]byte, error) {
	return json.Marshal(r)
}

func (r *Rates) LoadfromBin(in []byte) error {
	if err := json.Unmarshal(in, &r); err != nil {
		return err
	}
	return nil
}


func (r *Rates) load(openRates map[string]float64) {
	for ticker, rate := range openRates {
		currency := Currency{
			Ticker: ticker,
			Base:   rate,
			Rates:  r,
		}
		r.Rates[ticker] = currency
	}
}
