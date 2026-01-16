package tests

import (
    "context"
    "github.com/craftgate/craftgate-go-client/adapter"
    craftgate "github.com/craftgate/craftgate-go-client/adapter"
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

func Test_SearchFraudRule(t *testing.T) {
    request := adapter.SearchFraudRuleRequest{
        Page: 0, Size: 10,
    }

    res, err := fraudClient.Fraud.SearchFraudRule(context.Background(), request)
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
	err := fraudClient.Fraud.CreateFraudValueList(context.Background(), "ipList", craftgate.FraudValueType_IP)

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
		Type:     craftgate.FraudValueType_IP,
		Label:    "local ip 1",
		Value:    "127.0.0.1",
	}
	err := fraudClient.Fraud.AddValueToFraudValueList(context.Background(), request)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func Test_AddTemporaryValueToFraudValueList(t *testing.T) {
	request := adapter.FraudValueListRequest{
		ListName:          "ipList",
		Type:              craftgate.FraudValueType_IP,
		Label:             "local ip 2",
		Value:             "127.0.0.2",
		DurationInSeconds: 60,
	}

	err := fraudClient.Fraud.AddValueToFraudValueList(context.Background(), request)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func Test_AddCardFingerprintToFraudValueList(t *testing.T) {
	request := adapter.AddCardFingerprintFraudValueListRequest{
		Label:             "John Doe's Card",
		Operation:         craftgate.FraudOperation_PAYMENT,
		OperationId:       "11675", //PaymentId
		DurationInSeconds: 3600,
	}
	err := fraudClient.Fraud.AddCardFingerprintToFraudValueList(context.Background(), request, "cardList")

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func Test_RemoveValueFromFraudValueList(t *testing.T) {
	err := fraudClient.Fraud.RemoveValueFromFraudValueList(context.Background(), "ipList", "7aac0ad8-d170-4c2b-89d3-440fcf145b35")

	if err != nil {
		t.Errorf("Error %s", err)
	}
}
