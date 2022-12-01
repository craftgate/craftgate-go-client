package adapter

import (
	"context"
	"fmt"
	"net/http"
)

type PaymentReporting struct {
	Client *Client
}

func (api *PaymentReporting) SearchPayments(ctx context.Context, request SearchPaymentsRequest) (*DataResponse[ReportingPaymentResponse], error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, "/payment-reporting/v1/payments", request)
	if err != nil {
		return nil, err
	}

	response := &Response[DataResponse[ReportingPaymentResponse]]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *PaymentReporting) SearchPaymentRefunds(ctx context.Context, request SearchPaymentRefundsRequest) (*DataResponse[ReportingPaymentRefundResponse], error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, "/payment-reporting/v1/refunds", request)
	if err != nil {
		return nil, err
	}

	response := &Response[DataResponse[ReportingPaymentRefundResponse]]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *PaymentReporting) SearchPaymentTransactionRefunds(ctx context.Context, request SearchPaymentTransactionRefundsRequest) (*DataResponse[ReportingPaymentTransactionRefundResponse], error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, "/payment-reporting/v1/refund-transactions", request)
	if err != nil {
		return nil, err
	}

	response := &Response[DataResponse[ReportingPaymentTransactionRefundResponse]]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *PaymentReporting) RetrievePayment(ctx context.Context, id int64) (*ReportingPaymentResponse, error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("/payment-reporting/v1/payments/%d", id), nil)
	if err != nil {
		return nil, err
	}

	response := &Response[ReportingPaymentResponse]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *PaymentReporting) RetrievePaymentTransactions(ctx context.Context, paymentId int64) (*DataResponse[ReportingPaymentTransactionResponse], error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("/payment-reporting/v1/payments/%d/transactions", paymentId), nil)
	if err != nil {
		return nil, err
	}

	response := &Response[DataResponse[ReportingPaymentTransactionResponse]]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *PaymentReporting) RetrievePaymentRefunds(ctx context.Context, paymentId int64) (*DataResponse[ReportingPaymentRefundResponse], error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("/payment-reporting/v1/payments/%d/refunds", paymentId), nil)
	if err != nil {
		return nil, err
	}

	response := &Response[DataResponse[ReportingPaymentRefundResponse]]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (api *PaymentReporting) RetrievePaymentTransactionRefunds(ctx context.Context, paymentId, paymentTransactionId int64) (*DataResponse[ReportingPaymentTransactionRefundResponse], error) {
	newRequest, err := api.Client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("/payment-reporting/v1/payments/%d/transactions/%d/refunds", paymentId, paymentTransactionId), nil)
	if err != nil {
		return nil, err
	}

	response := &Response[DataResponse[ReportingPaymentTransactionRefundResponse]]{}
	err = api.Client.Do(ctx, newRequest, response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
