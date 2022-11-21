package tests

import (
    "context"
    "github.com/craftgate/craftgate-go-client/v1/adapter"
    craftgate "github.com/craftgate/craftgate-go-client/v1/adapter"
    "github.com/davecgh/go-spew/spew"
    "testing"
)

var fraudClient, _ = craftgate.New("api-key", "secret-key", "https://sandbox-api.craftgate.io")

func Test_SearchFraudChecks(t *testing.T) {
    request := adapter.SearchFraudChecksRequest{
        Page: 0, Size: 10,
    }

    res, err := fraudClient.Fraud.SearchFraudChecks(context.Background(), request)
    _, _ = spew.Printf("%#v\n", res)

    if err != nil {
        t.Errorf("Error %s", err)
    }
}

func Test_RetrieveAllFraudValueList(t *testing.T) {
    res, err := fraudClient.Fraud.RetrieveAllFraudValueList(context.Background())
    _, _ = spew.Printf("%#v\n", res)

    if err != nil {
        t.Errorf("Error %s", err)
    }
}

func Test_RetrieveFraudValueList(t *testing.T) {
    res, err := fraudClient.Fraud.RetrieveFraudValueList(context.Background(), "blockedBuyerIdList")
    _, _ = spew.Printf("%#v\n", res)

    if err != nil {
        t.Errorf("Error %s", err)
    }
}

func Test_CreateFraudValueList(t *testing.T) {
    err := fraudClient.Fraud.CreateFraudValueList(context.Background(), "myTestList")

    if err != nil {
        t.Errorf("Error %s", err)
    }
}

func Test_DeleteFraudValueList(t *testing.T) {
    err := fraudClient.Fraud.DeleteFraudValueList(context.Background(), "myTestList")

    if err != nil {
        t.Errorf("Error %s", err)
    }
}

func Test_AddValueToFraudValueList(t *testing.T) {
    request := adapter.FraudValueListRequest{
        ListName: "ipList",
        Value:    "999.999.999.999",
    }
    err := fraudClient.Fraud.AddValueToFraudValueList(context.Background(), request)

    if err != nil {
        t.Errorf("Error %s", err)
    }
}

func Test_RemoveValueFromFraudValueList(t *testing.T) {
    err := fraudClient.Fraud.RemoveValueFromFraudValueList(context.Background(), "ipList", "999.999.999.999")

    if err != nil {
        t.Errorf("Error %s", err)
    }
}
