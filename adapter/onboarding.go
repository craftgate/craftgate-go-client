package adapter

import (
	"context"
	"fmt"
	"net/http"
)

type Onboarding struct {
	Client *Client
}

func (api *Onboarding) CreateMember(ctx context.Context, request CreateMemberRequest) (*MemberResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/onboarding/v1/members", request)

	if err != nil {
		return nil, err
	}

	response := &Response[MemberResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

func (api *Onboarding) UpdateMember(ctx context.Context, id int64, request UpdateMemberRequest) (*MemberResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPut, fmt.Sprintf("/onboarding/v1/members/%d", id), request)
	if err != nil {
		return nil, err
	}

	response := &Response[MemberResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *Onboarding) RetrieveMember(ctx context.Context, id int64) (*MemberResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("/onboarding/v1/members/%d", id), nil)
	if err != nil {
		return nil, err
	}

	response := &Response[MemberResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *Onboarding) SearchMembers(ctx context.Context, request SearchMembersRequest) (*DataResponse[MemberResponse], error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, "/onboarding/v1/members", request)
	if err != nil {
		return nil, err
	}

	response := &Response[DataResponse[MemberResponse]]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *Onboarding) CreateMerchant(ctx context.Context, request CreateMerchantRequest) (*CreateMerchantResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/onboarding/v1/merchants", request)
	if err != nil {
		return nil, err
	}

	response := &Response[CreateMerchantResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
