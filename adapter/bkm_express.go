package adapter

import (
	"context"
	"fmt"
	"net/http"
)

type BkmExpress struct {
	Client *Client
}

func (api *BkmExpress) Init(ctx context.Context, request InitBkmExpressRequest) (*InitBkmExpressResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/payment/v1/bkm-express/init", request)
	if err != nil {
		return nil, err
	}

	response := &Response[InitBkmExpressResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *BkmExpress) Complete(ctx context.Context, request CompleteBkmExpressRequest) (*PaymentResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/payment/v1/bkm-express/complete", request)
	if err != nil {
		return nil, err
	}

	response := &Response[PaymentResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *BkmExpress) RetrievePayment(ctx context.Context, ticketId string) error {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("/payment/v1/bkm-express/payments/%s", ticketId), nil)
	if err != nil {
		return fmt.Errorf("failed to create new request: %w", err)
	}
	response := &Void{}

	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	if response == nil {
		return fmt.Errorf("empty response")
	}

	return nil
}
