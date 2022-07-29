package adapter

import (
	"craftgate-go-client/model"
	"fmt"
	"testing"
)

const (
	baseURL   = "https://sandbox-api.craftgate.io"
	apiKey    = "sandbox-YEhueLgomBjqsnvBlWVVuFsVhlvJlMHE"
	secretKey = "sandbox-tBdcdKVGmGupzfaWcULcwDLMoglZZvTz"
)

var wallet = Wallet{
	Opts: model.RequestOptions{
		BaseURL:   baseURL,
		ApiKey:    apiKey,
		SecretKey: secretKey,
	},
}

// test function
func TestWallet_RetrieveMemberWallet(t *testing.T) {
	res, err := wallet.RetrieveMemberWallet(RetrieveMemberWalletRequest{MemberId: *NewLong(66988)})
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestWallet_SearchWalletTransactions(t *testing.T) {
	res, err := wallet.SearchWalletTransactions(SearchWalletTransactionsRequest{WalletId: 62181})
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
}
