package tests

import (
	"craftgate-go-client/adapter"
	"craftgate-go-client/model"
	"fmt"
	"testing"
	"time"
)

var onboarding = adapter.Onboarding{
	Opts: model.RequestOptions{
		BaseURL:   "https://sandbox-api.craftgate.io",
		ApiKey:    "sandbox-SpqVrfuINfhbFtDEWBqQTCAhIzTEOedj",
		SecretKey: "sandbox-aJGxugIvDEdmgUYFByWAyNCrgaEpYWOw",
	},
}

func Test_CreateBuyerMember(t *testing.T) {
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

	res, err := onboarding.CreateMember(request)
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func Test_CreateSubMerchantMember(t *testing.T) {
	request := adapter.CreateMemberRequest{
		MemberExternalId:  fmt.Sprintf("%d", time.Now().Nanosecond()),
		ContactName:       "Haluk",
		ContactSurname:    "Demir",
		Email:             "haluk.demir@example.com",
		PhoneNumber:       "905551111111",
		Iban:              "TR930006701000000001111111",
		IdentityNumber:    "11111111110",
		LegalCompanyTitle: "Dem Zeytinyağı Üretim Ltd. Şti.",
		Name:              "Dem Zeytinyağı Üretim Ltd. Şti.",
		MemberType:        model.MemberType(model.LIMITED_OR_JOINT_STOCK_COMPANY),
		TaxNumber:         "1111111114",
		TaxOffice:         "Erenköy",
		Address:           "Suadiye Mah. Örnek Cd. No:23, 34740 Kadıköy/İstanbul",
		IsBuyer:           false,
		IsSubMerchant:     true,
	}

	res, err := onboarding.CreateMember(request)
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func Test_CreateBuyerAndSubMerchantMember(t *testing.T) {
	request := adapter.CreateMemberRequest{
		MemberExternalId:  fmt.Sprintf("%d", time.Now().Nanosecond()),
		ContactName:       "Haluk",
		ContactSurname:    "Demir",
		Email:             "haluk.demir@example.com",
		PhoneNumber:       "905551111111",
		Iban:              "TR930006701000000001111111",
		IdentityNumber:    "11111111110",
		LegalCompanyTitle: "Dem Zeytinyağı Üretim Ltd. Şti.",
		Name:              "Dem Zeytinyağı Üretim Ltd. Şti.",
		MemberType:        model.MemberType(model.LIMITED_OR_JOINT_STOCK_COMPANY),
		TaxNumber:         "1111111114",
		TaxOffice:         "Erenköy",
		Address:           "Suadiye Mah. Örnek Cd. No:23, 34740 Kadıköy/İstanbul",
		IsBuyer:           true,
		IsSubMerchant:     true,
	}

	res, err := onboarding.CreateMember(request)
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func Test_UpdateSubMerchantMember(t *testing.T) {
	request := adapter.UpdateMemberRequest{
		ContactName:                   "Haluk",
		ContactSurname:                "Demir",
		Email:                         "haluk.demir@example.com",
		PhoneNumber:                   "905551111111",
		Iban:                          "TR930006701000000001111111",
		IdentityNumber:                "11111111110",
		LegalCompanyTitle:             "Dem Zeytinyağı Üretim Ltd. Şti.",
		Name:                          "Dem Zeytinyağı Üretim Ltd. Şti.",
		MemberType:                    model.MemberType(model.LIMITED_OR_JOINT_STOCK_COMPANY),
		SettlementEarningsDestination: model.SettlementEarningsDestination(model.SettlementEarningsDestinationIBAN),
		TaxNumber:                     "1111111114",
		TaxOffice:                     "Erenköy",
		Address:                       "Suadiye Mah. Örnek Cd. No:23, 34740 Kadıköy/İstanbul",
		IsBuyer:                       false,
		IsSubMerchant:                 true,
	}

	res, err := onboarding.UpdateMember(69271, request)
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func Test_UpdateBuyerMember(t *testing.T) {
	request := adapter.UpdateMemberRequest{
		ContactName:       "Haluk",
		ContactSurname:    "Demir",
		Email:             "haluk.demir@example.com",
		PhoneNumber:       "905551111111",
		Iban:              "TR930006701000000001111111",
		IdentityNumber:    "11111111110",
		LegalCompanyTitle: "Dem Zeytinyağı Üretim Ltd. Şti.",
		Name:              "Dem Zeytinyağı Üretim Ltd. Şti.",
		MemberType:        model.MemberType(model.PERSONAL),
		TaxNumber:         "1111111114",
		TaxOffice:         "Erenköy",
		Address:           "Suadiye Mah. Örnek Cd. No:23, 34740 Kadıköy/İstanbul",
		IsBuyer:           true,
		IsSubMerchant:     false,
	}

	res, err := onboarding.UpdateMember(69271, request)
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func Test_RetrieveMember(t *testing.T) {
	res, err := onboarding.RetrieveMember(69271)
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func Test_SearchMembers(t *testing.T) {
	request := adapter.SearchMembersRequest{
		Page: 0, Size: 10,
		IsBuyer:       true,
		IsSubMerchant: true,
	}

	res, err := onboarding.SearchMembers(request)
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}
