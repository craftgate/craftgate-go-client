package tests

import (
	"context"
	craftgate "github.com/craftgate/craftgate-go-client/adapter"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

var merchantApmClient, _ = craftgate.New("api-key", "secret-key", "https://sandbox-api.craftgate.io")

func Test_RetrieveApms(t *testing.T) {
	res, err := merchantApmClient.MerchantApm.RetrieveApms(context.Background())

	_, _ = spew.Printf("%#v\n", res)
	if err != nil {
		t.Errorf("Error %s", err)
	}
}
