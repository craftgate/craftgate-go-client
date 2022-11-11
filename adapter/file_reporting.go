package adapter

import (
	"context"
	"net/http"
)

type FileReporting struct {
	Client *Client
}

func (api *FileReporting) RetrieveDailyTransactionReport(ctx context.Context, request RetrieveDailyTransactionReportRequest) (*[]byte, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, "/file-reporting/v1/transaction-reports", request)
	if err != nil {
		return nil, err
	}

	response := &[]byte{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
