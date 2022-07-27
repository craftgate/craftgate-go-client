package adapter

import (
	"craftgate-go-client/adapter/rest"
	"craftgate-go-client/model"
	"fmt"
	"net/http"
)

type Installment struct {
	Opts model.RequestOptions
}

type InstallmentPrice struct {
	InstallmentPrice  float64 `json:"installmentPrice"`
	TotalPrice        float64 `json:"totalPrice"`
	InstallmentNumber int     `json:"installmentNumber"`
	InstallmentLabel  string  `json:"installmentLabel"`
}

type SearchInstallmentRequest struct {
	BinNumber                               string         `schema:"binNumber"`
	Price                                   float64        `schema:"price"`
	Currency                                model.Currency `schema:"currency"`
	DistinctCardBrandsWithLowestCommissions bool           `schema:"distinctCardBrandsWithLowestCommissions"`
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

type RetrieveBinNumberRequest struct {
	BinNumber string
}

type RetrieveBinNumberResponse struct {
	BinNumber       string `json:"binNumber"`
	CardType        string `json:"cardType"`
	CardAssociation string `json:"cardAssociation"`
	CardBrand       string `json:"cardBrand"`
	BankName        string `json:"bankName"`
	BankCode        int    `json:"bankCode"`
	Commercial      bool   `json:"commercial"`
}

func (api *Installment) SearchInstallments(request SearchInstallmentRequest) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/installment/v1/installments", api.Opts.BaseURL), nil)

	request.Currency = "TRY"
	req.URL.RawQuery, _ = QueryParams(request)

	res := model.Response[model.DataResponse[InstallmentResponse]]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Installment) RetrieveBinNumber(request RetrieveBinNumberRequest) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/installment/v1/bins/%s", api.Opts.BaseURL, request.BinNumber), nil)

	res := model.Response[RetrieveBinNumberResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}
