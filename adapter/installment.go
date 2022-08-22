package adapter

import (
	"context"
	craftgate "craftgate-go-client"
	"craftgate-go-client/adapter/rest"
	"craftgate-go-client/model"
	"fmt"
	"net/http"
)

type Installment struct {
	client *craftgate.Client
}

type InstallmentPrice struct {
	InstallmentPrice  float64 `json:"installmentPrice"`
	TotalPrice        float64 `json:"totalPrice"`
	InstallmentNumber int     `json:"installmentNumber"`
	InstallmentLabel  string  `json:"installmentLabel"`
}

type SearchInstallmentsRequest struct {
	BinNumber                               string         `schema:"binNumber"`
	Price                                   float64        `schema:"price"`
	Currency                                model.Currency `schema:"currency"`
	DistinctCardBrandsWithLowestCommissions bool           `schema:"distinctCardBrandsWithLowestCommissions"`
}

type SearchInstallmentResponse struct {
	BinNumber         *string            `json:"binNumber"`
	Price             *float64           `json:"price"`
	CardType          *string            `json:"cardType"`
	CardAssociation   *string            `json:"cardAssociation"`
	CardBrand         *string            `json:"cardBrand"`
	BankName          *string            `json:"bankName"`
	BankCode          *int               `json:"bankCode"`
	Force3Ds          *bool              `json:"force3ds"`
	CvcRequired       *bool              `json:"cvcRequired"`
	Commercial        *bool              `json:"commercial"`
	PosAlias          *string            `json:"posAlias"`
	InstallmentPrices []InstallmentPrice `json:"installmentPrices"`
}

type RetrieveBinNumberRequest struct {
	BinNumber string
}

type RetrieveBinNumberResponse struct {
	BinNumber       *string `json:"binNumber"`
	CardType        *string `json:"cardType"`
	CardAssociation *string `json:"cardAssociation"`
	CardBrand       *string `json:"cardBrand"`
	BankName        *string `json:"bankName"`
	BankCode        *int    `json:"bankCode"`
	Commercial      *bool   `json:"commercial"`
}

func (api *Installment) SearchInstallments(ctx context.Context, request SearchInstallmentsRequest) (*SearchInstallmentResponse, error) {

	newRequest, err := api.client.NewRequest(ctx, http.MethodGet, "/installment/v1/installments", request)
	if err != nil {
		return nil, err
	}
	installmentResponse := new(SearchInstallmentResponse)
	_, respErr := api.client.Do(ctx, newRequest, installmentResponse)
	if err != nil {
		return nil, respErr
	}

	return installmentResponse, err
}

func (api *Installment) RetrieveBinNumber(binNumber string) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/installment/v1/bins/%s", api.Opts.BaseURL, binNumber), nil)

	res := model.Response[RetrieveBinNumberResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}
