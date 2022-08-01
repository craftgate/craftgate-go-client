package adapter

import (
	"craftgate-go-client/adapter/rest"
	"craftgate-go-client/model"
	"fmt"
	"net/http"
	"net/url"
)

type PaymentReporting struct {
	Opts model.RequestOptions
}

type SearchPaymentsRequest struct {
	Page                 int                   `schema:"page,omitempty"`
	Size                 int                   `schema:"size,omitempty"`
	PaymentId            int64                 `schema:"paymentId,omitempty"`
	PaymentTransactionId int64                 `schema:"paymentTransactionId,omitempty"`
	BuyerMemberId        int64                 `schema:"buyerMemberId,omitempty"`
	SubMerchantMemberId  int64                 `schema:"subMerchantMemberId,omitempty"`
	ConversationId       string                `schema:"conversationId,omitempty"`
	ExternalId           string                `schema:"externalId,omitempty"`
	OrderId              string                `schema:"orderId,omitempty"`
	PaymentType          model.PaymentType     `schema:"paymentType,omitempty"`
	PaymentProvider      model.PaymentProvider `schema:"paymentProvider,omitempty"`
	PaymentStatus        model.PaymentStatus   `schema:"paymentStatus,omitempty"`
	PaymentSource        model.PaymentSource   `schema:"paymentSource,omitempty"`
	PaymentChannel       string                `schema:"paymentChannel,omitempty"`
	BinNumber            string                `schema:"binNumber,omitempty"`
	LastFourDigits       string                `schema:"lastFourDigits,omitempty"`
	Currency             model.Currency        `schema:"currency,omitempty"`
	MinPaidPrice         float64               `schema:"minPaidPrice,omitempty"`
	MaxPaidPrice         float64               `schema:"maxPaidPrice,omitempty"`
	Installment          int                   `schema:"installment,omitempty"`
	IsThreeDS            bool                  `schema:"isThreeDS,omitempty"`
	MinCreatedDate       CraftgateTime         `schema:"minCreatedDate,omitempty"`
	MaxCreatedDate       CraftgateTime         `schema:"maxCreatedDate,omitempty"`
}

type SearchPaymentRefundsRequest struct {
	Page           int                `schema:"page,omitempty"`
	Size           int                `schema:"size,omitempty"`
	Id             int64              `schema:"id,omitempty"`
	PaymentId      int64              `schema:"paymentId,omitempty"`
	BuyerMemberId  int64              `schema:"buyerMemberId,omitempty"`
	ConversationId string             `schema:"conversationId,omitempty"`
	Status         model.RefundStatus `schema:"status,omitempty"`
	Currency       model.Currency     `schema:"currency,omitempty"`
	MinRefundPrice float64            `schema:"minRefundPrice,omitempty"`
	MaxRefundPrice float64            `schema:"maxRefundPrice,omitempty"`
	MinCreatedDate CraftgateTime      `schema:"minCreatedDate,omitempty"`
	MaxCreatedDate CraftgateTime      `schema:"maxCreatedDate,omitempty"`
}

type SearchPaymentTransactionRefundsRequest struct {
	Page                 int                `schema:"page,omitempty"`
	Size                 int                `schema:"size,omitempty"`
	Id                   int64              `schema:"id,omitempty"`
	PaymentId            int64              `schema:"paymentId,omitempty"`
	PaymentTransactionId int64              `schema:"paymentTransactionId,omitempty"`
	BuyerMemberId        int64              `schema:"buyerMemberId,omitempty"`
	ConversationId       string             `schema:"conversationId,omitempty"`
	Status               model.RefundStatus `schema:"status,omitempty"`
	Currency             model.Currency     `schema:"currency,omitempty"`
	IsAfterSettlement    bool               `schema:"isAfterSettlement,omitempty"`
	MinRefundPrice       float64            `schema:"minRefundPrice,omitempty"`
	MaxRefundPrice       float64            `schema:"maxRefundPrice,omitempty"`
	MinCreatedDate       CraftgateTime      `schema:"minCreatedDate,omitempty"`
	MaxCreatedDate       CraftgateTime      `schema:"maxCreatedDate,omitempty"`
}

type ReportingPaymentResponse struct {
	Id                           int64                            `json:"id"`
	CreatedDate                  CraftgateTime                    `json:"createdDate"`
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
	BuyerMember                  MemberResponse                   `json:"buyerMember"`
	Refunds                      []ReportingPaymentRefundResponse `json:"refunds"`
	Pos                          model.MerchantPos                `json:"pos"`
	Loyalty                      model.Loyalty                    `json:"loyalty"`
	PaymentError                 model.PaymentError               `json:"paymentError"`
}

