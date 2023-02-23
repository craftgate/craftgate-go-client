package adapter

import (
	"context"
	"net/http"
)

type Merchant struct {
	Client *Client
}

func (api *Merchant) CreateMerchantPos(ctx context.Context, request CreateMerchantPosRequest) (*CreateMerchantPosResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/settlement/v1/instant-wallet-settlements", request)
	if err != nil {
		return nil, err
	}

	response := &Response[CreateMerchantPosResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
