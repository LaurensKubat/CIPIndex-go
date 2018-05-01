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
	rates := client.WatchRates("USD", 15)
	for {
		fmt.Println(rates)
		time.Sleep(15*time.Second)
	}
}