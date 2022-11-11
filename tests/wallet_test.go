package tests

import (
	"context"
	"craftgate-go-client/adapter"
	craftgate "craftgate-go-client/adapter"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

var walletClient, _ = craftgate.New("sandbox-SpqVrfuINfhbFtDEWBqQTCAhIzTEOedj", "sandbox-aJGxugIvDEdmgUYFByWAyNCrgaEpYWOw", "https://sandbox-api.craftgate.io")

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

//
//func TestWallet_SearchWalletTransactions(t *testing.T) {
//	res, err := wallet.SearchWalletTransactions(adapter.SearchWalletTransactionsRequest{WalletId: 62181})
//	_, _ = spew.Printf("%#v\n", res)
//
//	if err != nil {
//		t.Errorf("Error %s", err.Error())
//	}
//}
//
//func TestWallet_RetrieveRefundableAmountOfWalletTransaction(t *testing.T) {
//	res, err := wallet.RetrieveRefundableAmountOfWalletTransaction(137832)
//	_, _ = spew.Printf("%#v\n", res)
//
//	if err != nil {
//		t.Errorf("Error %s", err.Error())
//	}
//}
//
//func TestWallet_RefundWalletTransactionToCard(t *testing.T) {
//	res, err := wallet.RefundWalletTransactionToCard(adapter.RefundWalletTransactionToCardRequest{WalletTransactionId: 137832, RefundPrice: 3.25})
//	_, _ = spew.Printf("%#v\n", res)
//
//	if err != nil {
//		t.Errorf("Error %s", err.Error())
//	}
//}
//
//func TestWallet_RetrieveRefundWalletTransactionToCard(t *testing.T) {
//	res, err := wallet.RetrieveRefundWalletTransactionToCard(137832)
//	_, _ = spew.Printf("%#v\n", res)
//
//	if err != nil {
//		t.Errorf("Error %s", err.Error())
//	}
//}
//
//func TestWallet_SendRemittance(t *testing.T) {
//	res, err := wallet.SendRemittance(adapter.RemittanceRequest{
//		Price:                100,
//		MemberId:             68350,
//		Description:          "bonus",
//		RemittanceReasonType: "REDEEM_ONLY_LOYALTY",
//	})
//	_, _ = spew.Printf("%#v\n", res)
//
//	if err != nil {
//		t.Errorf("Error %s", err.Error())
//	}
//}
//
//func TestWallet_ReceiveRemittance(t *testing.T) {
//	res, err := wallet.ReceiveRemittance(adapter.RemittanceRequest{
//		Price:                10,
//		MemberId:             68350,
//		Description:          "bonus",
//		RemittanceReasonType: "REDEEM_ONLY_LOYALTY",
//	})
//	spew.Printf("%#v\n", res)
//
//	if err != nil {
//		t.Errorf("Error %s", err.Error())
//	}
//}
//
//func TestWallet_RetrieveRemittance(t *testing.T) {
//	res, err := wallet.RetrieveRemittance(adapter.RetrieveRemittanceRequest{RemittanceId: 66148})
//	spew.Printf("%#v\n", res)
//
//	if err != nil {
//		t.Errorf("Error %s", err.Error())
//	}
//}
//
//func TestWallet_CreateWithdraw(t *testing.T) {
//	res, err := wallet.CreateWithdraw(adapter.CreateWithdrawRequest{
//		Price:       5.25,
//		MemberId:    3,
//		Description: "Para çekme talebi",
//		Currency:    "TRY",
//	})
//	_, _ = spew.Printf("%#v\n", res)
//
//	if err != nil {
//		t.Errorf("Error %s", err.Error())
//	}
//}
//
//func TestWallet_CancelWithdraw(t *testing.T) {
//	res, err := wallet.CancelWithdraw(3)
//	_, _ = spew.Printf("%#v\n", res)
//
//	if err != nil {
//		t.Errorf("Error %s", err.Error())
//	}
//}
//
//func TestWallet_RetrieveWithdraw(t *testing.T) {
//	res, err := wallet.RetrieveWithdraw(3)
//	_, _ = spew.Printf("%#v\n", res)
//
//	if err != nil {
//		t.Errorf("Error %s", err.Error())
//	}
//}
//
//func TestWallet_SearchWithdraws(t *testing.T) {
//	res, err := wallet.SearchWithdraws(adapter.SearchWithdrawsRequest{
//		Currency:         "TRY",
//		MinWithdrawPrice: 0,
//		MaxWithdrawPrice: 10000,
//		MinCreatedDate:   time.Now().AddDate(0, 0, -180),
//		MaxCreatedDate:   time.Now(),
//	})
//
//	_, _ = spew.Printf("%#v\n", res)
//
//	if err != nil {
//		t.Errorf("Error %s", err.Error())
//	}
//}
