package adapter

import (
	"craftgate-go-client/adapter/rest"
	"craftgate-go-client/model"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/schema"
)

var encoder = schema.NewEncoder()

type InstallmentApi struct {
	Opts model.RequestOptions
}

type SearchInstallmentRequest struct {
	BinNumber                               string  `schema:"binNumber"`
	Price                                   float64 `schema:"price"`
	Currency                                string  `schema:"currency"`
	DistinctCardBrandsWithLowestCommissions bool    `schema:"distinctCardBrandsWithLowestCommissions"`
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

func (api *InstallmentApi) SearchInstallments(request SearchInstallmentRequest) (*model.Response[InstallmentResponse], error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/installment/v1/installments", api.Opts.BaseURL), nil)

	request.Currency = "TRY"
	queryParams := url.Values{}
	err := encoder.Encode(request, queryParams)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = queryParams.Encode()

	res := model.Response[InstallmentResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}
