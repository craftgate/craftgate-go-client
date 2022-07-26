package main

import (
	"fmt"
	"net/http"
)

type InstallmentParams struct {
	binNumber                               string
	price                                   int
	currency                                string
	distinctCardBrandsWithLowestCommissions bool
}

type InstallmentPrice struct {
	InstallmentPrice  float64 `json:"installmentPrice"`
	TotalPrice        float64 `json:"totalPrice"`
	InstallmentNumber int     `json:"installmentNumber"`
	InstallmentLabel  string  `json:"installmentLabel"`
}

type Installment struct {
	BinNumber         string             `json:"binNumber"`
	Price             float64            `json:"price"`
	CardType          string             `json:"cardType"`
	CardAssociation   string             `json:"cardAssociation"`
	CardBrand         string             `json:"cardBrand"`
	BankName          string             `json:"bankName"`
	BankCode          int                `json:"bankCode"`
	Force3Ds          bool               `json:"force3ds"`
	CvcRequired       bool               `json:"cvcRequired"`
	Commercial        bool               `json:"commercial"`
	PosAlias          string             `json:"posAlias"`
	InstallmentPrices []InstallmentPrice `json:"installmentPrices"`
}

func (c *Client) Search(installmentParams InstallmentParams) (*Response[Installment], error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/installment/v1/installments", c.BaseURL), nil)
	if err != nil {
		panic(err)
		return nil, err
	}

	q := req.URL.Query()
	q.Add("binNumber", installmentParams.binNumber)
	q.Add("price", fmt.Sprintf("%d", installmentParams.price))
	req.URL.RawQuery = q.Encode()

	//req = req.WithContext(ctx)

	res := Response[Installment]{}
	if err := c.sendRequest(req, &res); err != nil {
		panic(err)
		return nil, err
	}

	return &res, nil
}
