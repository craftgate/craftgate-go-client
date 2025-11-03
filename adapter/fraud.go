package adapter

import (
	"context"
	"net/http"
)

type Fraud struct {
	Client *Client
}

func (api *Fraud) SearchFraudChecks(ctx context.Context, request SearchFraudChecksRequest) (*DataResponse[FraudCheckResponse], error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, "/fraud/v1/fraud-checks", request)
	if err != nil {
		return nil, err
	}

	response := &Response[DataResponse[FraudCheckResponse]]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *Fraud) RetrieveAllFraudValueList(ctx context.Context) (*DataResponse[FraudValuesResponse], error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, "/fraud/v1/value-lists/all", nil)
	if err != nil {
		return nil, err
	}

	response := &Response[DataResponse[FraudValuesResponse]]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *Fraud) RetrieveFraudValueList(ctx context.Context, listName string) (*FraudValuesResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, "/fraud/v1/value-lists/"+listName, nil)
	if err != nil {
		return nil, err
	}

	response := &Response[FraudValuesResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *Fraud) CreateFraudValueList(ctx context.Context, listName string, fraudValueType FraudValueType) error {
	request := FraudValueListRequest{
		ListName: listName,
		Type:     fraudValueType,
	}

	newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/fraud/v1/value-lists", request)
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

func (api *Fraud) DeleteFraudValueList(ctx context.Context, listName string) error {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodDelete, "/fraud/v1/value-lists/"+listName, nil)
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

func (api *Fraud) AddCardValueToFraudValueList(ctx context.Context, request AddCardFingerprintFraudValueListRequest, listName string) error {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/fraud/v1/value-lists/"+listName+"/card-fingerprints", request)
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

func (api *Fraud) AddValueToFraudValueList(ctx context.Context, request FraudValueListRequest) error {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/fraud/v1/value-lists", request)
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

func (api *Fraud) RemoveValueFromFraudValueList(ctx context.Context, listName, valueId string) error {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodDelete, "/fraud/v1/value-lists/"+listName+"/values/"+valueId, nil)
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
