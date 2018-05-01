package CIPIndex_go

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
	"testing"
)

func init(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func GetRateservice () (*Rates, *ForexClient) {
	client := ForexClient{}
	client.Init(os.Getenv("OPEN_EXCHANGE_APP_ID"))
	rates := client.NewRateService("USD", 0)
	time.Sleep(1 * time.Second)
	return rates, &client
}



func TestCoin_CalculateMarketcap(t *testing.T) {
	rates, _ := GetRateservice()

	value := Value{
		Base:   		8000,
		Currency: 		Currency{
			"USD",
			1,
			rates,
		},
	}

	BTC := Coin {
		Ticker:		"BTC",
		Value:		value,
		TES:		17000000,
	}

	if BTC.CalculateMarketcap().Base != 136000000000 {
		t.Errorf("Marketcap not properly calculated")
	}
}

func TestCIPIndex_Value(t *testing.T) {
	rates, _ := GetRateservice()
	BTC := Coin {
		Ticker:			"BTC",
		Value:			Value{
							Base:   	8000,
							Currency: 	Currency{
									"USD",
									1,
									rates,
								},
		},
		TES:	17000000,
	}

	ETH := Coin {
		Ticker:					"ETH",
		Value:					Value{
									Base:   		500,
									Currency: 		Currency{
										"USD",
										1,
										rates,
									},
		},
		TES:	100000000,
	}

	RIP := Coin {
		Ticker:					"RIP",
		Value:					Value{
							Base:   		2,
							Currency: 		Currency{
								"USD",
								1,
								rates,
							},
		},
		TES:	10000000000,
	}

	index := CIPIndex{
		Coins: []Coin{BTC, ETH, RIP},
		Currency: 	Currency{
			Ticker: 	"USD",
			Base: 		1,
			Rates: 		rates,
		},
	}

	value := index.Value()
	if int(value) != 5403 {
		t.Errorf("Index not properly calculated!")
	}
}
