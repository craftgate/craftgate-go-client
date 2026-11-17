package adapter

import (
	"context"
	"fmt"
	"net/http"
)

type Settlement struct {
	Client *Client
}

func (api *Settlement) CreateInstantWalletSettlement(ctx context.Context, request CreateInstantWalletSettlementRequest) (*CreateInstantWalletSettlementResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/settlement/v1/instant-wallet-settlements", request)
	if err != nil {
		return nil, err
	}

	response := &Response[CreateInstantWalletSettlementResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *Settlement) ApprovePaymentTransactions(ctx context.Context, request PaymentTransactionsApprovalRequest) (*DataResponse[PaymentTransactionsApprovalResponse], error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/settlement/v1/payment-transactions/approve", request)
	if err != nil {
		return nil, err
	}

	response := &Response[DataResponse[PaymentTransactionsApprovalResponse]]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *Settlement) DisapprovePaymentTransactions(ctx context.Context, request PaymentTransactionsApprovalRequest) (*DataResponse[PaymentTransactionsApprovalResponse], error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/settlement/v1/payment-transactions/disapprove", request)
	if err != nil {
		return nil, err
	}

	response := &Response[DataResponse[PaymentTransactionsApprovalResponse]]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *Settlement) UpdatePaymentTransaction(ctx context.Context, id int64, request UpdatePaymentTransactionRequest) (*PaymentTxPayoutResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPut, fmt.Sprintf("/settlement/v1/payment-transactions/%d", id), request)
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
