package main

import (
	"fmt"
	"net/http"
)

type InstallmentApi struct {
	Opts RequestOptions
}

type SearchInstallmentRequest struct {
	BinNumber                               string
	Price                                   float64
	Currency                                string
	DistinctCardBrandsWithLowestCommissions bool
}

type InstallmentPrice struct {
	InstallmentPrice  float64 `json:"installmentPrice"`
	TotalPrice        float64 `json:"totalPrice"`
	InstallmentNumber int     `json:"installmentNumber"`
	InstallmentLabel  string  `json:"installmentLabel"`
}

type InstallmentResponse struct {
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

func (api *InstallmentApi) SearchInstallments(request SearchInstallmentRequest) (*Response[InstallmentResponse], error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/installment/v1/installments", api.Opts.BaseURL), nil)

	q := req.URL.Query()
	q.Add("binNumber", request.BinNumber)
	q.Add("price", fmt.Sprintf("%f", request.Price))
	req.URL.RawQuery = q.Encode()

	res := Response[InstallmentResponse]{}
	resErr := sendRequest(req, &res, api.Opts)
	return &res, resErr
}
