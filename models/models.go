package models

import "github.com/shopspring/decimal"

type ConfigResp struct {
	baseResp
	Data struct {
		CoinConfig []struct {
			MakerFeeRate   decimal.Decimal `json:"makerFeeRate"`
			MinWithdraw    decimal.Decimal `json:"minWithdraw"`
			WithdrawFee    decimal.Decimal `json:"withdrawFee"`
			Name           string          `json:"name"`
			DepositStatus  string          `json:"depositStatus"`
			FullName       string          `json:"fullName"`
			TakerFeeRate   decimal.Decimal `json:"takerFeeRate"`
			WithdrawStatus decimal.Decimal `json:"withdrawStatus"`
			MinTxAmount    decimal.Decimal `json:"minTxAmt"`
		} `json:"coinConfig"`
		ContractConfig []struct {
			Symbol       string          `json:"symbol"`
			MakerFeeRate decimal.Decimal `json:"makerFeeRate"`
			TakerFeeRate decimal.Decimal `json:"takerFeeRate"`
		} `json:"contractConfig"`
		SpotConfig []struct {
			Symbol       string   `json:"symbol"`
			Accuracy     []string `json:"accuracy"`
			PercentPrice struct {
				MultiplierDown decimal.Decimal `json:"multiplierDown"`
				MultiplierUp   decimal.Decimal `json:"multiplierUp"`
			} `json:"percentPrice"`
		} `json:"spotConfig"`
	} `json:"data"`
}

type CreateOrderResp struct {
	baseResp
	Data struct {
		OrderId string
		Symbol  string
	} `json:"data"`
}

type DepositHistory struct {
	baseResp
	Data []struct {
		CoinType   string          `json:"coinType"`
		Address    string          `json:"address"`
		Quantity   decimal.Decimal `json:"quantity"`
		CreateTime int64           `json:"createTime"`
		Txid       string          `json:"txid"`
		AcountName string          `json:"acountName"`
		ID         string          `json:"id"`
		Status     string          `json:"status"`
	} `json:"data"`
}

type WithdrawHistory struct {
	baseResp
	Data []struct {
		CoinType     string          `json:"coinType"`
		Address      string          `json:"address"`
		Quantity     decimal.Decimal `json:"quantity"`
		CreateTime   int64           `json:"createTime"`
		Fee          decimal.Decimal `json:"fee"`
		WithdrawType string          `json:"withdrawType"`
		Memo         string          `json:"memo"`
		ID           string          `json:"id"`
		Status       string          `json:"status"`
		Txid         string          `json:"txid"`
	} `json:"data"`
}

type OrderDetailResp struct {
	baseResp
	Data struct {
		OrderID    string          `json:"orderId"`
		Symbol     string          `json:"symbol"`
		Price      decimal.Decimal `json:"price"`
		TradedNum  decimal.Decimal `json:"tradedNum"`
		Quantity   decimal.Decimal `json:"quantity"`
		AvgPrice   decimal.Decimal `json:"avgPrice"`
		Status     string          `json:"status"`
		Type       string          `json:"type"`
		Side       string          `json:"side"`
		CreateTime string          `json:"createTime"`
		TradeTotal decimal.Decimal `json:"tradeTotal"`
	} `json:"data"`
}

type AssetsResp struct {
	baseResp
	Data []struct {
		CoinType    string
		Count       decimal.Decimal
		Frozen      decimal.Decimal
		Type        string
		BtcQuantity decimal.Decimal
	}
}

type baseResp struct {
	Code      string
	Msg       string
	Timestamp int64
	Data      interface{}
}
