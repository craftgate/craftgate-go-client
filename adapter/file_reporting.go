package adapter

import (
	"context"
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
