package tests

import (
    "context"
    "github.com/craftgate/craftgate-go-client/adapter"
    craftgate "github.com/craftgate/craftgate-go-client/adapter"
    "github.com/davecgh/go-spew/spew"
    "testing"
    "time"
)

var fileReportingClient, _ = craftgate.New("api-key", "secret-key", "https://sandbox-api.craftgate.io")

func TestFileReporting_RetrieveDailyTransactionReport(t *testing.T) {
    request := adapter.RetrieveDailyTransactionReportRequest{
        ReportDate: craftgate.DateOf(time.Date(2022, 11, 11, 0, 0, 0, 0, time.UTC)),
        FileType:   craftgate.CSV,
    }

    res, err := fileReportingClient.FileReporting.RetrieveDailyTransactionReport(context.Background(), request)
    _, _ = spew.Printf("%#v\n", res)

    if err != nil {
        t.Errorf("Error %s", err)
    }
}
