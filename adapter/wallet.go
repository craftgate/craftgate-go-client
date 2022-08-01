package adapter

import (
	"craftgate-go-client/adapter/rest"
	"craftgate-go-client/model"
	"fmt"
	"net/http"
	"time"
)

type Wallet struct {
	Opts model.RequestOptions
}

type RetrieveMemberWalletRequest struct {
	MemberId int64
}

type RetrieveRemittanceRequest struct {
	RemittanceId int64
}

type WithdrawRequest struct {
	WithdrawId int64
}

type RetrieveRefundWalletTransactionRequest struct {
	WalletTransactionId int64
}

type RefundWalletTransactionToCardRequest struct {
	WalletTransactionId int64
	RefundPrice         float64 `json:"refundPrice"`
}

type RemittanceRequest struct {
	MemberId             int64   `json:"memberId"`
	Price                float64 `json:"price"`
	Description          string  `json:"description"`
	RemittanceReasonType string  `json:"remittanceReasonType"`
}

type CreateWithdrawRequest struct {
	MemberId    int64          `json:"memberId"`
	Price       float64        `json:"price"`
	Description string         `json:"description"`
	Currency    model.Currency `json:"currency"`
}

type SearchWalletTransactionsRequest struct {
	WalletId              int64
	WalletTransactionType string `schema:"walletTransactionType,omitempty"`
	Page                  int    `schema:"page,omitempty"`
	Size                  int    `schema:"size,omitempty"`
}

type SearchWithdrawsRequest struct {
	MemberId         int64          `schema:"walletId,omitempty"`
	PayoutStatus     string         `schema:"payoutStatus,omitempty"`
	Currency         model.Currency `schema:"currency,omitempty"`
	MinWithdrawPrice float64        `schema:"minWithdrawPrice,omitempty"`
	MaxWithdrawPrice float64        `schema:"maxWithdrawPrice,omitempty"`
	MinCreatedDate   time.Time      `schema:"minCreatedDate,omitempty"`
	MaxCreatedDate   time.Time      `schema:"maxCreatedDate,omitempty"`
	Page             int            `schema:"page,omitempty"`
	Size             int            `schema:"size,omitempty"`
}

type MemberWalletResponse struct {
	Id               int64          `json:"id"`
	CreatedDate      CraftgateTime  `json:"createdDate"`
	UpdatedDate      CraftgateTime  `json:"updatedDate"`
	Amount           float64        `json:"amount"`
	WithdrawalAmount float64        `json:"withdrawalAmount"`
	Currency         model.Currency `json:"currency"`
	MemberId         int64          `json:"memberId"`
}

type RemittanceResponse struct {
	Id                   int64         `json:"id"`
	CreatedDate          CraftgateTime `json:"createdDate"`
	Active               int           `json:"active"`
	Price                float64       `json:"price"`
	MemberId             int64         `json:"memberId"`
	RemittanceType       string        `json:"remittanceType"`
	RemittanceReasonType string        `json:"remittanceReasonType"`
	Description          string        `json:"description"`
}

type WithdrawResponse struct {
	Id           int64          `json:"id"`
	CreatedDate  CraftgateTime  `json:"createdDate"`
	Status       int            `json:"status"`
	MemberId     int64          `json:"memberId"`
	PayoutId     int64          `json:"payoutId"`
	Price        float64        `json:"price"`
	Description  string         `json:"description"`
	Currency     model.Currency `json:"currency"`
	PayoutStatus string         `json:"payoutStatus"`
}

type RefundWalletTransactionToCardResponse struct {
	Id                  int64         `json:"id"`
	CreatedDate         CraftgateTime `json:"createdDate"`
	RefundStatus        string        `json:"refundStatus"`
	RefundPrice         float64       `json:"refundPrice"`
	AuthCode            string        `json:"authCode"`
	HostReference       string        `json:"hostReference"`
	TransId             string        `json:"transId"`
	TransactionId       int64         `json:"transactionId"`
	WalletTransactionId int64         `json:"walletTransactionId"`
	PaymentError        string        `json:"paymentError"`
	TransactionType     string        `json:"transactionType"`
}

