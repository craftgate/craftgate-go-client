package adapter

import (
	"context"
	"net/http"
)

type Masterpass struct {
	Client *Client
}

func (api *Masterpass) CheckMasterpassUser(ctx context.Context, request CheckMasterpassUserRequest) (*CheckMasterpassUserResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "payment/v1/masterpass-payments/check-user", request)
	if err != nil {
		return nil, err
	}

	response := &Response[CheckMasterpassUserResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *Masterpass) GenerateMasterpassPaymentToken(ctx context.Context, request MasterpassPaymentTokenGenerateRequest) (*MasterpassPaymentTokenGenerateResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "payment/v2/masterpass-payments/generate-token", request)
	if err != nil {
		return nil, err
	}

	response := &Response[MasterpassPaymentTokenGenerateResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *Masterpass) CompleteMasterpassPayment(ctx context.Context, request MasterpassPaymentCompleteRequest) (*PaymentResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "payment/v2/masterpass-payments/complete", request)
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

func (api *Masterpass) Init3DSMasterpassPayment(ctx context.Context, request MasterpassPaymentThreeDSInitRequest) (*MasterpassPaymentThreeDSInitResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "payment/v2/masterpass-payments/3ds-init", request)
	if err != nil {
		return nil, err
	}

	response := &Response[MasterpassPaymentThreeDSInitResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *Masterpass) Complete3DSMasterpassPayment(ctx context.Context, request MasterpassPaymentThreeDSCompleteRequest) (*PaymentResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "payment/v2/masterpass-payments/3ds-complete", request)
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

func (api *Masterpass) RetrieveLoyalties(ctx context.Context, request MasterpassRetrieveLoyaltiesRequest) (*RetrieveLoyaltiesResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/payment/v2/masterpass-payments/loyalties/retrieve", request)
	if err != nil {
		return nil, err
	}

	response := &Response[RetrieveLoyaltiesResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
