package adapter

import (
	"context"
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
