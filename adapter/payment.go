package adapter

import (
	"craftgate-go-client/adapter/rest"
	"craftgate-go-client/model"
	"fmt"
	"net/http"
)

type Payment struct {
	Opts model.RequestOptions
}

//type Payout struct {
//	PaidPrice                     float64 `json:"paidPrice"`
//	Currency                      string  `json:"currency"`
//	MerchantPayoutAmount          float64 `json:"merchantPayoutAmount"`
//	SubMerchantMemberPayoutAmount float64 `json:"subMerchantMemberPayoutAmount"`
//}

type PaymentTransactionResponse struct {
	ID                            *int     `json:"id"`
	ExternalID                    *string  `json:"externalId"`
	Name                          *string  `json:"name"`
	TransactionStatus             *string  `json:"transactionStatus"`
	BlockageResolvedDate          *string  `json:"blockageResolvedDate"`
	Price                         *float64 `json:"price"`
	PaidPrice                     *float64 `json:"paidPrice"`
	WalletPrice                   *int     `json:"walletPrice"`
	MerchantCommissionRate        *int     `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount  *float64 `json:"merchantCommissionRateAmount"`
	MerchantPayoutAmount          *float64 `json:"merchantPayoutAmount"`
	SubMerchantMemberID           *int     `json:"subMerchantMemberId"`
	SubMerchantMemberPrice        *float64 `json:"subMerchantMemberPrice"`
	SubMerchantMemberPayoutRate   *int     `json:"subMerchantMemberPayoutRate"`
	SubMerchantMemberPayoutAmount *float64 `json:"subMerchantMemberPayoutAmount"`
	//Payout                        Payout  `json:"payout"`
}

type PaymentResponse struct {
	Id                           *int64                       `json:"id"`
	CreatedDate                  *TimeResponse                `json:"createdDate"`
	Price                        *float64                     `json:"price"`
	PaidPrice                    *float64                     `json:"paidPrice"`
	WalletPrice                  *float64                     `json:"walletPrice"`
	Currency                     *string                      `json:"currency"`
	BuyerMemberId                *int64                       `json:"buyerMemberId"`
	Installment                  *int                         `json:"installment"`
	ConversationId               *string                      `json:"conversationId"`
	ExternalId                   *string                      `json:"externalId"`
	PaymentType                  *model.PaymentType           `json:"paymentType"`
	PaymentProvider              *model.PaymentProvider       `json:"paymentProvider"`
	PaymentSource                *model.PaymentSource         `json:"paymentSource"`
	PaymentGroup                 *model.PaymentGroup          `json:"paymentGroup"`
	PaymentStatus                *model.PaymentStatus         `json:"paymentStatus"`
	PaymentPhase                 *model.PaymentPhase          `json:"paymentPhase"`
	PaymentChannel               *string                      `json:"paymentChannel"`
	IsThreeDS                    *bool                        `json:"isThreeDS"`
	MerchantCommissionRate       *float64                     `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount *float64                     `json:"merchantCommissionRateAmount"`
	BankCommissionRate           *float64                     `json:"bankCommissionRate"`
	BankCommissionRateAmount     *float64                     `json:"bankCommissionRateAmount"`
	PaidWithStoredCard           *bool                        `json:"paidWithStoredCard"`
	BinNumber                    *string                      `json:"binNumber"`
	LastFourDigits               *string                      `json:"lastFourDigits"`
	AuthCode                     *string                      `json:"authCode"`
	HostReference                *string                      `json:"hostReference"`
	TransId                      *string                      `json:"transId"`
	OrderId                      *string                      `json:"orderId"`
	CardHolderName               *string                      `json:"cardHolderName"`
	BankCardHolderName           *string                      `json:"bankCardHolderName"`
	CardType                     *string                      `json:"cardType"`
	CardAssociation              *string                      `json:"cardAssociation"`
	CardBrand                    *string                      `json:"cardBrand"`
	RequestedPosAlias            *string                      `json:"requestedPosAlias"`
	Pos                          *model.MerchantPos           `json:"pos"`
	Loyalty                      *model.Loyalty               `json:"loyalty"`
	PaymentError                 *model.PaymentError          `json:"paymentError"`
	CardUserKey                  *string                      `json:"cardUserKey"`
	CardToken                    *string                      `json:"cardToken"`
	PaymentTransactions          []PaymentTransactionResponse `json:"paymentTransactions"`
}

