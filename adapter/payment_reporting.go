package adapter

import (
	"craftgate-go-client/adapter/rest"
	"craftgate-go-client/model"
	"fmt"
	"net/http"
	"time"
)

type PaymentReporting struct {
	Opts model.RequestOptions
}

type SearchPaymentsRequest struct {
	Page                 int
	Size                 int
	PaymentId            int64
	PaymentTransactionId int64
	BuyerMemberId        int64
	SubMerchantMemberId  int64
	ConversationId       string
	ExternalId           string
	OrderId              string
	PaymentType          model.PaymentType
	PaymentProvider      model.PaymentProvider
	PaymentStatus        model.PaymentStatus
	PaymentSource        model.PaymentSource
	PaymentChannel       string
	BinNumber            string
	LastFourDigits       string
	Currency             model.Currency
	MinPaidPrice         float64
	MaxPaidPrice         float64
	Installment          int
	IsThreeDS            bool
	MinCreatedDate       time.Time
	MaxCreatedDate       time.Time
}

type ReportingPaymentResponse struct {
	Id int64 `json:"id"`
	//CreatedDate time.Time      `json:"createdDate"`
	OrderId     string         `json:"orderId"`
	Price       float64        `json:"price"`
	PaidPrice   float64        `json:"paidPrice"`
	WalletPrice float64        `json:"walletPrice"`
	Currency    model.Currency `json:"currency"`
}

func (api *PaymentReporting) SearchPayments(request SearchPaymentsRequest) (*model.Response[ReportingPaymentResponse], error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/payment-reporting/v1/payments?", api.Opts.BaseURL), nil)

	q := req.URL.Query()
	//q.Add("currency", string(request.Currency))

	//t := reflect.TypeOf(request)
	//for i := 0; i < t.NumField(); i++ {
	//	field := t.Field(i)
	//	value := reflect.ValueOf(field)
	//	fmt.Println(field, value)
	//
	//	//q.Add(field.Name, string(field.))
	//}

	req.URL.RawQuery = q.Encode()

	res := model.Response[ReportingPaymentResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}