type SearchWalletTransactionsResponse struct {
	ID                    int64         `json:"id"`
	CreatedDate           CraftgateTime `json:"createdDate"`
	WalletTransactionType string        `json:"walletTransactionType"`
	Amount                float64       `json:"amount"`
	TransactionID         int64         `json:"transactionId"`
	WalletID              int64         `json:"walletId"`
}

type ResetMerchantMemberWalletBalanceRequest struct {
	WalletAmount float64 `json:"walletAmount"`
}

type RetrieveWalletTransactionRefundableAmountResponse struct {
	RefundableAmount float64 `json:"refundableAmount"`
}

func (api *Wallet) RetrieveMemberWallet(request RetrieveMemberWalletRequest) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/wallet/v1/members/%d/wallet", api.Opts.BaseURL, request.MemberId), nil)
	res := model.Response[MemberWalletResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Wallet) RetrieveMerchantMemberWallet() (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/wallet/v1/merchants/me/wallet", api.Opts.BaseURL), nil)
	res := model.Response[MemberWalletResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Wallet) ResetMerchantMemberWalletBalance(request ResetMerchantMemberWalletBalanceRequest) (interface{}, error) {
	body, _ := PrepareBody(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/wallet/v1/merchants/me/wallet/reset-balance", api.Opts.BaseURL), body)
	res := model.Response[MemberWalletResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Wallet) SearchWalletTransactions(request SearchWalletTransactionsRequest) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/wallet/v1/wallets/%d/wallet-transactions", api.Opts.BaseURL, request.WalletId), nil)
	req.URL.RawQuery, _ = QueryParams(request)
	res := model.Response[model.DataResponse[SearchWalletTransactionsResponse]]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Wallet) RetrieveRefundableAmountOfWalletTransaction(request RetrieveRefundWalletTransactionRequest) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/payment/v1/wallet-transactions/%d/refundable-amount", api.Opts.BaseURL, request.WalletTransactionId), nil)
	res := model.Response[RetrieveWalletTransactionRefundableAmountResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Wallet) RefundWalletTransactionToCard(request RefundWalletTransactionToCardRequest) (interface{}, error) {
	body, _ := PrepareBody(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/payment/v1/wallet-transactions/%d/refunds", api.Opts.BaseURL, request.WalletTransactionId), body)
	res := model.Response[RefundWalletTransactionToCardResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Wallet) RetrieveRefundWalletTransactionToCard(request RetrieveRefundWalletTransactionRequest) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/payment/v1/wallet-transactions/%d/refunds", api.Opts.BaseURL, request.WalletTransactionId), nil)
	res := model.Response[model.DataResponse[RefundWalletTransactionToCardResponse]]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Wallet) SendRemittance(request RemittanceRequest) (interface{}, error) {
	body, _ := PrepareBody(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/wallet/v1/remittances/send", api.Opts.BaseURL), body)
	res := model.Response[RemittanceResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Wallet) ReceiveRemittance(request RemittanceRequest) (interface{}, error) {
	body, _ := PrepareBody(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/wallet/v1/remittances/receive", api.Opts.BaseURL), body)
	res := model.Response[RemittanceResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Wallet) RetrieveRemittance(request RetrieveRemittanceRequest) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/wallet/v1/remittances/%d", api.Opts.BaseURL, request.RemittanceId), nil)
	req.URL.RawQuery, _ = QueryParams(request)
	res := model.Response[RemittanceResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Wallet) CreateWithdraw(request CreateWithdrawRequest) (interface{}, error) {
	body, _ := PrepareBody(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/wallet/v1/withdraws", api.Opts.BaseURL), body)
	res := model.Response[WithdrawResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Wallet) CancelWithdraw(request WithdrawRequest) (interface{}, error) {
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/wallet/v1/withdraws/%d/cancel", api.Opts.BaseURL, request.WithdrawId), nil)
	res := model.Response[WithdrawResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Wallet) RetrieveWithdraw(request WithdrawRequest) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/wallet/v1/withdraws/%d", api.Opts.BaseURL, request.WithdrawId), nil)
	res := model.Response[WithdrawResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Wallet) SearchWithdraws(request SearchWithdrawsRequest) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/wallet/v1/withdraws", api.Opts.BaseURL), nil)
	req.URL.RawQuery, _ = QueryParams(request)
	res := model.Response[model.DataResponse[RemittanceResponse]]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}
