package adapter

import (
	"context"
	"net/http"
)

type MerchantApm struct {
	Client *Client
}

func (api *MerchantApm) RetrieveApms(ctx context.Context) (*MerchantApmListResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, "/merchant/v1/merchant-apms", nil)
	if err != nil {
		return nil, err
	}

	response := &Response[MerchantApmListResponse]{}
	respErr := api.Client.Do(ctx, newRequest, response)
	if respErr != nil {
		return nil, respErr
	}

	return response.Data, nil
}
