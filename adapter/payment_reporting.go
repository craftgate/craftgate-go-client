package adapter

import (
	"craftgate-go-client/adapter/rest"
	"craftgate-go-client/model"
	"fmt"
	"net/http"
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
	MinCreatedDate       Time
	MaxCreatedDate       Time
}

type ReportingPaymentResponse struct {
	Id                           int64                            `json:"id"`
	CreatedDate                  Time                             `json:"createdDate"`
	Price                        float64                          `json:"price"`
	PaidPrice                    float64                          `json:"paidPrice"`
	WalletPrice                  float64                          `json:"walletPrice"`
	Currency                     model.Currency                   `json:"currency"`
	BuyerMemberId                int64                            `json:"buyerMemberId"`
	Installment                  int32                            `json:"installment"`
	ConversationId               string                           `json:"conversationId"`
	ExternalId                   string                           `json:"externalId"`
	PaymentType                  model.PaymentType                `json:"paymentType"`
	PaymentProvider              model.PaymentProvider            `json:"paymentProvider"`
	PaymentSource                model.PaymentSource              `json:"paymentSource"`
	PaymentGroup                 model.PaymentGroup               `json:"paymentGroup"`
	PaymentStatus                model.PaymentStatus              `json:"paymentStatus"`
	PaymentPhase                 model.PaymentPhase               `json:"paymentPhase"`
	PaymentChannel               string                           `json:"paymentChannel"`
	IsThreeDS                    bool                             `json:"isThreeDS"`
	MerchantCommissionRate       float64                          `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount float64                          `json:"merchantCommissionRateAmount"`
	BankCommissionRate           float64                          `json:"bankCommissionRate"`
	BankCommissionRateAmount     float64                          `json:"bankCommissionRateAmount"`
	PaidWithStoredCard           bool                             `json:"paidWithStoredCard"`
	BinNumber                    string                           `json:"binNumber"`
	LastFourDigits               string                           `json:"lastFourDigits"`
	AuthCode                     string                           `json:"authCode"`
	HostReference                string                           `json:"hostReference"`
	OrderId                      string                           `json:"orderId"`
	TransId                      string                           `json:"transId"`
	CardHolderName               string                           `json:"cardHolderName"`
	BankCardHolderName           string                           `json:"bankCardHolderName"`
	CardType                     model.CardType                   `json:"cardType"`
	CardAssociation              model.CardAssociation            `json:"cardAssociation"`
	CardBrand                    string                           `json:"cardBrand"`
	RequestedPosAlias            string                           `json:"requestedPosAlias"`
	RetryCount                   int                              `json:"retryCount"`
	RefundablePrice              float64                          `json:"refundablePrice"`
	RefundStatus                 model.PaymentRefundStatus        `json:"refundStatus"`
	CardIssuerBankName           string                           `json:"cardIssuerBankName"`
	MdStatus                     int                              `json:"mdStatus"`
	BuyerMember                  model.MemberResponse             `json:"buyerMember"`
	Refunds                      []ReportingPaymentRefundResponse `json:"refunds"`
	Pos                          model.MerchantPos                `json:"pos"`
	Loyalty                      model.Loyalty                    `json:"loyalty"`
	PaymentError                 model.PaymentError               `json:"paymentError"`
}

type ReportingPaymentRefundResponse struct {
	Id                    int64                       `json:"id"`
	CreatedDate           Time                        `json:"createdDate"`
	Status                model.RefundStatus          `json:"status"`
	RefundDestinationType model.RefundDestinationType `json:"refundDestinationType"`
	RefundPrice           float64                     `json:"refundPrice"`
	RefundBankPrice       float64                     `json:"refundBankPrice"`
	RefundWalletPrice     float64                     `json:"refundWalletPrice"`
	ConversationId        string                      `json:"conversationId"`
	AuthCode              string                      `json:"authCode"`
	HostReference         string                      `json:"hostReference"`
	PaymentType           model.PaymentType           `json:"paymentType"`
	TransId               string                      `json:"transId"`
	PaymentId             int64                       `json:"paymentId"`
	PaymentError          model.PaymentError          `json:"paymentError"`
}

func (api *PaymentReporting) SearchPayments(request SearchPaymentsRequest) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/payment-reporting/v1/payments?", api.Opts.BaseURL), nil)

	//req.URL.RawQuery, _ = QueryParams(request)

	res := model.Response[model.DataResponse[ReportingPaymentResponse]]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *PaymentReporting) RetrievePayment(id int64) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/payment-reporting/v1/payments/%d", api.Opts.BaseURL, id), nil)

	res := model.Response[ReportingPaymentResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}
