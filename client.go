package main

import (
	"craftgate-go-client/adapter"
)

type Client struct {
	Installment      adapter.Installment
	PaymentReporting adapter.PaymentReporting
	Wallet           adapter.Wallet
}

func CraftgateClient(apiKey, secretKey, baseURL string) *Client {
	options := adapter.RequestOptions{
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
