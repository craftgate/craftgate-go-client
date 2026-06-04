package tests

import (
	"context"
	"testing"

	"github.com/craftgate/craftgate-go-client/adapter"
	craftgate "github.com/craftgate/craftgate-go-client/adapter"
	"github.com/davecgh/go-spew/spew"
    "github.com/stretchr/testify/require"
)

var mealVoucherClient, _ = craftgate.New("api-key", "secret-key", "https://sandbox-api.craftgate.io")

func TestMealVoucherCardTokenization_Init(t *testing.T) {
	request := adapter.MealVoucherCardTokenizationInitRequest{
		ApmType: craftgate.ApmType_SETCARD,
		MealVoucherCardTokenizationData: adapter.MealVoucherCardTokenizationData{
			CallbackUrl: "https://webhook.site/e806070a-da76-4d02-a67b-54ba9e8332d3",
		},
	}

	res, err := mealVoucherClient.MealVoucherCardTokenization.Init(context.Background(), request)

    _, _ = spew.Printf("%#v\n", res)
    require.Nil(t, err, "Error: %v", err)
    require.NotNil(t, res.HtmlContent)
}

func TestMealVoucherCardTokenization_Regenerate(t *testing.T) {
	request := adapter.MealVoucherCardTokenizationRegenerateRequest{
		MealVoucherCardTokenizationData: adapter.MealVoucherCardTokenizationData{
			CardNumber: "123456",
		},
	}

	res, err := mealVoucherClient.MealVoucherCardTokenization.Regenerate(context.Background(), "session-id", request)

    _, _ = spew.Printf("%#v\n", res)
    require.Nil(t, err, "Error: %v", err)
    require.NotNil(t, res.HtmlContent)
}

func TestMealVoucherCardTokenization_Complete(t *testing.T) {
	request := adapter.MealVoucherCardTokenizationCompleteRequest{
		ValidationCode: "123456",
	}

	res, err := mealVoucherClient.MealVoucherCardTokenization.Complete(context.Background(), "session-id", request)

    _, _ = spew.Printf("%#v\n", res)
    require.Nil(t, err, "Error: %v", err)
    require.NotNil(t, res.MaskedCardNumber)
}
