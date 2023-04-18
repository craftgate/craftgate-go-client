package adapter

import (
	"context"
	"fmt"
	"net/http"
)

type SettlementReporting struct {
	Client *Client
}

func (api *SettlementReporting) SearchPayoutCompletedTransactions(ctx context.Context, request SearchPayoutCompletedTransactionsRequest) (*DataResponse[SearchPayoutCompletedTransactionsResponse], error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, "/settlement-reporting/v1/settlement-file/payout-completed-transactions", request)
	if err != nil {
		return nil, err
	}

	response := &Response[DataResponse[SearchPayoutCompletedTransactionsResponse]]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *SettlementReporting) SearchPayoutBouncedTransactions(ctx context.Context, request SearchPayoutBouncedTransactionsRequest) (*DataResponse[SearchPayoutBouncedTransactionsResponse], error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, "/settlement-reporting/v1/settlement-file/bounced-sub-merchant-rows", request)
	if err != nil {
		return nil, err
	}

	response := &Response[DataResponse[SearchPayoutBouncedTransactionsResponse]]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *SettlementReporting) RetrievePayoutDetails(ctx context.Context, payoutDetailId int64) (*PayoutDetailResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("/settlement-reporting/v1/settlement-file/payout-details/%d", payoutDetailId), nil)
	if err != nil {
		return nil, err
	}

	response := &Response[PayoutDetailResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *SettlementReporting) SearchPayoutRows(ctx context.Context, request SearchPayoutRowRequest) (*DataResponse[PayoutRowResponse], error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, "/settlement-reporting/v1/settlement-file-rows", request)
	if err != nil {
		return nil, err
	}

	response := &Response[DataResponse[PayoutRowResponse]]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