type CreatePaymentRequest struct {
	Price            float64                `json:"price,omitempty"`
	PaidPrice        float64                `json:"paidPrice,omitempty"`
	WalletPrice      float64                `json:"walletPrice,omitempty"`
	PosAlias         string                 `json:"posAlias,omitempty"`
	Installment      int                    `json:"installment,omitempty"`
	Currency         model.Currency         `json:"currency,omitempty"`
	PaymentGroup     model.PaymentGroup     `json:"paymentGroup,omitempty"`
	ConversationId   string                 `json:"conversationId,omitempty"`
	ExternalId       string                 `json:"externalId,omitempty"`
	ClientIp         string                 `json:"clientIp,omitempty"`
	PaymentPhase     model.PaymentPhase     `json:"paymentPhase,omitempty"`
	PaymentChannel   string                 `json:"paymentChannel,omitempty"`
	BuyerMemberId    int64                  `json:"buyerMemberId,omitempty"`
	BankOrderId      string                 `json:"bankOrderId,omitempty"`
	Card             model.Card             `json:"card,omitempty"`
	Items            []model.PaymentItem    `json:"items"`
	AdditionalParams map[string]interface{} `json:"additionalParams"`
}

type Init3DSPaymentRequest struct {
	Price            float64                `json:"price,omitempty"`
	PaidPrice        float64                `json:"paidPrice,omitempty"`
	WalletPrice      float64                `json:"walletPrice,omitempty"`
	PosAlias         string                 `json:"posAlias,omitempty"`
	Installment      int                    `json:"installment,omitempty"`
	Currency         model.Currency         `json:"currency,omitempty"`
	PaymentGroup     model.PaymentGroup     `json:"paymentGroup,omitempty"`
	ConversationId   string                 `json:"conversationId,omitempty"`
	ExternalId       string                 `json:"externalId,omitempty"`
	ClientIp         string                 `json:"clientIp,omitempty"`
	PaymentPhase     model.PaymentPhase     `json:"paymentPhase,omitempty"`
	PaymentChannel   string                 `json:"paymentChannel,omitempty"`
	BuyerMemberId    int64                  `json:"buyerMemberId,omitempty"`
	BankOrderId      string                 `json:"bankOrderId,omitempty"`
	Card             model.Card             `json:"card,omitempty"`
	CallbackUrl      string                 `json:"callbackUrl,omitempty"`
	Items            []model.PaymentItem    `json:"items"`
	AdditionalParams map[string]interface{} `json:"additionalParams"`
}

type Init3DSPaymentResponse struct {
	HtmlContent *string `json:"htmlContent"`
}

type Complete3DSPaymentRequest struct {
	PaymentId int64 `json:"paymentId"`
}

type PostAuthPaymentRequest struct {
	PaidPrice float64 `json:"paidPrice"`
}

