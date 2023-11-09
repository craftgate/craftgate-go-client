package adapter

import (
	"context"
	"fmt"
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

func (api *BankAccountTracking) RetrieveRecords(ctx context.Context, id int64) (*BankAccountTrackingRecordResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("/bank-account-tracking/v1/merchant-bank-account-trackings/records/%d", id), nil)
	if err != nil {
		return nil, err
	}
	response := &Response[BankAccountTrackingRecordResponse]{}
	err = api.Client.Do(ctx, newRequest, response)

	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
