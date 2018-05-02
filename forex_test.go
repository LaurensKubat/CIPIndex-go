package CIPIndex

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
	"time"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func TestForexClient_Init(t *testing.T) {
	client := ForexClient{}
	client.Init(os.Getenv("OPEN_EXCHANGE_APP_ID"))
	client.client.Rates.SetBaseCurrency("USD")
	_, err := client.client.Rates.All()
	if err != nil {
		fmt.Println(err)
		t.Errorf("Failed to get a response from OpenX")
	}
}

func TestForexClient_WatchRates(t *testing.T) {
	client := ForexClient{}
	client.Init(os.Getenv("OPEN_EXCHANGE_APP_ID"))
	rates := client.NewRateService("USD", 15)
	time.Sleep(3 * time.Second)
	if rates.Rates["EUR"] == 0 {
		t.Errorf("RateService failed")
	}
}

func TestCurrency_Value(t *testing.T) {
	rates, _ := GetRateservice()
	currency := Currency{
		"USD",
		1,
		rates,
	}
	converted := currency.Convert("EUR").Base
	if currency.Base < converted {
		t.Errorf("Are dollars worth more than EUROS!?")
	}
}

func TestCurrency_MarshalBinary(t *testing.T) {
	rates, _ := GetRateservice()
	currency := Currency{
		"USD",
		1,
		rates,
	}
	bin, err := currency.MarshalBinary()
	if err != nil {
		t.Error("Error in marshalling currency", err)
	}
	currency = Currency{}
	currency.Load(bin)
	if currency.Ticker != "USD" {
		t.Error("Error in unmarshalling currency Ticker")
	}
	if currency.Base != 1 {
		t.Error("Error in unmarshalling currency Base")
	}

	if currency.Rates.Rates["EUR"] != rates.Rates["EUR"] {
		t.Error("Error in unmarshalling currency Rates")
	}

	if currency.Rates.Rates["EUR"] == 0 {
		t.Error("Error in unmarshalling currency Rates")
	}
}