type InitCheckoutPaymentRequest struct {
	Price                       float64               `json:"price,omitempty"`
	PaidPrice                   float64               `json:"paidPrice,omitempty"`
	Currency                    model.Currency        `json:"currency,omitempty"`
	PaymentGroup                model.PaymentGroup    `json:"paymentGroup,omitempty"`
	ConversationId              string                `json:"conversationId,omitempty"`
	ExternalId                  string                `json:"externalId,omitempty"`
	BankOrderId                 string                `json:"bankOrderId,omitempty"`
	CallbackUrl                 string                `json:"callbackUrl,omitempty"`
	ClientIp                    string                `json:"clientIp,omitempty"`
	PaymentPhase                model.PaymentPhase    `json:"paymentPhase,omitempty"`
	PaymentChannel              string                `json:"paymentChannel,omitempty"`
	EnabledPaymentMethods       []model.PaymentMethod `json:"enabledPaymentMethods,omitempty"`
	MasterpassGsmNumber         string                `json:"masterpassGsmNumber,omitempty"`
	MasterpassUserId            string                `json:"masterpassUserId,omitempty"`
	CardUserKey                 string                `json:"cardUserKey,omitempty"`
	BuyerMemberId               int64                 `json:"buyerMemberId,omitempty"`
	EnabledInstallments         []int                 `json:"enabledInstallments,omitempty"`
	AlwaysStoreCardAfterPayment bool                  `json:"alwaysStoreCardAfterPayment,omitempty"`
	AllowOnlyStoredCards        bool                  `json:"allowOnlyStoredCards,omitempty"`
	AllowOnlyCreditCard         bool                  `json:"allowOnlyCreditCard,omitempty"`
	ForceThreeDS                bool                  `json:"forceThreeDS,omitempty"`
	Items                       []model.PaymentItem   `json:"items"`
}

type InitCheckoutPaymentResponse struct {
	Token   *string `json:"token"`
	PageUrl *string `json:"pageUrl"`
}

type DepositPaymentRequest struct {
	Price          float64    `json:"price,omitempty"`
	BuyerMemberId  int64      `json:"buyerMemberId,omitempty"`
	ConversationId string     `json:"conversationId,omitempty"`
	CallbackUrl    string     `json:"callbackUrl,omitempty"`
	PosAlias       string     `json:"posAlias,omitempty"`
	ClientIp       string     `json:"clientIp,omitempty"`
	Card           model.Card `json:"card"`
}

type DepositPaymentResponse struct {
	Id                       *int64               `json:"id"`
	CreatedDate              *TimeResponse        `json:"createdDate"`
	Price                    *float64             `json:"price"`
	Currency                 *string              `json:"currency"`
	BuyerMemberId            *int64               `json:"buyerMemberId"`
	ConversationId           *string              `json:"conversationId"`
	BankCommissionRate       *float64             `json:"bankCommissionRate"`
	BankCommissionRateAmount *float64             `json:"bankCommissionRateAmount"`
	AuthCode                 *string              `json:"authCode"`
	HostReference            *string              `json:"hostReference"`
	TransId                  *string              `json:"transId"`
	OrderId                  *string              `json:"orderId"`
	PaymentType              *model.PaymentType   `json:"paymentType"`
	PaymentStatus            *model.PaymentStatus `json:"paymentStatus"`
	CardUserKey              *string              `json:"cardUserKey"`
	CardToken                *string              `json:"cardToken"`
}

type CreateFundTransferDepositPaymentRequest struct {
	Price          float64 `json:"price,omitempty"`
	BuyerMemberId  int64   `json:"buyerMemberId,omitempty"`
	ConversationId string  `json:"conversationId,omitempty"`
	ClientIp       string  `json:"clientIp,omitempty"`
}

type InitGarantiPayPaymentRequest struct {
	Price          float64                 `json:"price,omitempty"`
	PaidPrice      float64                 `json:"paidPrice,omitempty"`
	Currency       model.Currency          `json:"currency,omitempty"`
	PosAlias       string                  `json:"posAlias,omitempty"`
	PaymentGroup   model.PaymentGroup      `json:"paymentGroup,omitempty"`
	ConversationId string                  `json:"conversationId,omitempty"`
	ExternalId     string                  `json:"externalId,omitempty"`
	CallbackUrl    string                  `json:"callbackUrl,omitempty"`
	ClientIp       string                  `json:"clientIp,omitempty"`
	PaymentChannel string                  `json:"paymentChannel,omitempty"`
	BuyerMemberId  int64                   `json:"buyerMemberId,omitempty"`
	BankOrderId    string                  `json:"bankOrderId,omitempty"`
	Items          []model.PaymentItem     `json:"items"`
	Installments   []GarantiPayInstallment `json:"installments,omitempty"`
}

