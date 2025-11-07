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
		FileType:   craftgate.ReportFileType_CSV,
	}

	res, err := fileReportingClient.FileReporting.RetrieveDailyTransactionReport(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestFileReporting_RetrieveDailyPaymentReport(t *testing.T) {
	request := adapter.RetrieveDailyPaymentReportRequest{
		ReportDate: craftgate.DateOf(time.Date(2025, 1, 26, 0, 0, 0, 0, time.UTC)),
		FileType:   craftgate.ReportFileType_CSV,
	}

	res, err := fileReportingClient.FileReporting.RetrieveDailyPaymentReport(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestFileReporting_CreateReportDemand(t *testing.T) {
	startDate := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.UTC).Add(72 * time.Hour)
	endDate := startDate.Add(24 * time.Hour)

	request := adapter.CreateReportRequest{
		ReportType:   craftgate.ReportType_TRANSACTION,
		StartDate:    startDate,
		EndDate:      endDate,
		ReportPeriod: craftgate.ReportPeriod_INSTANT,
	}

	res, err := fileReportingClient.FileReporting.CreateReport(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestFileReporting_RetrieveReport(t *testing.T) {
	retrieveReportRequest := adapter.RetrieveReportRequest{
		FileType: craftgate.ReportFileType_CSV,
	}
	reportId := int64(22)
	res, err := fileReportingClient.FileReporting.RetrieveReport(context.Background(), retrieveReportRequest, reportId)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}
