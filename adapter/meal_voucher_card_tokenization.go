package adapter

import (
    "context"
    "fmt"
    "net/http"
)

type MealVoucherCardTokenization struct {
	Client *Client
}

func (api *MealVoucherCardTokenization) Init(ctx context.Context, request MealVoucherCardTokenizationInitRequest) (*MealVoucherCardTokenizationInitResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/payment/v1/meal-voucher/card-tokenizations/init", request)

	if err != nil {
		return nil, err
	}

	response := &Response[MealVoucherCardTokenizationInitResponse]{}
	err = api.Client.Do(ctx, newRequest, response)

	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *MealVoucherCardTokenization) Regenerate(ctx context.Context, sessionId string, request MealVoucherCardTokenizationRegenerateRequest) (*MealVoucherCardTokenizationInitResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, fmt.Sprintf("/payment/v1/meal-voucher/card-tokenizations/%s/regenerate", sessionId), request)

	if err != nil {
		return nil, err
	}

	response := &Response[MealVoucherCardTokenizationInitResponse]{}
	err = api.Client.Do(ctx, newRequest, response)

	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *MealVoucherCardTokenization) Complete(ctx context.Context, sessionId string, request MealVoucherCardTokenizationCompleteRequest) (*MealVoucherCardTokenizationCompleteResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, fmt.Sprintf("/payment/v1/meal-voucher/card-tokenizations/%s/complete", sessionId), request)

	if err != nil {
		return nil, err
	}

	response := &Response[MealVoucherCardTokenizationCompleteResponse]{}
	err = api.Client.Do(ctx, newRequest, response)

	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

