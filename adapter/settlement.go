package adapter

import (
	"craftgate-go-client/adapter/rest"
	"craftgate-go-client/model"
	"fmt"
	"net/http"
)

type Settlement struct {
	Opts model.RequestOptions
}

type CreateInstantWalletSettlementRequest struct {
	ExcludedSubMerchantMemberIds []int64
}

type CreateInstantWalletSettlementResponse struct {
	SettlementResultStatus string `json:"settlementResultStatus"`
}

func (api *Settlement) CreateInstantWalletSettlement(request CreateInstantWalletSettlementRequest) (interface{}, error) {
	body, _ := PrepareBody(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/settlement/v1/instant-wallet-settlements", api.Opts.BaseURL), body)
	res := model.Response[CreateInstantWalletSettlementResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}
