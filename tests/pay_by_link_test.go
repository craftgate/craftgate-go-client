package tests

import (
    "context"
    "github.com/craftgate/craftgate-go-client/adapter"
    craftgate "github.com/craftgate/craftgate-go-client/adapter"
    "github.com/davecgh/go-spew/spew"
    "testing"
)

var payByLinkClient, _ = craftgate.New("api-key", "secret-key", "https://sandbox-api.craftgate.io")

func Test_CreateProduct(t *testing.T) {
    request := adapter.CreateProductRequest{
        Name:                "A new Product",
        Channel:             "API",
        Price:               10,
        Currency:            craftgate.TRY,
        EnabledInstallments: []int{1, 2, 3, 6},
    }

    res, err := payByLinkClient.PayByLink.CreateProduct(context.Background(), request)
    _, _ = spew.Printf("%#v\n", res)

    if err != nil {
        t.Errorf("Error %s", err)
    }
}

func Test_UpdateProduct(t *testing.T) {
    request := adapter.UpdateProductRequest{
        Name:                "A new Product",
        Status:              craftgate.ACTIVE,
        Channel:             "API",
        Price:               10,
        Currency:            craftgate.TRY,
        EnabledInstallments: []int{1, 2, 3, 6, 9},
    }

    res, err := payByLinkClient.PayByLink.UpdateProduct(context.Background(), 1, request)
    _, _ = spew.Printf("%#v\n", res)

    if err != nil {
        t.Errorf("Error %s", err)
    }
}

func Test_RetrieveProduct(t *testing.T) {
    res, err := payByLinkClient.PayByLink.RetrieveProduct(context.Background(), 1)
    _, _ = spew.Printf("%#v\n", res)

    if err != nil {
        t.Errorf("Error %s", err)
    }
}

func Test_DeleteProduct(t *testing.T) {
    err := payByLinkClient.PayByLink.DeleteProduct(context.Background(), 1)

    if err != nil {
        t.Errorf("Error %s", err)
    }
}

func Test_SearchProducts(t *testing.T) {
    request := adapter.SearchProductsRequest{
        Page:     0,
        Size:     10,
        Currency: craftgate.TRY,
    }

    res, err := payByLinkClient.PayByLink.SearchProducts(context.Background(), request)
    _, _ = spew.Printf("%#v\n", res)

    if err != nil {
        t.Errorf("Error %s", err)
    }
}
