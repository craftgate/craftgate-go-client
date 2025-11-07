package adapter

import (
	"context"
	"fmt"
	"net/http"
)

type FileReporting struct {
	Client *Client
}

func (api *FileReporting) RetrieveDailyTransactionReport(ctx context.Context, request RetrieveDailyTransactionReportRequest) ([]byte, error) {
	newRequest, err := api.Client.NewRequestForByteResponse(ctx, http.MethodGet, "/file-reporting/v1/transaction-reports", request)
	if err != nil {
		return nil, err
	}

	response, err := api.Client.DoForByteResponse(ctx, newRequest)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (api *FileReporting) RetrieveDailyPaymentReport(ctx context.Context, request RetrieveDailyPaymentReportRequest) ([]byte, error) {
	newRequest, err := api.Client.NewRequestForByteResponse(ctx, http.MethodGet, "/file-reporting/v1/payment-reports", request)
	if err != nil {
		return nil, err
	}

	response, err := api.Client.DoForByteResponse(ctx, newRequest)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (api *FileReporting) CreateReport(ctx context.Context, request CreateReportRequest) (*ReportDemandResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/file-reporting/v1/report-demands", request)
	if err != nil {
		return nil, err
	}

	response := &Response[ReportDemandResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *FileReporting) RetrieveReport(ctx context.Context, retrieveReportRequest RetrieveReportRequest, reportId int64) ([]byte, error) {
	newRequest, err := api.Client.NewRequestForByteResponse(ctx, http.MethodGet, fmt.Sprintf("/file-reporting/v1/reports/%d", reportId), retrieveReportRequest)
	if err != nil {
		return nil, err
	}

	response, err := api.Client.DoForByteResponse(ctx, newRequest)
	if err != nil {
		return nil, err
	}

	return response, nil
}
