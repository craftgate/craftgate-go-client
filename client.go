package main

import (
	"craftgate-go-client/adapter"
	"craftgate-go-client/model"
)

type Client struct {
	Installment      adapter.Installment
	PaymentReporting adapter.PaymentReporting
	Wallet           adapter.Wallet
}

func CraftgateClient(apiKey, secretKey, baseURL string) *Client {
	options := model.RequestOptions{
		ApiKey:    apiKey,
		SecretKey: secretKey,
		BaseURL:   baseURL,
	}

	return &Client{
		Installment:      adapter.Installment{Opts: options},
		PaymentReporting: adapter.PaymentReporting{Opts: options},
		Wallet:           adapter.Wallet{Opts: options},
	}
}
