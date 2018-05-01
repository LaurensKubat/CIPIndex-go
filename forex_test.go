package CIPIndex_go

import (
	"testing"
	"os"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"time"
)

func init(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func TestForexClient_Init(t *testing.T){
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
	if rates.rates["EUR"].Ticker != "EUR"{
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
	converted := currency.Value("EUR")
	if currency.Base < converted {
		t.Errorf("Are dollars worth more than EUROS!?")
	}
}