type ReportingPaymentRefundResponse struct {
	Id                    int64                       `json:"id"`
	CreatedDate           CraftgateTime               `json:"createdDate"`
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

type ReportingPaymentTransactionRefundResponse struct {
	Id                    int64                       `json:"id"`
	CreatedDate           CraftgateTime               `json:"createdDate"`
	Status                model.RefundStatus          `json:"status"`
	RefundDestinationType model.RefundDestinationType `json:"refundDestinationType"`
	RefundPrice           float64                     `json:"refundPrice"`
	RefundBankPrice       float64                     `json:"refundBankPrice"`
	RefundWalletPrice     float64                     `json:"refundWalletPrice"`
	ConversationId        string                      `json:"conversationId"`
	AuthCode              string                      `json:"authCode"`
	HostReference         string                      `json:"hostReference"`
	TransId               string                      `json:"transId"`
	IsAfterSettlement     bool                        `json:"isAfterSettlement"`
	PaymentTransactionId  int64                       `json:"paymentTransactionId"`
	PaymentError          model.PaymentError          `json:"paymentError"`
}

type ReportingPaymentTransaction struct {
	Id                            int64                     `json:"id"`
	Name                          string                    `json:"name"`
	ExternalId                    string                    `json:"externalId"`
	Price                         float64                   `json:"price"`
	PaidPrice                     float64                   `json:"paidPrice"`
	WalletPrice                   float64                   `json:"walletPrice"`
	MerchantCommissionRate        float64                   `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount  float64                   `json:"merchantCommissionRateAmount"`
	MerchantPayoutAmount          float64                   `json:"merchantPayoutAmount"`
	SubMerchantMemberId           int64                     `json:"subMerchantMemberId"`
	SubMerchantMemberPrice        float64                   `json:"subMerchantMemberPrice"`
	SubMerchantMemberPayoutRate   float64                   `json:"subMerchantMemberPayoutRate"`
	SubMerchantMemberPayoutAmount float64                   `json:"subMerchantMemberPayoutAmount"`
	TransactionStatus             model.TransactionStatus   `json:"transactionStatus"`
	BlockageResolvedDate          CraftgateTime             `json:"blockageResolvedDate"`
	CreatedDate                   CraftgateTime             `json:"createdDate"`
	TransactionStatusDate         CraftgateTime             `json:"transactionStatusDate"`
	RefundablePrice               float64                   `json:"refundablePrice"`
	BankCommissionRate            float64                   `json:"bankCommissionRate"`
	BankCommissionRateAmount      float64                   `json:"bankCommissionRateAmount"`
	Payout                        model.Payout              `json:"payout"`
	SubMerchantMember             MemberResponse            `json:"subMerchantMember"`
	RefundStatus                  model.PaymentRefundStatus `json:"refundStatus"`
	PayoutStatus                  model.PayoutStatus        `json:"payoutStatus"`
}

func (api *PaymentReporting) SearchPayments(request SearchPaymentsRequest) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/payment-reporting/v1/payments?", api.Opts.BaseURL), nil)

	//q := request.ToParams(*req.URL)
	//req.URL.RawQuery = q.Encode()
	req.URL.RawQuery, _ = QueryParams(request)

	res := model.Response[model.DataResponse[ReportingPaymentResponse]]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *PaymentReporting) SearchPaymentRefunds(request SearchPaymentRefundsRequest) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/payment-reporting/v1/refunds", api.Opts.BaseURL), nil)

	req.URL.RawQuery, _ = QueryParams(request)
	res := model.Response[model.DataResponse[ReportingPaymentRefundResponse]]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *PaymentReporting) SearchPaymentTransactionRefunds(request SearchPaymentTransactionRefundsRequest) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/payment-reporting/v1/refund-transactions", api.Opts.BaseURL), nil)

	req.URL.RawQuery, _ = QueryParams(request)
	res := model.Response[model.DataResponse[ReportingPaymentTransactionRefundResponse]]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *PaymentReporting) RetrievePayment(id int64) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/payment-reporting/v1/payments/%d", api.Opts.BaseURL, id), nil)

	res := model.Response[ReportingPaymentResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *PaymentReporting) RetrievePaymentTransactions(id int64) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/payment-reporting/v1/payments/%d/transactions", api.Opts.BaseURL, id), nil)

	res := model.Response[model.DataResponse[ReportingPaymentTransaction]]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *PaymentReporting) RetrievePaymentRefunds(id int64) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/payment-reporting/v1/payments/%d/refunds", api.Opts.BaseURL, id), nil)

	res := model.Response[model.DataResponse[ReportingPaymentRefundResponse]]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *PaymentReporting) RetrievePaymentTransactionRefunds(paymentId, paymentTransactionId int64) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/payment-reporting/v1/payments/%d/transactions/%d/refunds", api.Opts.BaseURL, paymentId, paymentTransactionId), nil)

	res := model.Response[model.DataResponse[ReportingPaymentTransactionRefundResponse]]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (request SearchPaymentsRequest) ToParams(url url.URL) url.Values {
	q := url.Query()
	q.Add("currency", string(request.Currency))
	return q
}