type GarantiPayInstallment struct {
	Number     int     `json:"number,omitempty"`
	TotalPrice float64 `json:"totalPrice,omitempty"`
}

type InitGarantiPayPaymentResponse struct {
	HtmlContent *string `json:"htmlContent"`
}

type RetrieveLoyaltiesRequest struct {
	CardNumber  string `json:"cardNumber"`
	ExpireYear  string `json:"expireYear"`
	ExpireMonth string `json:"expireMonth"`
	Cvc         string `json:"cvc"`
	CardUserKey string `json:"cardUserKey"`
	CardToken   string `json:"cardToken"`
}

type RetrieveLoyaltiesResponse struct {
	CardBrand *string            `json:"cardBrand"`
	Force3ds  *bool              `json:"force3Ds"`
	Pos       *model.MerchantPos `json:"pos"`
	Loyalties []model.Loyalty    `json:"loyalties"`
}

type RefundPaymentTransactionRequest struct {
	PaymentTransactionId  int64                       `json:"paymentTransactionId"`
	ConversationId        string                      `json:"conversationId"`
	RefundPrice           float64                     `json:"refundPrice"`
	RefundDestinationType model.RefundDestinationType `json:"refundDestinationType"`
	ChargeFromMe          bool                        `json:"chargeFromMe"`
}

type PaymentTransactionRefundResponse struct {
	Id                    *int64                       `json:"id"`
	CreatedDate           *TimeResponse                `json:"createdDate"`
	Status                *model.RefundStatus          `json:"status"`
	RefundDestinationType *model.RefundDestinationType `json:"refundDestinationType"`
	RefundPrice           *float64                     `json:"refundPrice"`
	RefundBankPrice       *float64                     `json:"refundBankPrice"`
	RefundWalletPrice     *float64                     `json:"refundWalletPrice"`
	ConversationId        *string                      `json:"conversationId"`
	AuthCode              *string                      `json:"authCode"`
	HostReference         *string                      `json:"hostReference"`
	TransId               *string                      `json:"transId"`
	IsAfterSettlement     *bool                        `json:"isAfterSettlement"`
	PaymentTransactionId  *int64                       `json:"paymentTransactionId"`
	Currency              *model.Currency              `json:"currency"`
	PaymentId             *int64                       `json:"paymentId"`
}

type RefundPaymentRequest struct {
	PaymentId             int64                       `json:"paymentId,omitempty"`
	ConversationId        string                      `json:"conversationId,omitempty"`
	RefundDestinationType model.RefundDestinationType `json:"refundDestinationType,omitempty"`
	ChargeFromMe          bool                        `json:"chargeFromMe,omitempty"`
}

type PaymentRefundResponse struct {
	Id                        *int64                             `json:"id"`
	CreatedDate               *TimeResponse                      `json:"createdDate"`
	Status                    *model.RefundStatus                `json:"status"`
	RefundDestinationType     *model.RefundDestinationType       `json:"refundDestinationType"`
	RefundPrice               *float64                           `json:"refundPrice"`
	RefundBankPrice           *float64                           `json:"refundBankPrice"`
	RefundWalletPrice         *float64                           `json:"refundWalletPrice"`
	ConversationId            *string                            `json:"conversationId"`
	AuthCode                  *string                            `json:"authCode"`
	HostReference             *string                            `json:"hostReference"`
	TransId                   *string                            `json:"transId"`
	PaymentId                 *int64                             `json:"paymentId"`
	RefundType                *model.RefundType                  `json:"refundType"`
	Currency                  *model.Currency                    `json:"currency"`
	PaymentTransactionRefunds []PaymentTransactionRefundResponse `json:"paymentTransactionRefunds"`
}

type StoreCardRequest struct {
	CardHolderName string `json:"cardHolderName,omitempty"`
	CardNumber     string `json:"cardNumber,omitempty"`
	ExpireYear     string `json:"expireYear,omitempty"`
	ExpireMonth    string `json:"expireMonth,omitempty"`
	CardAlias      string `json:"cardAlias,omitempty"`
	CardUserKey    string `json:"cardUserKey,omitempty"`
}

