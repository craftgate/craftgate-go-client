package adapter

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

type Hook struct {
	Client *Client
}

func (c *Hook) IsWebhookVerified(merchantHookKey, incomingSignature string, webhookData WebhookData) bool {
	data := fmt.Sprintf("%s%d%s%s", webhookData.EventType, webhookData.EventTimestamp, webhookData.Status, webhookData.PayloadId)

	hmac := hmac.New(sha256.New, []byte(merchantHookKey))
	hmac.Write([]byte(data))
	signature := base64.StdEncoding.EncodeToString(hmac.Sum(nil))

	return incomingSignature == signature
}
