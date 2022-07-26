package tests

import (
	"craftgate-go-client/adapter"
	"craftgate-go-client/model"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

var payByLink = adapter.PayByLink{
	Opts: model.RequestOptions{
		BaseURL:   "https://sandbox-api.craftgate.io",
		ApiKey:    "api-key",
		SecretKey: "secret-key",
	},
}

func Test_CreateProduct(t *testing.T) {
	request := adapter.CreateProductRequest{
		Name:                "A new Product",
		Channel:             "API",
		Price:               12.32,
		Currency:            model.Currency(model.TRY),
		EnabledInstallments: []int{1, 2, 3, 6},
	}

	res, err := payByLink.CreateProduct(request)
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func Test_UpdateProduct(t *testing.T) {
	request := adapter.UpdateProductRequest{
		Name:                "A new Product",
		Status:              model.Status(model.ACTIVE),
		Channel:             "API",
		Price:               12.32,
		Currency:            model.Currency(model.TRY),
		EnabledInstallments: []int{1, 2, 3, 6, 9},
	}

	res, err := payByLink.UpdateProduct(193, request)
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func Test_RetrieveProduct(t *testing.T) {
	res, err := payByLink.RetrieveProduct(193)
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func Test_DeleteProduct(t *testing.T) {
	err := payByLink.DeleteProduct(202)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func Test_SearchProducts(t *testing.T) {
	request := adapter.SearchProductsRequest{
		Page:     0,
		Size:     10,
		Currency: model.Currency(model.TRY),
	}

	res, err := payByLink.SearchProducts(request)
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}
