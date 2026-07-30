package adapter

import (
	"context"
	"fmt"
	"net/http"
)

type Installment struct {
	Client *Client
}

func (api *Installment) SearchInstallments(ctx context.Context, request SearchInstallmentsRequest) (*InstallmentListResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, "/installment/v1/installments", request)
	if err != nil {
		return nil, err
	}
	response := &Response[InstallmentListResponse]{}
	respErr := api.Client.Do(ctx, newRequest, response)
	if respErr != nil {
		return nil, respErr
	}

	return response.Data, nil
}

func (api *Installment) RetrieveBinNumber(ctx context.Context, binNumber string) (*RetrieveBinNumberResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("/installment/v1/bins/%s", binNumber), nil)

	if err != nil {
		return nil, err
	}

	response := &Response[RetrieveBinNumberResponse]{}
	respErr := api.Client.Do(ctx, newRequest, response)
	if respErr != nil {
		return nil, respErr
	}

	return response.Data, nil
}
