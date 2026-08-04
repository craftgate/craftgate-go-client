package tests

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/craftgate/craftgate-go-client/adapter"
)

const idempotencyKeyHeaderName = "x-idempotency-key"

var idempotencyClient, _ = adapter.New("api-key", "secret-key", "https://sandbox-api.craftgate.io")

func readBody(t *testing.T, req *http.Request) string {
	t.Helper()
	if req.Body == nil {
		return ""
	}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		t.Fatalf("could not read request body: %s", err)
	}
	return string(body)
}

func Test_Idempotency_BodyRequest_SendsKeyAsHeader(t *testing.T) {
	request := adapter.CreatePaymentRequest{Price: 100}
	request.HeaderOptions = adapter.HeaderOptions{IdempotencyKey: "idempotency-key-1"}

	req, err := idempotencyClient.NewRequest(context.Background(), http.MethodPost, "/payment/v1/card-payments", request)
	if err != nil {
		t.Fatalf("Error %s", err)
	}

	if got := req.Header.Get(idempotencyKeyHeaderName); got != "idempotency-key-1" {
		t.Errorf("expected header idempotency-key-1, got %q", got)
	}
}

func Test_Idempotency_BodyRequest_WithoutKeySendsNoHeader(t *testing.T) {
	request := adapter.CreatePaymentRequest{Price: 100}

	req, err := idempotencyClient.NewRequest(context.Background(), http.MethodPost, "/payment/v1/card-payments", request)
	if err != nil {
		t.Fatalf("Error %s", err)
	}

	if _, present := req.Header[http.CanonicalHeaderKey(idempotencyKeyHeaderName)]; present {
		t.Error("expected no idempotency key header")
	}
}

func Test_Idempotency_KeyIsExcludedFromBody(t *testing.T) {
	request := adapter.CreatePaymentRequest{Price: 100}
	request.HeaderOptions = adapter.HeaderOptions{IdempotencyKey: "idempotency-key-1"}

	req, err := idempotencyClient.NewRequest(context.Background(), http.MethodPost, "/payment/v1/card-payments", request)
	if err != nil {
		t.Fatalf("Error %s", err)
	}

	body := readBody(t, req)
	if !strings.Contains(body, `"price":100`) {
		t.Errorf("expected price in body, got %q", body)
	}
	if strings.Contains(body, "HeaderOptions") || strings.Contains(body, "headerOptions") ||
		strings.Contains(body, "IdempotencyKey") || strings.Contains(body, "idempotencyKey") ||
		strings.Contains(body, "idempotency-key-1") {
		t.Errorf("idempotency key leaked into body: %q", body)
	}
}

func Test_Idempotency_BodySignatureIsUnchangedByKey(t *testing.T) {
	withKey := adapter.CreatePaymentRequest{Price: 100}
	withKey.HeaderOptions = adapter.HeaderOptions{IdempotencyKey: "idempotency-key-1"}
	withoutKey := adapter.CreatePaymentRequest{Price: 100}

	a, err := idempotencyClient.NewRequest(context.Background(), http.MethodPost, "/payment/v1/card-payments", withKey)
	if err != nil {
		t.Fatalf("Error %s", err)
	}
	b, err := idempotencyClient.NewRequest(context.Background(), http.MethodPost, "/payment/v1/card-payments", withoutKey)
	if err != nil {
		t.Fatalf("Error %s", err)
	}

	expected := adapter.GenerateHash(b.URL.String(), "api-key", "secret-key",
		a.Header.Get(adapter.RandomHeaderName), readBody(t, b))
	if got := a.Header.Get(adapter.SignatureHeaderName); got != expected {
		t.Errorf("signature changed by idempotency key:\n got      %s\n expected %s", got, expected)
	}
}

func Test_Idempotency_PathOnlyRequest_SendsKeyAndNoQueryOrBody(t *testing.T) {
	request := adapter.ExpireCheckoutPaymentRequest{Token: "token-1"}
	request.HeaderOptions = adapter.HeaderOptions{IdempotencyKey: "idempotency-key-1"}

	req, err := idempotencyClient.NewRequestWithoutBody(context.Background(), http.MethodDelete,
		"/payment/v1/checkout-payments/token-1", request)
	if err != nil {
		t.Fatalf("Error %s", err)
	}

	if got := req.Header.Get(idempotencyKeyHeaderName); got != "idempotency-key-1" {
		t.Errorf("expected header idempotency-key-1, got %q", got)
	}
	if req.URL.RawQuery != "" {
		t.Errorf("expected no query string, got %q", req.URL.RawQuery)
	}
	if body := readBody(t, req); body != "" {
		t.Errorf("expected no body, got %q", body)
	}
}

