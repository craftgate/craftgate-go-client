package tests

import (
	craftgate "github.com/craftgate/craftgate-go-client/adapter"
	"github.com/stretchr/testify/require"
	"testing"
)

var cg, _ = craftgate.New("api-key", "secret-key", "https://sandbox-api.craftgate.io")

func TestHook_VerifyWebhook(t *testing.T) {
	merchantHookKey := "Aoh7tReTybO6wOjBmOJFFsOR53SBojEp"
	incomingSignature := "0wRB5XqWJxwwPbn5Z9TcbHh8EGYFufSYTsRMB74N094="
	webhookData := craftgate.WebhookData{
		EventType:      craftgate.WebhookEventType(craftgate.API_VERIFY_AND_AUTH),
		EventTimestamp: 1661521221,
		Status:         craftgate.WebhookStatus(craftgate.WebhookStatusSUCCESS),
		PayloadId:      "584",
	}

	isVerified := cg.Hook.IsWebhookVerified(merchantHookKey, incomingSignature, webhookData)

	require.True(t, isVerified)
}

func TestHook_NotVerifyWebhook(t *testing.T) {
	merchantHookKey := "Aoh7tReTybO6wOjBmOJFFsOR53SBojEp"
	incomingSignature := "Bsa498wcnaasd4bhx8anxıxcsdnxanalkdjcahxhd"
	webhookData := craftgate.WebhookData{
		EventType:      craftgate.WebhookEventType(craftgate.API_VERIFY_AND_AUTH),
		EventTimestamp: 1661521221,
		Status:         craftgate.WebhookStatus(craftgate.WebhookStatusSUCCESS),
		PayloadId:      "584",
	}

	isVerified := cg.Hook.IsWebhookVerified(merchantHookKey, incomingSignature, webhookData)

	require.False(t, isVerified)
}
