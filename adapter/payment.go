package adapter

import (
    "context"
    "crypto/sha256"
    "encoding/base64"
    "encoding/hex"
    "fmt"
    "log"
    "net/http"
    "strings"
)

type Payment struct {
    Client *Client
}

func (r *Init3DSPaymentResponse) GetDecodedHtmlContent() string {
    return Decode(*r.HtmlContent)
}

func (r *InitGarantiPayPaymentResponse) GetDecodedHtmlContent() string {
    return Decode(*r.HtmlContent)
}

func Decode(encoded string) string {
    data, err := base64.StdEncoding.DecodeString(encoded)
    if err != nil {
        log.Fatal("error:", err)
    }
    return string(data)
}

func (api *Payment) CreatePayment(ctx context.Context, request CreatePaymentRequest) (*PaymentResponse, error) {
    newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/payment/v1/card-payments", request)
    if err != nil {
        return nil, err
    }
    response := &Response[PaymentResponse]{}
    err = api.Client.Do(ctx, newRequest, response)
    if err != nil {
        return nil, err
    }

    return response.Data, nil
}

func (api *Payment) CreateApmPayment(ctx context.Context, request CreateApmPaymentRequest) (*PaymentResponse, error) {
    newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/payment/v1/apm-payments", request)
    if err != nil {
        return nil, err
    }
    response := &Response[PaymentResponse]{}
    err = api.Client.Do(ctx, newRequest, response)
    if err != nil {
        return nil, err
    }

    return response.Data, nil
}

func (api *Payment) RetrievePayment(ctx context.Context, paymentId int64) (*PaymentResponse, error) {
    newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("/payment/v1/card-payments/%d", paymentId), nil)
    if err != nil {
        return nil, err
    }
    response := &Response[PaymentResponse]{}
    err = api.Client.Do(ctx, newRequest, response)
    if err != nil {
        return nil, err
    }

    return response.Data, nil
}

func (api *Payment) Init3DSPayment(ctx context.Context, request Init3DSPaymentRequest) (*Init3DSPaymentResponse, error) {
    newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/payment/v1/card-payments/3ds-init", request)
    if err != nil {
        return nil, err
    }
    response := &Response[Init3DSPaymentResponse]{}
    err = api.Client.Do(ctx, newRequest, response)
    if err != nil {
        return nil, err
    }

    return response.Data, nil
}

func (api *Payment) Complete3DSPayment(ctx context.Context, request Complete3DSPaymentRequest) (*PaymentResponse, error) {
    newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/payment/v1/card-payments/3ds-complete", request)
    if err != nil {
        return nil, err
    }
    response := &Response[PaymentResponse]{}
    err = api.Client.Do(ctx, newRequest, response)
    if err != nil {
        return nil, err
    }

    return response.Data, nil
}

func (api *Payment) PostAuthPayment(ctx context.Context, paymentId int64, request PostAuthPaymentRequest) (*PaymentResponse, error) {
    newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, fmt.Sprintf("/payment/v1/card-payments/%d/post-auth", paymentId), request)
    if err != nil {
        return nil, err
    }

    response := &Response[PaymentResponse]{}
    err = api.Client.Do(ctx, newRequest, response)
    if err != nil {
        return nil, err
    }

    return response.Data, nil
}

func (api *Payment) InitCheckoutPayment(ctx context.Context, request InitCheckoutPaymentRequest) (*InitCheckoutPaymentResponse, error) {
    newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/payment/v1/checkout-payments/init", request)
    if err != nil {
        return nil, err
    }
    response := &Response[InitCheckoutPaymentResponse]{}
    err = api.Client.Do(ctx, newRequest, response)
    if err != nil {
        return nil, err
    }
    return response.Data, nil
}

func (api *Payment) RetrieveCheckoutPayment(ctx context.Context, token string) (*PaymentResponse, error) {
    newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("/payment/v1/checkout-payments/%s", token), nil)
    if err != nil {
        return nil, err
    }

    response := &Response[PaymentResponse]{}
    err = api.Client.Do(ctx, newRequest, response)
    if err != nil {
        return nil, err
    }

    return response.Data, nil
}

func (api *Payment) CreateDepositPayment(ctx context.Context, request DepositPaymentRequest) (*DepositPaymentResponse, error) {
    newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/payment/v1/deposits", request)

    if err != nil {
        return nil, err
    }

    response := &Response[DepositPaymentResponse]{}
    err = api.Client.Do(ctx, newRequest, response)

    if err != nil {
        return nil, err
    }

    return response.Data, nil
}

