package adapter

import (
	"craftgate-go-client/adapter/rest"
	"craftgate-go-client/model"
	"fmt"
	"net/http"
)

type Wallet struct {
	Opts model.RequestOptions
}

type RetrieveMemberWalletRequest struct {
	MemberId int64 `schema:"memberId"`
}

type RetrieveMemberWalletResponse struct {
	Id               int64          `schema:"id"`
	CreatedDate      Time           `schema:"createdDate"`
	UpdatedDate      Time           `schema:"updatedDate"`
	Amount           float64        `schema:"amount"`
	WithdrawalAmount float64        `schema:"withdrawalAmount"`
	Currency         model.Currency `schema:"currency"`
	MemberId         int64          `schema:"memberId"`
}

func (api *Wallet) RetrieveMemberWallet(request RetrieveMemberWalletRequest) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/wallet/v1/members/%d/wallet", api.Opts.BaseURL, request.MemberId), nil)

	res := model.Response[RetrieveMemberWalletResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}
