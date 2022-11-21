package tests

import (
    "context"
    "github.com/craftgate/craftgate-go-client/v1/adapter"
    craftgate "github.com/craftgate/craftgate-go-client/v1/adapter"
    "github.com/davecgh/go-spew/spew"
    "testing"
)

var settlementClient, _ = craftgate.New("api-key", "secret-key", "https://sandbox-api.craftgate.io")

func TestSettlement_CreateInstantWalletSettlement(t *testing.T) {
    request := adapter.CreateInstantWalletSettlementRequest{ExcludedSubMerchantMemberIds: []int64{1, 2}}

    res, err := settlementClient.Settlement.CreateInstantWalletSettlement(context.Background(), request)
    _, _ = spew.Printf("%#v\n", res)

    if err != nil {
        t.Errorf("Error %s", err)
    }
}
