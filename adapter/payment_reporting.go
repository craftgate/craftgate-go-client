package adapter

import (
	"craftgate-go-client/adapter/rest"
	"fmt"
	"net/http"
)

type PaymentReporting struct {
	Opts RequestOptions
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
	PaymentType          PaymentType
	PaymentProvider      PaymentProvider
	PaymentStatus        PaymentStatus
	PaymentSource        PaymentSource
	PaymentChannel       string
	BinNumber            string
	LastFourDigits       string
	Currency             Currency
	MinPaidPrice         float64
	MaxPaidPrice         float64
	Installment          int
	IsThreeDS            bool
	MinCreatedDate       Time
	MaxCreatedDate       Time
}

type SearchPaymentRefundsRequest struct {
	Page           int
	Size           int
	Id             int64
	PaymentId      int64
	BuyerMemberId  int64
	ConversationId string
	Status         RefundStatus
	Currency       Currency
	MinRefundPrice float64
	MaxRefundPrice float64
	MinCreatedDate Time
	MaxCreatedDate Time
}

type SearchPaymentTransactionRefundsRequest struct {
	Page                 int
	Size                 int
	Id                   int64
	PaymentId            int64
	PaymentTransactionId int64
	BuyerMemberId        int64
	ConversationId       string
	Status               RefundStatus
	Currency             Currency
	IsAfterSettlement    bool
	MinRefundPrice       float64
	MaxRefundPrice       float64
	MinCreatedDate       Time
	MaxCreatedDate       Time
}

type ReportingPaymentResponse struct {
	Id                           int64                            `json:"id"`
	CreatedDate                  Time                             `json:"createdDate"`
	Price                        float64                          `json:"price"`
	PaidPrice                    float64                          `json:"paidPrice"`
	WalletPrice                  float64                          `json:"walletPrice"`
	Currency                     Currency                         `json:"currency"`
	BuyerMemberId                int64                            `json:"buyerMemberId"`
	Installment                  int32                            `json:"installment"`
	ConversationId               string                           `json:"conversationId"`
	ExternalId                   string                           `json:"externalId"`
	PaymentType                  PaymentType                      `json:"paymentType"`
	PaymentProvider              PaymentProvider                  `json:"paymentProvider"`
	PaymentSource                PaymentSource                    `json:"paymentSource"`
	PaymentGroup                 PaymentGroup                     `json:"paymentGroup"`
	PaymentStatus                PaymentStatus                    `json:"paymentStatus"`
	PaymentPhase                 PaymentPhase                     `json:"paymentPhase"`
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
	CardType                     CardType                         `json:"cardType"`
	CardAssociation              CardAssociation                  `json:"cardAssociation"`
	CardBrand                    string                           `json:"cardBrand"`
	RequestedPosAlias            string                           `json:"requestedPosAlias"`
	RetryCount                   int                              `json:"retryCount"`
	RefundablePrice              float64                          `json:"refundablePrice"`
	RefundStatus                 PaymentRefundStatus              `json:"refundStatus"`
	CardIssuerBankName           string                           `json:"cardIssuerBankName"`
	MdStatus                     int                              `json:"mdStatus"`
	BuyerMember                  MemberResponse                   `json:"buyerMember"`
	Refunds                      []ReportingPaymentRefundResponse `json:"refunds"`
	Pos                          MerchantPos                      `json:"pos"`
	Loyalty                      Loyalty                          `json:"loyalty"`
	PaymentError                 PaymentError                     `json:"paymentError"`
}

type ReportingPaymentRefundResponse struct {
	Id                    int64                 `json:"id"`
	CreatedDate           Time                  `json:"createdDate"`
	Status                RefundStatus          `json:"status"`
	RefundDestinationType RefundDestinationType `json:"refundDestinationType"`
	RefundPrice           float64               `json:"refundPrice"`
	RefundBankPrice       float64               `json:"refundBankPrice"`
	RefundWalletPrice     float64               `json:"refundWalletPrice"`
	ConversationId        string                `json:"conversationId"`
	AuthCode              string                `json:"authCode"`
	HostReference         string                `json:"hostReference"`
	PaymentType           PaymentType           `json:"paymentType"`
	TransId               string                `json:"transId"`
	PaymentId             int64                 `json:"paymentId"`
	PaymentError          PaymentError          `json:"paymentError"`
}

