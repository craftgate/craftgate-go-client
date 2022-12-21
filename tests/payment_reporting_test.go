package tests

import (
    "context"
    "github.com/craftgate/craftgate-go-client/adapter"
    craftgate "github.com/craftgate/craftgate-go-client/adapter"
    "github.com/davecgh/go-spew/spew"
    "testing"
    "time"
)

var paymentReportingClient, _ = craftgate.New("api-key", "secret-key", "https://sandbox-api.craftgate.io")

func Test_SearchPayment(t *testing.T) {
    request := adapter.SearchPaymentsRequest{
        Page: 0, Size: 10,
        PaymentType:    craftgate.CARD_PAYMENT,
        PaymentStatus:  craftgate.SUCCESS,
        Currency:       craftgate.TRY,
        MinCreatedDate: time.Now().AddDate(0, 0, -180),
        MaxCreatedDate: time.Now(),
    }
    res, err := paymentReportingClient.PaymentReporting.SearchPayments(context.Background(), request)
    _, _ = spew.Printf("%#v\n", res)

    if err != nil {
        t.Errorf("Error %s", err)
    }
}

func Test_SearchPaymentRefunds(t *testing.T) {
    request := adapter.SearchPaymentRefundsRequest{
        Page: 0, Size: 10,
        Currency:       craftgate.TRY,
        MinCreatedDate: time.Now().AddDate(0, 0, -180),
        MaxCreatedDate: time.Now(),
    }
    res, err := paymentReportingClient.PaymentReporting.SearchPaymentRefunds(context.Background(), request)
    _, _ = spew.Printf("%#v\n", res)

    if err != nil {
        t.Errorf("Error %s", err)
    }
}

func Test_SearchPaymentTransactionRefunds(t *testing.T) {
    request := adapter.SearchPaymentTransactionRefundsRequest{
        Page: 0, Size: 10,
        Currency:       craftgate.TRY,
        MinCreatedDate: time.Now().AddDate(0, 0, -180),
        MaxCreatedDate: time.Now(),
    }
    res, err := paymentReportingClient.PaymentReporting.SearchPaymentTransactionRefunds(context.Background(), request)
    _, _ = spew.Printf("%#v\n", res)

    if err != nil {
        t.Errorf("Error %s", err)
    }
}

func Test_RetrievePayment(t *testing.T) {
    res, err := paymentReportingClient.PaymentReporting.RetrievePayment(context.Background(), 123)
    _, _ = spew.Printf("%#v\n", res)

    if err != nil {
        t.Errorf("Error %s", err)
    }
}

func Test_RetrievePaymentTransactions(t *testing.T) {
    res, err := paymentReportingClient.PaymentReporting.RetrievePaymentTransactions(context.Background(), 123)
    _, _ = spew.Printf("%#v\n", res)

    if err != nil {
        t.Errorf("Error %s", err)
    }
}

func Test_RetrievePaymentRefunds(t *testing.T) {
    res, err := paymentReportingClient.PaymentReporting.RetrievePaymentRefunds(context.Background(), 123)
    _, _ = spew.Printf("%#v\n", res)

    if err != nil {
        t.Errorf("Error %s", err)
    }
}

func Test_RetrievePaymentTransactionRefunds(t *testing.T) {
    res, err := paymentReportingClient.PaymentReporting.RetrievePaymentTransactionRefunds(context.Background(), 123, 789)
    _, _ = spew.Printf("%#v\n", res)

    if err != nil {
        t.Errorf("Error %s", err)
    }
}
