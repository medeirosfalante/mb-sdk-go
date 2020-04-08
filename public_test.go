package mercadobitcoin_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	mercadobitcoin "github.com/rafaeltokyo/mb-sdk-go"
)

func TestGetTicker(t *testing.T) {
	godotenv.Load()
	client := mercadobitcoin.New("", os.Getenv("ENV"))
	response, errAPI, err := client.Public().Ticker("BTC")

	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if response == nil {
		t.Error("response is null")
		return
	}

}

func TestGetOrders(t *testing.T) {
	godotenv.Load()
	client := mercadobitcoin.New("", os.Getenv("ENV"))
	response, errAPI, err := client.Public().OrderBook("BTC")
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if response == nil {
		t.Error("response is null")
		return
	}
}

func TestGetTrades(t *testing.T) {
	godotenv.Load()
	client := mercadobitcoin.New("", os.Getenv("ENV"))
	response, errAPI, err := client.Public().Trades("BTC")
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if response == nil {
		t.Error("response is null")
		return
	}
}
