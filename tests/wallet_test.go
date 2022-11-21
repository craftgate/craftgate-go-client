package tests

import (
	"context"
	"craftgate-go-client/adapter"
	craftgate "craftgate-go-client/adapter"
	"github.com/davecgh/go-spew/spew"
	"testing"
	"time"
)

var walletClient, _ = craftgate.New("api-key", "secret-key", "https://sandbox-api.craftgate.io")

func TestWallet_RetrieveMemberWallet(t *testing.T) {
	res, err := walletClient.Wallet.RetrieveMemberWallet(context.Background(), 68350)
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
	request := adapter.ResetMerchantMemberWalletBalanceRequest{WalletAmount: -190}
	res, err := walletClient.Wallet.ResetMerchantMemberWalletBalance(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}

func TestWallet_SearchWalletTransactions(t *testing.T) {
	request := adapter.SearchWalletTransactionsRequest{}
	res, err := walletClient.Wallet.SearchWalletTransactions(context.Background(), 81686, request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}

func TestWallet_RetrieveRefundableAmountOfWalletTransaction(t *testing.T) {
	res, err := walletClient.Wallet.RetrieveRefundableAmountOfWalletTransaction(context.Background(), 172485)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}

func TestWallet_RefundWalletTransactionToCard(t *testing.T) {
	request := adapter.RefundWalletTransactionToCardRequest{RefundPrice: 3.25}
	res, err := walletClient.Wallet.RefundWalletTransactionToCard(context.Background(), 172485, request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}

func TestWallet_RetrieveRefundWalletTransactionsToCard(t *testing.T) {
	res, err := walletClient.Wallet.RetrieveRefundWalletTransactionsToCard(context.Background(), 172485)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}

func TestWallet_SendRemittance(t *testing.T) {
	request := adapter.RemittanceRequest{
		Price:                100,
		MemberId:             86816,
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
		MemberId:             86816,
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
	res, err := walletClient.Wallet.RetrieveRemittance(context.Background(), 82068)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}

func TestWallet_CreateWithdraw(t *testing.T) {
	request := adapter.CreateWithdrawRequest{
		Price:       5.25,
		MemberId:    86747,
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
	res, err := walletClient.Wallet.CancelWithdraw(context.Background(), 10)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}

func TestWallet_RetrieveWithdraw(t *testing.T) {
	res, err := walletClient.Wallet.RetrieveWithdraw(context.Background(), 9)
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
