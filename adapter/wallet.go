package adapter

import (
	"craftgate-go-client/adapter/rest"
	"fmt"
	"net/http"
)

type Wallet struct {
	Opts RequestOptions
}

type RetrieveMemberWalletRequest struct {
	MemberId Long `schema:"memberId"`
}

type RetrieveMemberWalletResponse struct {
	Id               Long       `schema:"id,omitempty"`
	CreatedDate      Time       `schema:"createdDate,omitempty"`
	UpdatedDate      Time       `schema:"updatedDate,omitempty"`
	Amount           BigDecimal `schema:"amount,omitempty"`
	WithdrawalAmount BigDecimal `schema:"withdrawalAmount,omitempty"`
	Currency         Currency   `schema:"currency,omitempty"`
	MemberId         Long       `schema:"memberId,omitempty"`
}

type SearchWalletTransactionsRequest struct {
	WalletId              int64  `schema:"walletId"`
	WalletTransactionType string `schema:"walletTransactionType,omitempty"`
	Page                  int    `schema:"page,omitempty"`
	Size                  int    `schema:"size,omitempty"`
}

type SearchWalletTransactionsResponse struct {
	ID                    Long       `json:"id"`
	CreatedDate           Time       `json:"createdDate"`
	WalletTransactionType String     `json:"walletTransactionType"`
	Amount                BigDecimal `json:"amount"`
	TransactionID         Long       `json:"transactionId"`
	WalletID              Long       `json:"walletId"`
}

func (api *Wallet) RetrieveMemberWallet(request RetrieveMemberWalletRequest) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/wallet/v1/members/%d/wallet", api.Opts.BaseURL, request.MemberId.int64), nil)

	res := Response[RetrieveMemberWalletResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Wallet) SearchWalletTransactions(request SearchWalletTransactionsRequest) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/wallet/v1/wallets/%d/wallet-transactions", api.Opts.BaseURL, request.WalletId), nil)

	req.URL.RawQuery, _ = QueryParams(request)

	res := Response[DataResponse[SearchWalletTransactionsResponse]]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}