type StoredCardResponse struct {
	BinNumber       *string                `json:"binNumber"`
	LastFourDigits  *string                `json:"lastFourDigits"`
	CardUserKey     *string                `json:"cardUserKey"`
	CardToken       *string                `json:"cardToken"`
	CardAlias       *string                `json:"cardAlias"`
	CardType        *model.CardType        `json:"cardType"`
	CardAssociation *model.CardAssociation `json:"cardAssociation"`
	CardBrand       *string                `json:"cardBrand"`
	CardBankName    *string                `json:"cardBankName"`
	CardBankId      *int64                 `json:"cardBankId"`
}

type UpdateStoredCardRequest struct {
	CardUserKey string `json:"cardUserKey,omitempty"`
	CardToken   string `json:"cardToken,omitempty"`
	ExpireYear  string `json:"expireYear,omitempty"`
	ExpireMonth string `json:"expireMonth,omitempty"`
}

type DeleteStoredCardRequest struct {
	CardUserKey string `schema:"cardUserKey,omitempty"`
	CardToken   string `schema:"cardToken,omitempty"`
}

type SearchStoredCardsRequest struct {
	CardAlias       string                `schema:"cardAlias,omitempty"`
	CardBrand       string                `schema:"cardBrand,omitempty"`
	CardType        model.CardType        `schema:"cardType,omitempty"`
	CardUserKey     string                `schema:"cardUserKey,omitempty"`
	CardToken       string                `schema:"cardToken,omitempty"`
	CardBankName    string                `schema:"cardBankName,omitempty"`
	CardAssociation model.CardAssociation `schema:"cardAssociation,omitempty"`
	Page            int                   `schema:"page,omitempty"`
	Size            int                   `schema:"size,omitempty"`
}

type PaymentTransactionsApprovalRequest struct {
	PaymentTransactionIds []int64 `json:"paymentTransactionIds,omitempty"`
	IsTransactional       bool    `json:"isTransactional,omitempty"`
}

type PaymentTransactionsApprovalResponse struct {
	PaymentTransactionId *int64                `json:"paymentTransactionId"`
	ApprovalStatus       *model.ApprovalStatus `json:"approvalStatus"`
	FailedReason         *string               `json:"failedReason"`
}

type CheckMasterpassUserRequest struct {
	MasterpassGsmNumber string `json:"masterpassGsmNumber"`
}

type CheckMasterpassUserResponse struct {
	IsEligibleToUseMasterpass             *bool   `json:"isEligibleToUseMasterpass"`
	IsAnyCardSavedInCustomerProgram       *bool   `json:"isAnyCardSavedInCustomerProgram"`
	HasMasterpassAccount                  *bool   `json:"hasMasterpassAccount"`
	HashAnyCardSavedInMasterpassAccount   *bool   `json:"hashAnyCardSavedInMasterpassAccount"`
	IsMasterpassAccountLinkedWithMerchant *bool   `json:"isMasterpassAccountLinkedWithMerchant"`
	IsMasterpassAccountLocked             *bool   `json:"isMasterpassAccountLocked"`
	IsPhoneNumberUpdatedInAnotherMerchant *bool   `json:"isPhoneNumberUpdatedInAnotherMerchant"`
	AccountStatus                         *string `json:"accountStatus"`
}

