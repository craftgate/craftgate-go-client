package tests

import (
    "context"
    "testing"

    "github.com/craftgate/craftgate-go-client/adapter"
    "github.com/davecgh/go-spew/spew"
    "github.com/stretchr/testify/require"
)

var bkmExpressClient, _ = adapter.New("api-key", "secret-key", "https://sandbox-api.craftgate.io")

func TestBkm_Express_GenerateToken(t *testing.T) {
    request := adapter.BkmExpressGenerateTokenRequest{
       GsmNumber: "905555555555",
       UserId: "1234567890",
    }
    res, err := bkmExpressClient.BkmExpress.GenerateToken(context.Background(), request)
    _, _ = spew.Printf("%#v\n", res)

    require.Nil(t, err, "Error: %v", err)
}
