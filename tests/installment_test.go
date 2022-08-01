package tests

import (
	"craftgate-go-client/adapter"
	"craftgate-go-client/model"
	"fmt"
	"testing"
	"time"
)

var installment = adapter.Installment{
	Opts: model.RequestOptions{
		BaseURL:   "https://sandbox-api.craftgate.io",
		ApiKey:    "sandbox-SpqVrfuINfhbFtDEWBqQTCAhIzTEOedj",
		SecretKey: "sandbox-aJGxugIvDEdmgUYFByWAyNCrgaEpYWOw",
	},
}

func Test_BinCheck(t *testing.T) {
	request := adapter.CreateMemberRequest{
		MemberExternalId:          fmt.Sprintf("%d", time.Now().Nanosecond()),
		Name:                      "Haluk Demir",
		Email:                     "haluk.demir@example.com",
		PhoneNumber:               "905551111111",
		Address:                   "Suadiye Mah. Örnek Cd. No:23, 34740 Kadıköy/İstanbul",
		ContactName:               "Haluk",
		ContactSurname:            "Demir",
		NegativeWalletAmountLimit: -50,
		IsBuyer:                   true,
		IsSubMerchant:             false,
	}

	res, err := installment.RetrieveBinNumber("487074")
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}
