package tests

import (
	"context"
	"fmt"
	"github.com/craftgate/craftgate-go-client/adapter"
	craftgate "github.com/craftgate/craftgate-go-client/adapter"
	"github.com/davecgh/go-spew/spew"
	"testing"
	"time"
)

var onboardingClient, _ = craftgate.New("api-key", "secret-key", "https://sandbox-api.craftgate.io")

func Test_CreateBuyerMember(t *testing.T) {
	request := adapter.CreateMemberRequest{
		MemberExternalId: fmt.Sprintf("%d", time.Now().Nanosecond()),
		Name:             "Haluk Demir",
		Email:            "haluk.demir@example.com",
		PhoneNumber:      "905551111111",
		Address:          "Suadiye Mah. Örnek Cd. No:23, 34740 Kadıköy/İstanbul",
		ContactName:      "Haluk",
		ContactSurname:   "Demir",
		IsBuyer:          true,
		IsSubMerchant:    false,
	}

	res, err := onboardingClient.Onboarding.CreateMember(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

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
		MemberType:        craftgate.MemberType_LIMITED_OR_JOINT_STOCK_COMPANY,
		TaxNumber:         "1111111114",
		TaxOffice:         "Erenköy",
		Address:           "Suadiye Mah. Örnek Cd. No:23, 34740 Kadıköy/İstanbul",
		IsBuyer:           false,
		IsSubMerchant:     true,
	}

	res, err := onboardingClient.Onboarding.CreateMember(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

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
		MemberType:        craftgate.MemberType_LIMITED_OR_JOINT_STOCK_COMPANY,
		TaxNumber:         "1111111114",
		TaxOffice:         "Erenköy",
		Address:           "Suadiye Mah. Örnek Cd. No:23, 34740 Kadıköy/İstanbul",
		IsBuyer:           true,
		IsSubMerchant:     true,
	}

	res, err := onboardingClient.Onboarding.CreateMember(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

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
		MemberType:                    craftgate.MemberType_LIMITED_OR_JOINT_STOCK_COMPANY,
		SettlementEarningsDestination: craftgate.SettlementEarningsDestination_WALLET,
		TaxNumber:                     "1111111114",
		TaxOffice:                     "Erenköy",
		Address:                       "Suadiye Mah. Örnek Cd. No:23, 34740 Kadıköy/İstanbul",
		IsBuyer:                       false,
		IsSubMerchant:                 true,
	}

	res, err := onboardingClient.Onboarding.UpdateMember(context.Background(), 1, request)
	_, _ = spew.Printf("%#v\n", res)

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
		MemberType:        craftgate.MemberType_PERSONAL,
		TaxNumber:         "1111111114",
		TaxOffice:         "Erenköy",
		Address:           "Suadiye Mah. Örnek Cd. No:23, 34740 Kadıköy/İstanbul",
		IsBuyer:           true,
		IsSubMerchant:     false,
	}

	res, err := onboardingClient.Onboarding.UpdateMember(context.Background(), 1, request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func Test_RetrieveMember(t *testing.T) {
	res, err := onboardingClient.Onboarding.RetrieveMember(context.Background(), 1)
	_, _ = spew.Printf("%#v\n", res)

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

	res, err := onboardingClient.Onboarding.SearchMembers(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func Test_CreateMerchant(t *testing.T) {
	request := adapter.CreateMerchantRequest{
		Name:               "newMerchant",
		LegalCompanyTitle:  "legalCompanyTitle",
		Email:              "new_merchant@merchant.com",
		Website:            "www.merchant.com",
		ContactName:        "newName",
		ContactSurname:     "newSurname",
		ContactPhoneNumber: "905555555566",
	}

	res, err := onboardingClient.Onboarding.CreateMerchant(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}
