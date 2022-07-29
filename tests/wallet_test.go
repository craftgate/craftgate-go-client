package tests

import (
	"craftgate-go-client/adapter"
	"craftgate-go-client/model"
	"fmt"
	"testing"
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
	res, err := wallet.RetrieveMemberWallet(adapter.RetrieveMemberWalletRequest{MemberId: *adapter.NewLong(66988)})
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestWallet_SearchWalletTransactions(t *testing.T) {
	res, err := wallet.SearchWalletTransactions(adapter.SearchWalletTransactionsRequest{WalletId: 62181})
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}
