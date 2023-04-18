package adapter

import (
	"context"
	"fmt"
	"net/http"
)

type Wallet struct {
	Client *Client
}

func (api *Wallet) CreateMemberWallet(ctx context.Context, memberId int64, request CreateMemberWalletRequest) (*MemberWalletResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, fmt.Sprintf("/wallet/v1/members/%d/wallets", memberId), request)
	if err != nil {
		return nil, err
	}

	response := &Response[MemberWalletResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *Wallet) RetrieveMemberWallet(ctx context.Context, memberId int64) (*MemberWalletResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("/wallet/v1/members/%d/wallet", memberId), nil)
	if err != nil {
		return nil, err
	}

	response := &Response[MemberWalletResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *Wallet) RetrieveMerchantMemberWallet(ctx context.Context) (*MemberWalletResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, "/wallet/v1/merchants/me/wallet", nil)
	if err != nil {
		return nil, err
	}

	response := &Response[MemberWalletResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *Wallet) UpdateMemberWallet(ctx context.Context, memberId int64, walletId int64, request UpdateMemberWalletRequest) (*MemberWalletResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPut, fmt.Sprintf("/wallet/v1/members/%d/wallets/%d", memberId, walletId), request)
	if err != nil {
		return nil, err
	}

	response := &Response[MemberWalletResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *Wallet) ResetMerchantMemberWalletBalance(ctx context.Context, request ResetMerchantMemberWalletBalanceRequest) (*MemberWalletResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/wallet/v1/merchants/me/wallet/reset-balance", request)
	if err != nil {
		return nil, err
	}

	response := &Response[MemberWalletResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *Wallet) SearchWalletTransactions(ctx context.Context, walletId int64, request SearchWalletTransactionsRequest) (*DataResponse[SearchWalletTransactionsResponse], error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("/wallet/v1/wallets/%d/wallet-transactions", walletId), request)
	if err != nil {
		return nil, err
	}

	response := &Response[DataResponse[SearchWalletTransactionsResponse]]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

func (api *Wallet) RetrieveRefundableAmountOfWalletTransaction(ctx context.Context, walletTransactionId int64) (*RetrieveWalletTransactionRefundableAmountResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("/payment/v1/wallet-transactions/%d/refundable-amount", walletTransactionId), nil)
	if err != nil {
		return nil, err
	}

	response := &Response[RetrieveWalletTransactionRefundableAmountResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *Wallet) RefundWalletTransaction(ctx context.Context, walletTransactionId int64, request RefundWalletTransactionRequest) (*RefundWalletTransactionResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, fmt.Sprintf("/payment/v1/wallet-transactions/%d/refunds", walletTransactionId), request)
	if err != nil {
		return nil, err
	}

	response := &Response[RefundWalletTransactionResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *Wallet) RetrieveRefundWalletTransactions(ctx context.Context, walletTransactionId int64) (*DataResponse[RefundWalletTransactionResponse], error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("/payment/v1/wallet-transactions/%d/refunds", walletTransactionId), nil)
	if err != nil {
		return nil, err
	}

	response := &Response[DataResponse[RefundWalletTransactionResponse]]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *Wallet) SendRemittance(ctx context.Context, request RemittanceRequest) (*RemittanceResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/wallet/v1/remittances/send", request)
	if err != nil {
		return nil, err
	}

	response := &Response[RemittanceResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *Wallet) ReceiveRemittance(ctx context.Context, request RemittanceRequest) (*RemittanceResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/wallet/v1/remittances/receive", request)
	if err != nil {
		return nil, err
	}

	response := &Response[RemittanceResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *Wallet) RetrieveRemittance(ctx context.Context, remittanceId int64) (*RemittanceResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("/wallet/v1/remittances/%d", remittanceId), nil)
	if err != nil {
		return nil, err
	}

	response := &Response[RemittanceResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *Wallet) CreateWithdraw(ctx context.Context, request CreateWithdrawRequest) (*WithdrawResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/wallet/v1/withdraws", request)
	if err != nil {
		return nil, err
	}

	response := &Response[WithdrawResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *Wallet) CancelWithdraw(ctx context.Context, withdrawId int64) (interface{}, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, fmt.Sprintf("/wallet/v1/withdraws/%d/cancel", withdrawId), nil)
	if err != nil {
		return nil, err
	}

	response := &Response[WithdrawResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *Wallet) RetrieveWithdraw(ctx context.Context, withdrawId int64) (*WithdrawResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("/wallet/v1/withdraws/%d", withdrawId), nil)
	if err != nil {
		return nil, err
	}

	response := &Response[WithdrawResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *Wallet) SearchWithdraws(ctx context.Context, request SearchWithdrawsRequest) (*DataResponse[WithdrawResponse], error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, "/wallet/v1/withdraws", request)
	if err != nil {
		return nil, err
	}

	response := &Response[DataResponse[WithdrawResponse]]{}
	err = api.Client.Do(ctx, newRequest, response)

	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
