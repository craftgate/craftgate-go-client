package adapter

import (
	"context"
	"fmt"
	"net/http"
)

type Merchant struct {
	Client *Client
}

func (api *Merchant) CreateMerchantPos(ctx context.Context, request CreateMerchantPosRequest) (*MerchantPosResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/merchant/v1/merchant-poses", request)

	if err != nil {
		return nil, err
	}

	response := &Response[MerchantPosResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

func (api *Merchant) RetrieveMerchantPos(ctx context.Context, id int64) (*MerchantPosResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("/merchant/v1/merchant-poses/%d", id), nil)

	if err != nil {
		return nil, err
	}

	response := &Response[MerchantPosResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

func (api *Merchant) SearchMerchantPos(ctx context.Context, request SearchMerchantPosRequest) (*DataResponse[MerchantPosResponse], error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, "/merchant/v1/merchant-poses", request)

	if err != nil {
		return nil, err
	}

	response := &Response[DataResponse[MerchantPosResponse]]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

func (api *Merchant) UpdateMerchantPosStatus(ctx context.Context, id int64, status PosStatus) error {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPut, fmt.Sprintf("/merchant/v1/merchant-poses/%d/status/%s", id, status), nil)

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

func (api *Merchant) DeleteMerchantPosStatus(ctx context.Context, id int64) error {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodDelete, fmt.Sprintf("/merchant/v1/merchant-poses/%d", id), nil)

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

func (api *Merchant) RetrieveMerchantPosCommissions(ctx context.Context, id int64) (*DataResponse[MerchantPosCommissionResponse], error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("/merchant/v1/merchant-poses/%d/commissions", id), nil)

	if err != nil {
		return nil, err
	}

	response := &Response[DataResponse[MerchantPosCommissionResponse]]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

func (api *Merchant) UpdateMerchantPosCommissions(ctx context.Context, id int64, request CreateMerchantPosCommissionRequest) (*DataResponse[MerchantPosCommissionResponse], error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, fmt.Sprintf("/merchant/v1/merchant-poses/%d/commissions", id), request)

	if err != nil {
		return nil, err
	}

	response := &Response[DataResponse[MerchantPosCommissionResponse]]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}
