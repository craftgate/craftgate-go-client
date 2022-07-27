package main

import (
	"craftgate-go-client/adapter"
	"craftgate-go-client/model"
)

type Client struct {
	Installment      InstallmentApi
	PaymentReporting adapter.PaymentReportingApi
}

func CraftgateClient(apiKey, secretKey, baseURL string) *Client {
	options := model.RequestOptions{
		ApiKey:    apiKey,
		SecretKey: secretKey,
		BaseURL:   baseURL,
	}

	return &Client{
		Installment:      InstallmentApi{Opts: options},
		PaymentReporting: adapter.PaymentReportingApi{Opts: options},
	}
}