func Test_Idempotency_PathOnlySignatureIsUnchangedByKey(t *testing.T) {
	path := "/payment/v1/checkout-payments/token-1"

	withKey := adapter.ExpireCheckoutPaymentRequest{Token: "token-1"}
	withKey.HeaderOptions = adapter.HeaderOptions{IdempotencyKey: "idempotency-key-1"}

	a, err := idempotencyClient.NewRequestWithoutBody(context.Background(), http.MethodDelete, path,
		withKey)
	if err != nil {
		t.Fatalf("Error %s", err)
	}
	b, err := idempotencyClient.NewRequestWithoutBody(context.Background(), http.MethodDelete, path,
		adapter.ExpireCheckoutPaymentRequest{Token: "token-1"})
	if err != nil {
		t.Fatalf("Error %s", err)
	}

	expected := adapter.GenerateHash(b.URL.String(), "api-key", "secret-key",
		a.Header.Get(adapter.RandomHeaderName), "")
	if got := a.Header.Get(adapter.SignatureHeaderName); got != expected {
		t.Errorf("signature changed by idempotency key:\n got      %s\n expected %s", got, expected)
	}
	if _, present := b.Header[http.CanonicalHeaderKey(idempotencyKeyHeaderName)]; present {
		t.Error("expected no idempotency key header when key is empty")
	}
}

func Test_Idempotency_KeyIsExcludedFromQueryParams(t *testing.T) {
	request := adapter.SearchPaymentsRequest{Page: 0, Size: 10}
	request.HeaderOptions = adapter.HeaderOptions{IdempotencyKey: "idempotency-key-1"}

	query, err := adapter.QueryParams(request)
	if err != nil {
		t.Fatalf("Error %s", err)
	}

	if strings.Contains(query, "HeaderOptions") || strings.Contains(query, "headerOptions") ||
		strings.Contains(query, "IdempotencyKey") || strings.Contains(query, "idempotencyKey") ||
		strings.Contains(query, "idempotency-key-1") {
		t.Errorf("idempotency key leaked into query params: %q", query)
	}

	values, err := url.ParseQuery(query)
	if err != nil {
		t.Fatalf("Error %s", err)
	}
	if values.Get("size") != "10" {
		t.Errorf("expected size=10 in query, got %q", query)
	}
}

func Test_Idempotency_HeaderOptionsAreReadFromValueAndPointerRequests(t *testing.T) {
	path := "/craftlink/v1/products/42"
	request := adapter.DeleteProductRequest{Id: 42}
	request.HeaderOptions = adapter.HeaderOptions{IdempotencyKey: "idempotency-key-1"}

	value, err := idempotencyClient.NewRequestWithoutBody(context.Background(), http.MethodDelete, path, request)
	if err != nil {
		t.Fatalf("Error %s", err)
	}
	if got := value.Header.Get(idempotencyKeyHeaderName); got != "idempotency-key-1" {
		t.Errorf("expected idempotency-key-1, got %q", got)
	}

	pointer, err := idempotencyClient.NewRequestWithoutBody(context.Background(), http.MethodDelete, path, &request)
	if err != nil {
		t.Fatalf("Error %s", err)
	}
	if got := pointer.Header.Get(idempotencyKeyHeaderName); got != "idempotency-key-1" {
		t.Errorf("expected idempotency-key-1 for pointer, got %q", got)
	}
}

func Test_Idempotency_NilBodySendsNoHeader(t *testing.T) {
	req, err := idempotencyClient.NewRequest(context.Background(), http.MethodGet, "/craftlink/v1/products/42", nil)
	if err != nil {
		t.Fatalf("Error %s", err)
	}

	if _, present := req.Header[http.CanonicalHeaderKey(idempotencyKeyHeaderName)]; present {
		t.Error("expected no idempotency key header for a nil body")
	}
}

func Test_Idempotency_WrappersCarryPathVariables(t *testing.T) {
	removeValue := adapter.RemoveValueFromValueListRequest{ListName: "ipList", ValueId: "value-1"}
	removeValue.HeaderOptions = adapter.HeaderOptions{IdempotencyKey: "idempotency-key-1"}

	posStatus := adapter.UpdateMerchantPosStatusRequest{MerchantPosId: 1, PosStatus: adapter.PosStatus_PASSIVE}
	posStatus.HeaderOptions = adapter.HeaderOptions{IdempotencyKey: "idempotency-key-2"}

	if removeValue.ListName != "ipList" || removeValue.ValueId != "value-1" ||
		removeValue.HeaderOptions.IdempotencyKey != "idempotency-key-1" {
		t.Errorf("unexpected wrapper contents: %+v", removeValue)
	}
	if posStatus.MerchantPosId != 1 || posStatus.PosStatus != adapter.PosStatus_PASSIVE ||
		posStatus.HeaderOptions.IdempotencyKey != "idempotency-key-2" {
		t.Errorf("unexpected wrapper contents: %+v", posStatus)
	}
}
