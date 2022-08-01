package adapter

import (
	"craftgate-go-client/adapter/rest"
	"craftgate-go-client/model"
	"fmt"
	"net/http"
)

type PayByLink struct {
	Opts model.RequestOptions
}

type CreateProductRequest struct {
	Name                string         `json:"name"`
	Channel             string         `json:"channel,omitempty"`
	Stock               int            `json:"stock,omitempty"`
	Price               float64        `json:"price"`
	Currency            model.Currency `json:"currency"`
	Description         string         `json:"description,omitempty"`
	EnabledInstallments []int          `json:"enabledInstallments"`
}

type UpdateProductRequest struct {
	Name                string         `json:"name"`
	Channel             string         `json:"channel,omitempty"`
	Stock               int            `json:"stock,omitempty"`
	Status              model.Status   `json:"status"`
	Price               float64        `json:"price"`
	Currency            model.Currency `json:"currency"`
	Description         string         `json:"description,omitempty"`
	EnabledInstallments []int          `json:"enabledInstallments"`
}

type SearchProductsRequest struct {
	Name     string         `schema:"name,omitempty"`
	MinPrice float64        `schema:"minPrice,omitempty"`
	MaxPrice float64        `schema:"maxPrice,omitempty"`
	Currency model.Currency `schema:"currency,omitempty"`
	Channel  string         `schema:"channel,omitempty"`
	Page     int            `schema:"page,omitempty"`
	Size     int            `schema:"size,omitempty"`
}

type ProductResponse struct {
	Id                  int64          `json:"id"`
	Name                string         `json:"name"`
	Description         string         `json:"description"`
	Status              model.Status   `json:"status"`
	Price               float64        `json:"price"`
	Currency            model.Currency `json:"currency"`
	Stock               int            `json:"stock"`
	SoldCount           int            `json:"soldCount"`
	Token               string         `json:"token"`
	EnabledInstallments []int          `json:"enabledInstallments"`
	Url                 string         `json:"url"`
	QrCodeUrl           string         `json:"qrCodeUrl"`
	Channel             string         `json:"channel"`
}

func (api *PayByLink) CreateProduct(request CreateProductRequest) (interface{}, error) {
	body, _ := PrepareBody(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/craftlink/v1/products", api.Opts.BaseURL), body)

	res := model.Response[ProductResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *PayByLink) UpdateProduct(id int64, request UpdateProductRequest) (interface{}, error) {
	body, _ := PrepareBody(request)
	req, _ := http.NewRequest("PUT", fmt.Sprintf("%s/craftlink/v1/products/%d", api.Opts.BaseURL, id), body)

	res := model.Response[ProductResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *PayByLink) RetrieveProduct(id int64) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/craftlink/v1/products/%d", api.Opts.BaseURL, id), nil)

	res := model.Response[ProductResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *PayByLink) DeleteProduct(id int64) error {
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/craftlink/v1/products/%d", api.Opts.BaseURL, id), nil)

	res := model.Void{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return resErr
}

func (api *PayByLink) SearchProducts(request SearchProductsRequest) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/craftlink/v1/products", api.Opts.BaseURL), nil)

	req.URL.RawQuery, _ = QueryParams(request)
	res := model.Response[model.DataResponse[ProductResponse]]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}
