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

func (api *Settlement) CreatePayoutAccount(ctx context.Context, request CreatePayoutAccountRequest) (*PayoutAccountResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/settlement/v1/payout-accounts", request)
	if err != nil {
		return nil, err
	}

	response := &Response[PayoutAccountResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *Settlement) UpdatePayoutAccount(ctx context.Context, id int64, request UpdatePayoutAccountRequest) error {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPut, fmt.Sprintf("/settlement/v1/payout-accounts/%d", id), request)
	if err != nil {
		return err
	}

	response := &Void{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return err
	}

	return nil
}

func (api *Settlement) DeletePayoutAccount(ctx context.Context, id int64) error {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodDelete, fmt.Sprintf("/settlement/v1/payout-accounts/%d", id), nil)
	if err != nil {
		return err
	}

	response := &Void{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return err
	}

	return nil
}

func (api *Settlement) SearchPayoutAccounts(ctx context.Context, request SearchPayoutAccountRequest) (*DataResponse[PayoutAccountResponse], error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, "/settlement/v1/payout-accounts", request)
	if err != nil {
		return nil, err
	}

	response := &Response[DataResponse[PayoutAccountResponse]]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
