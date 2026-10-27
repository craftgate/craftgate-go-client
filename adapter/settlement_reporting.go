package adapter

import (
	"context"
	"fmt"
	"net/http"
)

type SettlementReporting struct {
	Client *Client
}

func (api *SettlementReporting) SearchPayoutCompletedItems(ctx context.Context, request SearchPayoutCompletedItemsRequest) (*DataResponse[PayoutCompletedItem], error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, "/settlement-reporting/v1/settlement-file/payout-completed-items", request)
	if err != nil {
		return nil, err
	}

	response := &Response[DataResponse[PayoutCompletedItem]]{}
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

func (api *SettlementReporting) RetrieveSettlementRow(ctx context.Context, settlementRowId int64) (*SettlementRowResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("/settlement-reporting/v1/settlement-file/rows/%d", settlementRowId), nil)
	if err != nil {
		return nil, err
	}

	response := &Response[SettlementRowResponse]{}
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

func (api *SettlementReporting) RetrievePaymentTransactionPayout(ctx context.Context, paymentTransactionId int64) (*PaymentTxPayoutResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("/settlement-reporting/v1/payment-transactions/%d/payout", paymentTransactionId), nil)
	if err != nil {
		return nil, err
	}

	response := &Response[PaymentTxPayoutResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *SettlementReporting) SearchPaymentTransactionPayouts(ctx context.Context, request SearchPaymentTransactionPayoutsRequest) (*DataResponse[PaymentTxPayoutResponse], error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, "/settlement-reporting/v1/payment-transactions/payouts", request)
	if err != nil {
		return nil, err
	}

	response := &Response[DataResponse[PaymentTxPayoutResponse]]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