type InitApmPaymentRequest struct {
	ApmType          model.ApmType          `json:"apmType,omitempty"`
	Price            float64                `json:"price,omitempty"`
	PaidPrice        float64                `json:"paidPrice,omitempty"`
	Currency         model.Currency         `json:"currency,omitempty"`
	PaymentGroup     model.PaymentGroup     `json:"paymentGroup,omitempty"`
	PaymentChannel   string                 `json:"paymentChannel,omitempty"`
	ConversationId   string                 `json:"conversationId,omitempty"`
	ExternalId       string                 `json:"externalId,omitempty"`
	CallbackUrl      string                 `json:"callbackUrl,omitempty"`
	ApmOrderId       string                 `json:"apmOrderId,omitempty"`
	ApmUserIdentity  string                 `json:"apmUserIdentity,omitempty"`
	ClientIp         string                 `json:"clientIp,omitempty"`
	Items            []model.PaymentItem    `json:"items"`
	AdditionalParams map[string]interface{} `json:"additionalParams"`
}

type InitApmPaymentResponse struct {
	PaymentId        *int64                     `json:"paymentId"`
	RedirectUrl      *string                    `json:"redirectUrl"`
	PaymentStatus    *model.PaymentStatus       `json:"paymentStatus"`
	AdditionalAction *model.ApmAdditionalAction `json:"additionalAction"`
}

type CompleteApmPaymentRequest struct {
	PaymentId        int64                  `json:"paymentId"`
	AdditionalParams map[string]interface{} `json:"additionalParams"`
}

type CompleteApmPaymentResponse struct {
	PaymentId     *int64               `json:"paymentId"`
	PaymentStatus *model.PaymentStatus `json:"paymentStatus"`
}