type ReportingPaymentTransactionRefundResponse struct {
	Id                    int64                 `json:"id"`
	CreatedDate           Time                  `json:"createdDate"`
	Status                RefundStatus          `json:"status"`
	RefundDestinationType RefundDestinationType `json:"refundDestinationType"`
	RefundPrice           float64               `json:"refundPrice"`
	RefundBankPrice       float64               `json:"refundBankPrice"`
	RefundWalletPrice     float64               `json:"refundWalletPrice"`
	ConversationId        string                `json:"conversationId"`
	AuthCode              string                `json:"authCode"`
	HostReference         string                `json:"hostReference"`
	TransId               string                `json:"transId"`
	IsAfterSettlement     bool                  `json:"isAfterSettlement"`
	PaymentTransactionId  int64                 `json:"paymentTransactionId"`
	PaymentError          PaymentError          `json:"paymentError"`
}

type ReportingPaymentTransaction struct {
	Id                            int64               `json:"id"`
	Name                          string              `json:"name"`
	ExternalId                    string              `json:"externalId"`
	Price                         float64             `json:"price"`
	PaidPrice                     float64             `json:"paidPrice"`
	WalletPrice                   float64             `json:"walletPrice"`
	MerchantCommissionRate        float64             `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount  float64             `json:"merchantCommissionRateAmount"`
	MerchantPayoutAmount          float64             `json:"merchantPayoutAmount"`
	SubMerchantMemberId           int64               `json:"subMerchantMemberId"`
	SubMerchantMemberPrice        float64             `json:"subMerchantMemberPrice"`
	SubMerchantMemberPayoutRate   float64             `json:"subMerchantMemberPayoutRate"`
	SubMerchantMemberPayoutAmount float64             `json:"subMerchantMemberPayoutAmount"`
	TransactionStatus             TransactionStatus   `json:"transactionStatus"`
	BlockageResolvedDate          Time                `json:"blockageResolvedDate"`
	CreatedDate                   Time                `json:"createdDate"`
	TransactionStatusDate         Time                `json:"transactionStatusDate"`
	RefundablePrice               float64             `json:"refundablePrice"`
	BankCommissionRate            float64             `json:"bankCommissionRate"`
	BankCommissionRateAmount      float64             `json:"bankCommissionRateAmount"`
	Payout                        Payout              `json:"payout"`
	SubMerchantMember             MemberResponse      `json:"subMerchantMember"`
	RefundStatus                  PaymentRefundStatus `json:"refundStatus"`
	PayoutStatus                  PayoutStatus        `json:"payoutStatus"`
}

func (api *PaymentReporting) SearchPayments(request SearchPaymentsRequest) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/payment-reporting/v1/payments?", api.Opts.BaseURL), nil)

	//req.URL.RawQuery, _ = QueryParams(request)

	res := Response[DataResponse[ReportingPaymentResponse]]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *PaymentReporting) SearchPaymentRefunds(request SearchPaymentRefundsRequest) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/payment-reporting/v1/refunds", api.Opts.BaseURL), nil)

	res := Response[DataResponse[ReportingPaymentRefundResponse]]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *PaymentReporting) SearchPaymentTransactionRefunds(request SearchPaymentTransactionRefundsRequest) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/payment-reporting/v1/refund-transactions", api.Opts.BaseURL), nil)

	res := Response[DataResponse[ReportingPaymentTransactionRefundResponse]]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *PaymentReporting) RetrievePayment(id int64) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/payment-reporting/v1/payments/%d", api.Opts.BaseURL, id), nil)

	res := Response[ReportingPaymentResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *PaymentReporting) RetrievePaymentTransactions(id int64) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/payment-reporting/v1/payments/%d/transactions", api.Opts.BaseURL, id), nil)

	res := Response[DataResponse[ReportingPaymentTransaction]]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *PaymentReporting) RetrievePaymentRefunds(id int64) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/payment-reporting/v1/payments/%d/refunds", api.Opts.BaseURL, id), nil)

	res := Response[DataResponse[ReportingPaymentRefundResponse]]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *PaymentReporting) RetrievePaymentTransactionRefunds(paymentId, paymentTransactionId int64) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/payment-reporting/v1/payments/%d/transactions/%d/refunds", api.Opts.BaseURL, paymentId, paymentTransactionId), nil)

	res := Response[DataResponse[ReportingPaymentTransactionRefundResponse]]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}
