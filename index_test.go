package CIPIndex_go

import (
"testing"
)

func TestCoin_CalculateMarketcap(t *testing.T) {
	BTC := Coin {
		Ticker:					"BTC",
		Value:					8000,
		TotalEffectiveSupply:	17000000,
	}

	if BTC.CalculateMarketcap() != 136000000000 {
		t.Errorf("Marketcap not properly calculated")
	}
}

func TestCIPIndex_Value(t *testing.T) {
	BTC := Coin {
		Ticker:					"BTC",
		Value:					8000,
		TotalEffectiveSupply:	17000000,
	}

	ETH := Coin {
		Ticker:					"ETH",
		Value:					500,
		TotalEffectiveSupply:	100000000,
	}

	RIP := Coin {
		Ticker:					"RIP",
		Value:					2,
		TotalEffectiveSupply:	10000000000,
	}

	index := CIPIndex{
		Coins: []Coin{BTC, ETH, RIP},
	}

	value := index.Value()

	if int(value) != 5403 {
		t.Errorf("Index not properly calculated!")
	}
}
