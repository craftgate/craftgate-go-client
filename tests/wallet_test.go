package tests

import (
	"context"
	"github.com/craftgate/craftgate-go-client/adapter"
	craftgate "github.com/craftgate/craftgate-go-client/adapter"
	"github.com/davecgh/go-spew/spew"
	"testing"
	"time"
)

var walletClient, _ = craftgate.New("api-key", "secret-key", "https://sandbox-api.craftgate.io")

func TestWallet_RetrieveMemberWallet(t *testing.T) {
	res, err := walletClient.Wallet.RetrieveMemberWallet(context.Background(), 1)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestWallet_RetrieveMerchantMemberWallet(t *testing.T) {
	res, err := walletClient.Wallet.RetrieveMerchantMemberWallet(context.Background())
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestWallet_ResetMerchantMemberWalletBalance(t *testing.T) {
	request := adapter.ResetMerchantMemberWalletBalanceRequest{WalletAmount: -10}
	res, err := walletClient.Wallet.ResetMerchantMemberWalletBalance(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}

func TestWallet_SearchWalletTransactions(t *testing.T) {
	request := adapter.SearchWalletTransactionsRequest{}
	res, err := walletClient.Wallet.SearchWalletTransactions(context.Background(), 1, request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}

func TestWallet_RetrieveRefundableAmountOfWalletTransaction(t *testing.T) {
	res, err := walletClient.Wallet.RetrieveRefundableAmountOfWalletTransaction(context.Background(), 1)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}

func TestWallet_RefundWalletTransaction(t *testing.T) {
	request := adapter.RefundWalletTransactionRequest{RefundPrice: 10}
	res, err := walletClient.Wallet.RefundWalletTransaction(context.Background(), 1, request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}

func TestWallet_RetrieveRefundWalletTransactions(t *testing.T) {
	res, err := walletClient.Wallet.RetrieveRefundWalletTransactions(context.Background(), 1)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}

func TestWallet_SendRemittance(t *testing.T) {
	request := adapter.RemittanceRequest{
		Price:                100,
		MemberId:             1,
		Description:          "bonus",
		RemittanceReasonType: "REDEEM_ONLY_LOYALTY",
	}
	res, err := walletClient.Wallet.SendRemittance(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}

func TestWallet_ReceiveRemittance(t *testing.T) {
	request := adapter.RemittanceRequest{
		Price:                10,
		MemberId:             1,
		Description:          "bonus",
		RemittanceReasonType: "REDEEM_ONLY_LOYALTY",
	}
	res, err := walletClient.Wallet.ReceiveRemittance(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}

func TestWallet_RetrieveRemittance(t *testing.T) {
	res, err := walletClient.Wallet.RetrieveRemittance(context.Background(), 1)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}

func TestWallet_CreateWithdraw(t *testing.T) {
	request := adapter.CreateWithdrawRequest{
		Price:       5,
		MemberId:    1,
		Description: "Para çekme talebi",
		Currency:    "TRY",
	}
	res, err := walletClient.Wallet.CreateWithdraw(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}

func TestWallet_CancelWithdraw(t *testing.T) {
	res, err := walletClient.Wallet.CancelWithdraw(context.Background(), 1)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}

func TestWallet_RetrieveWithdraw(t *testing.T) {
	res, err := walletClient.Wallet.RetrieveWithdraw(context.Background(), 1)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}

func TestWallet_SearchWithdraws(t *testing.T) {
	request := adapter.SearchWithdrawsRequest{
		Currency:         "TRY",
		MinWithdrawPrice: 0,
		MaxWithdrawPrice: 10000,
		MinCreatedDate:   time.Now().AddDate(0, 0, -180),
		MaxCreatedDate:   time.Now(),
	}
	res, err := walletClient.Wallet.SearchWithdraws(context.Background(), request)

	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}

func TestWallet_CreateMemberWallet(t *testing.T) {
	request := adapter.CreateMemberWalletRequest{
		NegativeAmountLimit: 0,
		Currency:            "TRY",
	}
	res, err := walletClient.Wallet.CreateMemberWallet(context.Background(), 1, request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}
