package go_bithumb

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func init() {
	_ = godotenv.Load()
}

func TestPublicConfig(t *testing.T)  {
	bithumb := NewBithhumb()
	configRes, err := bithumb.GetConfig()
	assert.Nil(t, err)

	for _, coin := range configRes.Data.CoinConfig {
		if coin.Name == "GTH" || coin.Name == "LTC"{
			fmt.Printf("%+v\n", coin)
		}
	}
	for _, market := range configRes.Data.SpotConfig {
		if strings.Contains(market.Symbol, "GTH") {
			fmt.Printf("%+v\n", market)
		}
	}

}

/* func TestPublicOrderBook(t *testing.T)  {
	bithumb := NewBithhumb()
	orders, err := bithumb.GetOrderBook("GTH-USDT")
	assert.Nil(t, err)
	fmt.Println(orders)
}

func TestPublicTrades(t *testing.T)  {
	bithumb := NewBithhumb()
	orders, err := bithumb.GetTrades("GTH-USDT")
	assert.Nil(t, err)
	fmt.Println(orders)
}*/

func TestPrivateBalance(t *testing.T)  {
	bithumb := NewBithhumbAuth(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
	balance, err := bithumb.Assets("USDT")
	assert.Nil(t, err)
	fmt.Println(balance)
}

func TestPrivateDeposits(t *testing.T)  {
	bithumb := NewBithhumbAuth(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
	deposits, err := bithumb.DepositHistory("USDT")
	assert.Nil(t, err)
	fmt.Println(deposits)
}