func (api *Payment) CreatePayment(request CreatePaymentRequest) (interface{}, error) {
	body, _ := PrepareBody(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/payment/v1/card-payments", api.Opts.BaseURL), body)
	res := model.Response[PaymentResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Payment) RetrievePayment(paymentId int64) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/payment/v1/card-payments/%d", api.Opts.BaseURL, paymentId), nil)
	res := model.Response[PaymentResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Payment) Init3DSPayment(request Init3DSPaymentRequest) (interface{}, error) {
	body, _ := PrepareBody(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/payment/v1/card-payments/3ds-init", api.Opts.BaseURL), body)
	res := model.Response[Init3DSPaymentResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Payment) Complete3DSPayment(request Complete3DSPaymentRequest) (interface{}, error) {
	body, _ := PrepareBody(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/payment/v1/card-payments/3ds-complete", api.Opts.BaseURL), body)
	res := model.Response[PaymentResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Payment) PostAuthPayment(paymentId int64, request PostAuthPaymentRequest) (interface{}, error) {
	body, _ := PrepareBody(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/payment/v1/card-payments/%d/post-auth", api.Opts.BaseURL, paymentId), body)
	res := model.Response[PaymentResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Payment) InitCheckoutPayment(request InitCheckoutPaymentRequest) (interface{}, error) {
	body, _ := PrepareBody(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/payment/v1/checkout-payments/init", api.Opts.BaseURL), body)
	res := model.Response[InitCheckoutPaymentResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Payment) RetrieveCheckoutPayment(token string) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/payment/v1/checkout-payments/%s", api.Opts.BaseURL, token), nil)
	res := model.Response[PaymentResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Payment) CreateDepositPayment(request DepositPaymentRequest) (interface{}, error) {
	body, _ := PrepareBody(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/payment/v1/deposits", api.Opts.BaseURL), body)
	res := model.Response[DepositPaymentResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Payment) Init3DSDepositPayment(request DepositPaymentRequest) (interface{}, error) {
	body, _ := PrepareBody(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/payment/v1/deposits/3ds-init", api.Opts.BaseURL), body)
	res := model.Response[DepositPaymentResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Payment) Complete3DSDepositPayment(request Complete3DSPaymentRequest) (interface{}, error) {
	body, _ := PrepareBody(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/payment/v1/deposits/3ds-complete", api.Opts.BaseURL), body)
	res := model.Response[DepositPaymentResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Payment) CreateFundTransferDepositPayment(request CreateFundTransferDepositPaymentRequest) error {
	body, _ := PrepareBody(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/payment/v1/deposits/fund-transfer", api.Opts.BaseURL), body)
	res := model.Void{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return resErr
}

func (api *Payment) InitGarantiPayPayment(request InitGarantiPayPaymentRequest) (interface{}, error) {
	body, _ := PrepareBody(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/payment/v1/garanti-pay-payments", api.Opts.BaseURL), body)
	res := model.Response[InitGarantiPayPaymentResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Payment) InitApmPayment(request InitApmPaymentRequest) (interface{}, error) {
	body, _ := PrepareBody(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/payment/v1/apm-payments/init", api.Opts.BaseURL), body)
	res := model.Response[InitApmPaymentResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Payment) CompleteApmPayment(request CompleteApmPaymentRequest) (interface{}, error) {
	body, _ := PrepareBody(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/payment/v1/apm-payments/complete", api.Opts.BaseURL), body)
	res := model.Response[CompleteApmPaymentResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Payment) RetrieveLoyalties(request RetrieveLoyaltiesRequest) (interface{}, error) {
	body, _ := PrepareBody(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/payment/v1/card-loyalties/retrieve", api.Opts.BaseURL), body)
	res := model.Response[RetrieveLoyaltiesResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Payment) RefundPaymentTransaction(request RefundPaymentTransactionRequest) (interface{}, error) {
	body, _ := PrepareBody(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/payment/v1/refund-transactions", api.Opts.BaseURL), body)
	res := model.Response[PaymentTransactionRefundResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Payment) RetrievePaymentTransactionRefund(id int64) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/payment/v1/refund-transactions/%d", api.Opts.BaseURL, id), nil)
	res := model.Response[PaymentTransactionRefundResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Payment) RefundPayment(request RefundPaymentRequest) (interface{}, error) {
	body, _ := PrepareBody(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/payment/v1/refunds", api.Opts.BaseURL), body)
	res := model.Response[PaymentRefundResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Payment) RetrievePaymentRefund(id int64) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/payment/v1/refunds/%d", api.Opts.BaseURL, id), nil)
	res := model.Response[PaymentRefundResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Payment) StoreCard(request StoreCardRequest) (interface{}, error) {
	body, _ := PrepareBody(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/payment/v1/cards", api.Opts.BaseURL), body)
	res := model.Response[StoredCardResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Payment) UpdateStoredCard(request UpdateStoredCardRequest) (interface{}, error) {
	body, _ := PrepareBody(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/payment/v1/cards/update", api.Opts.BaseURL), body)
	res := model.Response[StoredCardResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Payment) DeleteStoredCard(request DeleteStoredCardRequest) error {
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/payment/v1/cards", api.Opts.BaseURL), nil)
	req.URL.RawQuery, _ = QueryParams(request)
	res := model.Void{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return resErr
}

func (api *Payment) SearchStoredCards(request SearchStoredCardsRequest) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/payment/v1/cards", api.Opts.BaseURL), nil)
	req.URL.RawQuery, _ = QueryParams(request)
	res := model.Response[model.DataResponse[StoredCardResponse]]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return res, resErr
}
func (api *Payment) ApprovePaymentTransactions(request PaymentTransactionsApprovalRequest) (interface{}, error) {
	body, _ := PrepareBody(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/payment/v1/payment-transactions/approve", api.Opts.BaseURL), body)
	res := model.Response[model.DataResponse[PaymentTransactionsApprovalResponse]]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return res, resErr
}

func (api *Payment) DisapprovePaymentTransactions(request PaymentTransactionsApprovalRequest) (interface{}, error) {
	body, _ := PrepareBody(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/payment/v1/payment-transactions/disapprove", api.Opts.BaseURL), body)
	res := model.Response[model.DataResponse[PaymentTransactionsApprovalResponse]]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return res, resErr
}

func (api *Payment) CheckMasterpassUser(request CheckMasterpassUserRequest) (interface{}, error) {
	body, _ := PrepareBody(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/payment/v1/masterpass-payments/check-user", api.Opts.BaseURL), body)
	res := model.Response[model.DataResponse[CheckMasterpassUserResponse]]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return res, resErr
}