func (api *Payment) InitApmPayment(ctx context.Context, request InitApmPaymentRequest) (*InitApmPaymentResponse, error) {
    newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/payment/v1/apm-payments/init", request)

    if err != nil {
        return nil, err
    }

    response := &Response[InitApmPaymentResponse]{}
    err = api.Client.Do(ctx, newRequest, response)

    if err != nil {
        return nil, err
    }

    return response.Data, nil
}

func (api *Payment) CompleteApmPayment(ctx context.Context, request CompleteApmPaymentRequest) (*CompleteApmPaymentResponse, error) {
    newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/payment/v1/apm-payments/complete", request)

    if err != nil {
        return nil, err
    }

    response := &Response[CompleteApmPaymentResponse]{}
    err = api.Client.Do(ctx, newRequest, response)

    if err != nil {
        return nil, err
    }

    return response.Data, nil
}

func (api *Payment) Init3DSDepositPayment(ctx context.Context, request DepositPaymentRequest) (*Init3DSPaymentResponse, error) {
    newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/payment/v1/deposits/3ds-init", request)

    if err != nil {
        return nil, err
    }

    response := &Response[Init3DSPaymentResponse]{}
    err = api.Client.Do(ctx, newRequest, response)

    if err != nil {
        return nil, err
    }

    return response.Data, nil
}

func (api *Payment) Complete3DSDepositPayment(ctx context.Context, request Complete3DSPaymentRequest) (*DepositPaymentResponse, error) {
    newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/payment/v1/deposits/3ds-complete", request)
    if err != nil {
        return nil, err
    }

    response := &Response[DepositPaymentResponse]{}
    err = api.Client.Do(ctx, newRequest, response)

    if err != nil {
        return nil, err
    }

    return response.Data, nil
}

func (api *Payment) CreateFundTransferDepositPayment(ctx context.Context, request CreateFundTransferDepositPaymentRequest) (*FundTransferDepositPaymentResponse, error) {
    newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/payment/v1/deposits/fund-transfer", request)
    if err != nil {
        return nil, err
    }

    response := &Response[FundTransferDepositPaymentResponse]{}
    err = api.Client.Do(ctx, newRequest, response)

    if err != nil {
        return nil, err
    }

    return response.Data, nil
}

func (api *Payment) InitGarantiPayPayment(ctx context.Context, request InitGarantiPayPaymentRequest) (*InitGarantiPayPaymentResponse, error) {
    newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/payment/v1/garanti-pay-payments", request)
    if err != nil {
        return nil, err
    }

    response := &Response[InitGarantiPayPaymentResponse]{}
    err = api.Client.Do(ctx, newRequest, response)
    if err != nil {
        return nil, err
    }
    return response.Data, nil
}

func (api *Payment) RetrieveLoyalties(ctx context.Context, request RetrieveLoyaltiesRequest) (*RetrieveLoyaltiesResponse, error) {
    newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/payment/v1/card-loyalties/retrieve", request)
    if err != nil {
        return nil, err
    }

    response := &Response[RetrieveLoyaltiesResponse]{}
    err = api.Client.Do(ctx, newRequest, response)
    if err != nil {
        return nil, err
    }

    return response.Data, nil
}

func (api *Payment) RefundPaymentTransaction(ctx context.Context, request RefundPaymentTransactionRequest) (*PaymentTransactionRefundResponse, error) {
    newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/payment/v1/refund-transactions", request)
    if err != nil {
        return nil, err
    }

    response := &Response[PaymentTransactionRefundResponse]{}
    err = api.Client.Do(ctx, newRequest, response)

    if err != nil {
        return nil, err
    }

    return response.Data, nil
}

func (api *Payment) RetrievePaymentTransactionRefund(ctx context.Context, id int64) (*PaymentTransactionRefundResponse, error) {
    newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("/payment/v1/refund-transactions/%d", id), nil)
    if err != nil {
        return nil, err
    }

    response := &Response[PaymentTransactionRefundResponse]{}
    err = api.Client.Do(ctx, newRequest, response)

    if err != nil {
        return nil, err
    }

    return response.Data, nil
}

func (api *Payment) RefundPayment(ctx context.Context, request RefundPaymentRequest) (*PaymentRefundResponse, error) {
    newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/payment/v1/refunds", request)
    if err != nil {
        return nil, err
    }

    response := &Response[PaymentRefundResponse]{}
    err = api.Client.Do(ctx, newRequest, response)

    if err != nil {
        return nil, err
    }

    return response.Data, nil
}

