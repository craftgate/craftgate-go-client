package tests

import (
	"craftgate-go-client/adapter"
	"craftgate-go-client/model"
	"fmt"
	"testing"
	"time"
)

var wallet = adapter.Wallet{
	Opts: model.RequestOptions{
		BaseURL:   "https://sandbox-api.craftgate.io",
		ApiKey:    "sandbox-YEhueLgomBjqsnvBlWVVuFsVhlvJlMHE",
		SecretKey: "sandbox-tBdcdKVGmGupzfaWcULcwDLMoglZZvTz",
	},
}

// test function
func TestWallet_RetrieveMemberWallet(t *testing.T) {
	res, err := wallet.RetrieveMemberWallet(adapter.RetrieveMemberWalletRequest{MemberId: 66988})
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

// test function
func TestWallet_RetrieveMerchantMemberWallet(t *testing.T) {
	res, err := wallet.RetrieveMerchantMemberWallet()
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

//?????????
func TestWallet_ResetMerchantMemberWalletBalance(t *testing.T) {
	res, err := wallet.ResetMerchantMemberWalletBalance(adapter.ResetMerchantMemberWalletBalanceRequest{WalletAmount: 50.25})
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}

func TestWallet_SearchWalletTransactions(t *testing.T) {
	res, err := wallet.SearchWalletTransactions(adapter.SearchWalletTransactionsRequest{WalletId: 62181})
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}

func TestWallet_RetrieveRefundableAmountOfWalletTransaction(t *testing.T) {
	res, err := wallet.RetrieveRefundableAmountOfWalletTransaction(adapter.RetrieveRefundWalletTransactionRequest{WalletTransactionId: 130459})
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}

//?????????
func TestWallet_RefundWalletTransactionToCard(t *testing.T) {
	res, err := wallet.RefundWalletTransactionToCard(adapter.RefundWalletTransactionToCardRequest{WalletTransactionId: 130459, RefundPrice: 3.25})
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}

func TestWallet_RetrieveRefundWalletTransactionToCard(t *testing.T) {
	res, err := wallet.RetrieveRefundWalletTransactionToCard(adapter.RetrieveRefundWalletTransactionRequest{WalletTransactionId: 130459})
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}

func TestWallet_SendRemittance(t *testing.T) {
	res, err := wallet.SendRemittance(adapter.RemittanceRequest{
		Price:                5.25,
		MemberId:             66988,
		Description:          "bonus",
		RemittanceReasonType: "REDEEM_ONLY_LOYALTY",
	})
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}

func TestWallet_ReceiveRemittance(t *testing.T) {
	res, err := wallet.ReceiveRemittance(adapter.RemittanceRequest{
		Price:                5.25,
		MemberId:             66988,
		Description:          "bonus",
		RemittanceReasonType: "REDEEM_ONLY_LOYALTY",
	})
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}

func TestWallet_RetrieveRemittance(t *testing.T) {
	res, err := wallet.RetrieveRemittance(adapter.RetrieveRemittanceRequest{RemittanceId: 64774})
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}

func TestWallet_CreateWithdraw(t *testing.T) {
	res, err := wallet.CreateWithdraw(adapter.CreateWithdrawRequest{
		Price:       5.25,
		MemberId:    66988,
		Description: "bonus",
		Currency:    "TRY",
	})
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}

func TestWallet_CancelWithdraw(t *testing.T) {
	res, err := wallet.CancelWithdraw(adapter.WithdrawRequest{WithdrawId: 55})
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}

func TestWallet_RetrieveWithdraw(t *testing.T) {
	res, err := wallet.RetrieveWithdraw(adapter.WithdrawRequest{WithdrawId: 55})
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}

func TestWallet_SearchWithdraws(t *testing.T) {
	t1 := time.Now().AddDate(0, 0, -180)
	t2 := time.Now()
	res, err := wallet.SearchWithdraws(adapter.SearchWithdrawsRequest{
		MemberId:         66988,
		Currency:         "TRY",
		MinWithdrawPrice: 10.75,
		MaxWithdrawPrice: 40.40,
		MinCreatedDate:   t1,
		MaxCreatedDate:   t2,
	})

	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}
