package tests

import (
	"context"
	"github.com/craftgate/craftgate-go-client/adapter"
	craftgate "github.com/craftgate/craftgate-go-client/adapter"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

var merchantClient, _ = craftgate.New("api-key", "secret-key", "http://localhost:8000")

func Test_CreateMerchantPos(t *testing.T) {
	request := adapter.CreateMerchantPosRequest{
		Name:                    "my test pos",
		ClientId:                "client id",
		TerminalId:              "terminal id",
		ThreedsKey:              "3d secure key",
		Status:                  craftgate.PosStatus_AUTOPILOT,
		Currency:                craftgate.Currency_TRY,
		OrderNumber:             1,
		EnableInstallment:       true,
		EnableForeignCard:       true,
		EnablePaymentWithoutCvc: true,
		PosIntegrator:           craftgate.PosIntegrator_AKBANK,
		EnabledPaymentAuthenticationTypes: []craftgate.PaymentAuthenticationType{
			craftgate.PaymentAuthenticationType_NON_THREE_DS,
			craftgate.PaymentAuthenticationType_NON_THREE_DS,
		},
		MerchantPosUsers: []craftgate.CreateMerchantPosUser{
			{
				PosOperationType: craftgate.PosOperationType_STANDARD,
				PosUserType:      craftgate.PosUserType_API,
				PosUsername:      "username",
				PosPassword:      "password",
			},
		},
	}

	res, err := merchantClient.Merchant.CreateMerchantPos(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func Test_RetrieveMerchantPos(t *testing.T) {
	res, err := merchantClient.Merchant.RetrieveMerchantPos(context.Background(), 1)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func Test_UpdateMerchantPosStatus(t *testing.T) {
	err := merchantClient.Merchant.UpdateMerchantPosStatus(context.Background(), 1, craftgate.PosStatus_ACTIVE)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func Test_DeleteMerchantPosStatus(t *testing.T) {
	err := merchantClient.Merchant.DeleteMerchantPosStatus(context.Background(), 1)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func Test_RetrieveMerchantPosCommission(t *testing.T) {
	res, err := merchantClient.Merchant.RetrieveMerchantPosCommissions(context.Background(), 14)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func Test_UpdateMerchantPosCommissions(t *testing.T) {
	req := adapter.CreateMerchantPosCommissionRequest{
		Commissions: []craftgate.CreateMerchantPosCommission{
			{
				Installment:                         1,
				BlockageDay:                         7,
				Status:                              craftgate.Status_ACTIVE,
				CardBrandName:                       craftgate.CardBrand_AXESS,
				InstallmentLabel:                    "Single installment",
				BankOnUsDebitCardCommissionRate:     1.0,
				BankOnUsCreditCardCommissionRate:    1.1,
				BankNotOnUsDebitCardCommissionRate:  1.2,
				BankNotOnUsCreditCardCommissionRate: 1.3,
				BankForeignCardCommissionRate:       1.5,
			},
			{
				Installment:                      2,
				BlockageDay:                      7,
				Status:                           craftgate.Status_ACTIVE,
				CardBrandName:                    craftgate.CardBrand_AXESS,
				InstallmentLabel:                 "installment 2",
				BankOnUsCreditCardCommissionRate: 2.1,
				MerchantCommissionRate:           2.3,
			},
		},
	}
	res, err := merchantClient.Merchant.UpdateMerchantPosCommissions(context.Background(), 14, req)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func Test_SearchMerchantPos(t *testing.T) {
	req := adapter.SearchMerchantPosRequest{
		Page:     0,
		Size:     10,
		Currency: adapter.Currency_TRY,
	}
	res, err := merchantClient.Merchant.SearchMerchantPos(context.Background(), req)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}
