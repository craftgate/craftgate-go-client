package adapter

import (
	"context"
	"fmt"
	"net/http"
)

type PayByLink struct {
	Client *Client
}

func (api *PayByLink) CreateProduct(ctx context.Context, request CreateProductRequest) (*ProductResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/craftlink/v1/products", request)
	if err != nil {
		return nil, err
	}

	response := &Response[ProductResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *PayByLink) UpdateProduct(ctx context.Context, id int64, request UpdateProductRequest) (*ProductResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPut, fmt.Sprintf("/craftlink/v1/products/%d", id), request)
	if err != nil {
		return nil, err
	}

	response := &Response[ProductResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *PayByLink) RetrieveProduct(ctx context.Context, id int64) (interface{}, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("/craftlink/v1/products/%d", id), nil)
	if err != nil {
		return nil, err
	}

	response := &Response[ProductResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *PayByLink) DeleteProduct(ctx context.Context, id int64) error {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodDelete, fmt.Sprintf("/craftlink/v1/products/%d", id), nil)
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

func (api *PayByLink) SearchProducts(ctx context.Context, request SearchProductsRequest) (*DataResponse[ProductResponse], error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, "/craftlink/v1/products", request)
	if err != nil {
		return nil, err
	}

	response := &Response[DataResponse[ProductResponse]]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
