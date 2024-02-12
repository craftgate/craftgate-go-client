package adapter

import (
	"context"
	"fmt"
	"net/http"
)

type Juzdan struct {
	Client *Client
}

func (api *Juzdan) InitJuzdanPayment(ctx context.Context, request InitJuzdanPaymentRequest) (*InitJuzdanPaymentResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/payment/v1/juzdan-payments/init", request)
	if err != nil {
		return nil, err
	}

	response := &Response[InitJuzdanPaymentResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *Juzdan) RetrieveJuzdanPayment(ctx context.Context, referenceId string) (*PaymentResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("/installment/v1/bins/%s", referenceId), nil)

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
