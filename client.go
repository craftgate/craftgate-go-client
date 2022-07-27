package main

import (
	"craftgate-go-client/adapter"
	"craftgate-go-client/model"
)

type Client struct {
	Installment      adapter.InstallmentApi
	PaymentReporting adapter.PaymentReportingApi
}

func CraftgateClient(apiKey, secretKey, baseURL string) *Client {
	options := model.RequestOptions{
		ApiKey:    apiKey,
		SecretKey: secretKey,
		BaseURL:   baseURL,
	}

	return &Client{
		Installment:      adapter.InstallmentApi{Opts: options},
		PaymentReporting: adapter.PaymentReportingApi{Opts: options},
	}
}