func (api *Payment) RetrievePaymentRefund(ctx context.Context, id int64) (*PaymentRefundResponse, error) {
    newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("/payment/v1/refunds/%d", id), nil)
    if err != nil {
        return nil, err
    }

    response := &Response[PaymentRefundResponse]{}
    err = api.Client.Do(ctx, newRequest, response)
    if err != nil {
        return nil, err
    }

    return response.Data, nil
}

func (api *Payment) StoreCard(ctx context.Context, request StoreCardRequest) (*StoredCardResponse, error) {
    newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/payment/v1/cards", request)
    if err != nil {
        return nil, err
    }

    response := &Response[StoredCardResponse]{}
    err = api.Client.Do(ctx, newRequest, response)
    if err != nil {
        return nil, err
    }

    return response.Data, nil
}

func (api *Payment) UpdateStoredCard(ctx context.Context, request UpdateStoredCardRequest) (*StoredCardResponse, error) {
    newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/payment/v1/cards/update", request)

    if err != nil {
        return nil, err
    }

    response := &Response[StoredCardResponse]{}
    err = api.Client.Do(ctx, newRequest, response)
    if err != nil {
        return nil, err
    }

    return response.Data, nil
}

func (api *Payment) DeleteStoredCard(ctx context.Context, request DeleteStoredCardRequest) error {
    newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/payment/v1/cards", request)
    if err != nil {
        return err
    }

    response := &Void{}
    err = api.Client.Do(ctx, newRequest, response)
    if err != nil {
        return err
    }

    return nil
}

func (api *Payment) SearchStoredCards(ctx context.Context, request SearchStoredCardsRequest) (*DataResponse[StoredCardResponse], error) {
    newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, "/payment/v1/cards", request)
    if err != nil {
        return nil, err
    }
    response := &Response[DataResponse[StoredCardResponse]]{}
    err = api.Client.Do(ctx, newRequest, response)

    if err != nil {
        return nil, err
    }

    return response.Data, nil
}

func (api *Payment) UpdatePaymentTransaction(ctx context.Context, id int64, request UpdatePaymentTransactionRequest) (*PaymentTransactionResponse, error) {
    newRequest, err := api.Client.NewRequest(ctx, http.MethodPut, fmt.Sprintf("/payment/v1/payment-transactions/%d", id), request)
    if err != nil {
        return nil, err
    }

    response := &Response[PaymentTransactionResponse]{}
    err = api.Client.Do(ctx, newRequest, response)
    if err != nil {
        return nil, err
    }

    return response.Data, nil
}

func (api *Payment) ApprovePaymentTransactions(ctx context.Context, request PaymentTransactionsApprovalRequest) (*DataResponse[PaymentTransactionsApprovalResponse], error) {
    newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/payment/v1/payment-transactions/approve", request)
    if err != nil {
        return nil, err
    }

    response := &Response[DataResponse[PaymentTransactionsApprovalResponse]]{}
    err = api.Client.Do(ctx, newRequest, response)
    if err != nil {
        return nil, err
    }

    return response.Data, nil
}

func (api *Payment) DisapprovePaymentTransactions(ctx context.Context, request PaymentTransactionsApprovalRequest) (*DataResponse[PaymentTransactionsApprovalResponse], error) {
    newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/payment/v1/payment-transactions/disapprove", request)
    if err != nil {
        return nil, err
    }

    response := &Response[DataResponse[PaymentTransactionsApprovalResponse]]{}
    err = api.Client.Do(ctx, newRequest, response)
    if err != nil {
        return nil, err
    }

    return response.Data, nil
}

func (api *Payment) CheckMasterpassUser(ctx context.Context, request CheckMasterpassUserRequest) (*CheckMasterpassUserResponse, error) {
    newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "payment/v1/masterpass-payments/check-user", request)
    if err != nil {
        return nil, err
    }

    response := &Response[CheckMasterpassUserResponse]{}
    err = api.Client.Do(ctx, newRequest, response)
    if err != nil {
        return nil, err
    }

    return response.Data, nil
}

func (c *Payment) Is3DSecureCallbackVerified(threeDSecureCallbackKey string, params map[string]string) bool {
    hash := params["hash"]
    hashString := strings.Join([]string{threeDSecureCallbackKey,
        params["status"],
        params["completeStatus"],
        params["paymentId"],
        params["conversationData"],
        params["conversationId"],
        params["callbackStatus"],
    }, "###")

    hasher := sha256.New()
    hasher.Write([]byte(hashString))

    hashed := hex.EncodeToString(hasher.Sum(nil))
    return hash == hashed
}
