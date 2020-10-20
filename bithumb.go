package go_bithumb

import (
	"fmt"
	"github.com/PrettyBoyHelios/go-bithumb/models"
	"github.com/shopspring/decimal"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Bithumb struct {
	client Client
}

func NewBithhumb() *Bithumb {
	b := new(Bithumb)
	b.client = Client{
		apiKey:    os.Getenv("ApiKey"),
		secretKey: os.Getenv("SecretKey"),
		client:    http.Client{},
	}
	return b
}

func NewBithhumbAuth(apiKey string, secretKey string) *Bithumb {
	b := new(Bithumb)
	b.client = Client{
		apiKey:    apiKey,
		secretKey: secretKey,
		client:    http.Client{},
		url: "https://global-openapi.bithumb.pro/openapi/v1",
	}
	return b
}


func (b *Bithumb) Assets(coinType string) (*models.AssetsResp, error) {
	var r models.AssetsResp
	p := struct {
		CoinType  string `json:"coinType"`
		AssetType string `json:"assetType"`
	}{
		coinType, "spot",
	}
	err := b.client.post(b.client.url + "/spot/assetList", p, &r)
	if err != nil {
		return &r, err
	}
	return &r, nil
}

func (b *Bithumb) Withdraw(asset string, address string, quantity decimal.Decimal, mark string) (bool, error) {
	var r interface{}
	p := struct {
		CoinType string `json:"coinType"`
		Address  string `json:"address"`
		Quantity string `json:"quantity"`
		Mark     string `json:"mark"`
	}{
		asset, address, quantity.String(), mark,
	}
	err := b.client.post(b.client.url + "/spot/assetList", p, &r)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (b *Bithumb) CreateOrder(symbol string, side string, quantity decimal.Decimal, price decimal.Decimal, orderType string) (*models.CreateOrderResp, error) {
	var c models.CreateOrderResp
	var pr = price
	if orderType == "market" {
		pr = decimal.NewFromFloat(-1)
	}
	p := struct {
		Symbol    string `json:"symbol"`
		Type      string `json:"type"`
		Side      string `json:"side"`
		Price     string `json:"price"`
		Quantity  string `json:"quantity"`
		Timestamp string `json:"timestamp"`
	}{
		symbol, orderType, side, pr.String(), quantity.String(), strconv.FormatInt(time.Now().UTC().UnixNano()/1e6, 10),
		// symbol, orderType, side, pr.String(), quantity.String(),
	}
	err := b.client.post(b.client.url + "/spot/placeOrder", p, &c)
	if err != nil {
		return &c, err
	}
	fmt.Println("message: ", c.Msg)
	return &c, nil
}

func (b *Bithumb) OrderDetail(symbol string, orderId string) (*models.OrderDetailResp, error) {
	var c models.OrderDetailResp
	p := struct {
		Symbol  string `json:"symbol"`
		OrderId string `json:"orderId"`
	}{
		symbol, orderId,
	}
	err := b.client.post(b.client.url + "/spot/singleOrder", p, &c)
	if err != nil {
		return &c, err
	}
	fmt.Println("message: ", &c.Msg)
	return &c, nil
}

func (b *Bithumb) GetConfig() (*models.ConfigResp, error) {
	var c models.ConfigResp
	err := b.client.get(b.client.url + "/spot/config", &c)
	if err != nil {
		return &c, err
	}
	fmt.Println("message: ", &c.Data)
	return &c, nil
}

func (b *Bithumb) DepositHistory(asset string) (*models.DepositHistory, error) {
	var c models.DepositHistory
	fromDate := int64(1000 * 60 * 60 * 24 * 90) // 90 days in milliseconds
	p := struct {
		Symbol  string `json:"coin"`
		Start string `json:"start"`
	}{
		asset, strconv.FormatInt(time.Now().UnixNano()/1e6 - fromDate, 10),
	}
	fmt.Println(time.Now().UnixNano()/1e6)
	err := b.client.post(b.client.url + "/wallet/depositHistory", p, &c)
	if err != nil {
		return &c, err
	}
	fmt.Println("message: ", &c.Msg)
	return &c, nil
}