package adapter

import (
	"context"
	"net/http"
)

type BankAccountTracking struct {
	Client *Client
}

func (api *BankAccountTracking) SearchRecords(ctx context.Context, request SearchBankAccountTrackingRecordRequest) (*DataResponse[BankAccountTrackingRecordResponse], error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, "/bank-account-tracking/v1/merchant-bank-account-trackings/records", request)
	if err != nil {
		return nil, err
	}
	response := &Response[DataResponse[BankAccountTrackingRecordResponse]]{}
	err = api.Client.Do(ctx, newRequest, response)

	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
