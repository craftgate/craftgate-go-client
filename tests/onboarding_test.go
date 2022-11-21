package tests

import (
    "context"
    "fmt"
    "github.com/craftgate/craftgate-go-client/v1/adapter"
    craftgate "github.com/craftgate/craftgate-go-client/v1/adapter"
    "github.com/davecgh/go-spew/spew"
    "testing"
    "time"
)

var onboardingClient, _ = craftgate.New("api-key", "secret-key", "https://sandbox-api.craftgate.io")

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
        MemberType:        craftgate.MemberType(craftgate.LIMITED_OR_JOINT_STOCK_COMPANY),
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
        MemberType:        craftgate.MemberType(craftgate.LIMITED_OR_JOINT_STOCK_COMPANY),
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
        MemberType:                    craftgate.MemberType(craftgate.LIMITED_OR_JOINT_STOCK_COMPANY),
        SettlementEarningsDestination: craftgate.SettlementEarningsDestination(craftgate.SettlementEarningsDestinationIBAN),
        TaxNumber:                     "1111111114",
        TaxOffice:                     "Erenköy",
        Address:                       "Suadiye Mah. Örnek Cd. No:23, 34740 Kadıköy/İstanbul",
        IsBuyer:                       false,
        IsSubMerchant:                 true,
    }

    res, err := onboardingClient.Onboarding.UpdateMember(context.Background(), 86821, request)
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
        MemberType:        craftgate.MemberType(craftgate.PERSONAL),
        TaxNumber:         "1111111114",
        TaxOffice:         "Erenköy",
        Address:           "Suadiye Mah. Örnek Cd. No:23, 34740 Kadıköy/İstanbul",
        IsBuyer:           true,
        IsSubMerchant:     false,
    }

    res, err := onboardingClient.Onboarding.UpdateMember(context.Background(), 86819, request)
    _, _ = spew.Printf("%#v\n", res)

    if err != nil {
        t.Errorf("Error %s", err)
    }
}

func Test_RetrieveMember(t *testing.T) {
    res, err := onboardingClient.Onboarding.RetrieveMember(context.Background(), 86821)
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
