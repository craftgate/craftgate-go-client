package main

import (
	"craftgate-go-client/adapter"
	"craftgate-go-client/model"
)

type Client struct {
	Installment         adapter.Installment
	Payment             adapter.Payment
	PaymentReporting    adapter.PaymentReporting
	Wallet              adapter.Wallet
	Onboarding          adapter.Onboarding
	PayByLink           adapter.PayByLink
	Settlement          adapter.Settlement
	SettlementReporting adapter.SettlementReporting
}

func CraftgateClient(apiKey, secretKey, baseURL string) *Client {
	options := model.RequestOptions{
		ApiKey:    apiKey,
		SecretKey: secretKey,
		BaseURL:   baseURL,
	}

	return &Client{
		Installment:         adapter.Installment{Opts: options},
		Payment:             adapter.Payment{Opts: options},
		PaymentReporting:    adapter.PaymentReporting{Opts: options},
		Wallet:              adapter.Wallet{Opts: options},
		Onboarding:          adapter.Onboarding{Opts: options},
		PayByLink:           adapter.PayByLink{Opts: options},
		Settlement:          adapter.Settlement{Opts: options},
		SettlementReporting: adapter.SettlementReporting{Opts: options},
	}
}
