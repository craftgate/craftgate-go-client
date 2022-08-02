package adapter

import (
	"craftgate-go-client/model"
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
	ID                            int     `json:"id"`
	ExternalID                    string  `json:"externalId"`
	Name                          string  `json:"name"`
	TransactionStatus             string  `json:"transactionStatus"`
	BlockageResolvedDate          string  `json:"blockageResolvedDate"`
	Price                         float64 `json:"price"`
	PaidPrice                     float64 `json:"paidPrice"`
	WalletPrice                   int     `json:"walletPrice"`
	MerchantCommissionRate        int     `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount  float64 `json:"merchantCommissionRateAmount"`
	MerchantPayoutAmount          float64 `json:"merchantPayoutAmount"`
	SubMerchantMemberID           int     `json:"subMerchantMemberId"`
	SubMerchantMemberPrice        float64 `json:"subMerchantMemberPrice"`
	SubMerchantMemberPayoutRate   int     `json:"subMerchantMemberPayoutRate"`
	SubMerchantMemberPayoutAmount float64 `json:"subMerchantMemberPayoutAmount"`
	//Payout                        Payout  `json:"payout"`
}

type PaymentResponse struct {
	Id                           int64                        `json:"id"`
	CreatedDate                  TimeResponse                 `json:"createdDate"`
	Price                        float64                      `json:"price"`
	PaidPrice                    float64                      `json:"paidPrice"`
	WalletPrice                  float64                      `json:"walletPrice"`
	Currency                     string                       `json:"currency"`
	BuyerMemberId                int64                        `json:"buyerMemberId"`
	Installment                  int                          `json:"installment"`
	ConversationId               string                       `json:"conversationId"`
	ExternalId                   string                       `json:"externalId"`
	PaymentType                  *model.PaymentType           `json:"paymentType"`
	PaymentProvider              *model.PaymentProvider       `json:"paymentProvider"`
	PaymentSource                *model.PaymentSource         `json:"paymentSource"`
	PaymentGroup                 *model.PaymentGroup          `json:"paymentGroup"`
	PaymentStatus                *model.PaymentStatus         `json:"paymentStatus"`
	PaymentPhase                 *model.PaymentPhase          `json:"paymentPhase"`
	PaymentChannel               string                       `json:"paymentChannel"`
	IsThreeDS                    bool                         `json:"isThreeDS"`
	MerchantCommissionRate       float64                      `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount float64                      `json:"merchantCommissionRateAmount"`
	BankCommissionRate           float64                      `json:"bankCommissionRate"`
	BankCommissionRateAmount     float64                      `json:"bankCommissionRateAmount"`
	PaidWithStoredCard           bool                         `json:"paidWithStoredCard"`
	BinNumber                    string                       `json:"binNumber"`
	LastFourDigits               string                       `json:"lastFourDigits"`
	AuthCode                     string                       `json:"authCode"`
	HostReference                string                       `json:"hostReference"`
	TransId                      string                       `json:"transId"`
	OrderId                      string                       `json:"orderId"`
	CardHolderName               string                       `json:"cardHolderName"`
	BankCardHolderName           string                       `json:"bankCardHolderName"`
	CardType                     string                       `json:"cardType"`
	CardAssociation              string                       `json:"cardAssociation"`
	CardBrand                    string                       `json:"cardBrand"`
	RequestedPosAlias            string                       `json:"requestedPosAlias"`
	Pos                          model.MerchantPos            `json:"pos"`
	Loyalty                      model.Loyalty                `json:"loyalty"`
	PaymentError                 model.PaymentError           `json:"paymentError"`
	CardUserKey                  string                       `json:"cardUserKey"`
	CardToken                    string                       `json:"cardToken"`
	PaymentTransactions          []PaymentTransactionResponse `json:"paymentTransactions"`
}

/*
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

func (api *Payment) PostAuthPayment(request PostAuthPaymentRequest) (interface{}, error) {
	body, _ := PrepareBody(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/payment/v1/card-payments/%d/post-auth", api.Opts.BaseURL, request.PaymentId), nil)
	res := model.Response[PaymentResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Payment) InitCheckoutPayment(request InitCheckoutPaymentRequest) (interface{}, error) {
	body, _ := PrepareBody(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/payment/v1/checkout-payments/init", api.Opts.BaseURL), nil)
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

func (api *Payment) Init3DSDepositPayment(request Init3DSDepositPaymentRequest) (interface{}, error) {
	body, _ := PrepareBody(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/payment/v1/deposits/3ds-init", api.Opts.BaseURL), body)
	res := model.Response[DepositPaymentResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *Payment) Complete3DSDepositPayment(request Complete3DSDepositPaymentRequest) (interface{}, error) {
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

func (api *Payment) SearchStoredCards(request SearchStoredCardsRequest) error {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/payment/v1/cards", api.Opts.BaseURL), nil)
	req.URL.RawQuery, _ = QueryParams(request)
	res := model.Response[model.DataResponse[StoredCardResponse]]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return resErr
}
*/
/*


public PaymentTransactionApprovalListResponse ApprovePaymentTransactions(
ApprovePaymentTransactionsRequest approvePaymentTransactionsRequest)
{
var path = "/payment/v1/payment-transactions/approve";
return RestClient.Post<PaymentTransactionApprovalListResponse>(RequestOptions.BaseUrl + path,
CreateHeaders(approvePaymentTransactionsRequest, path, RequestOptions),
approvePaymentTransactionsRequest);
}

public PaymentTransactionApprovalListResponse DisapprovePaymentTransactions(
DisapprovePaymentTransactionsRequest disapprovePaymentTransactionsRequest)
{
var path = "/payment/v1/payment-transactions/disapprove";
return RestClient.Post<PaymentTransactionApprovalListResponse>(RequestOptions.BaseUrl + path,
CreateHeaders(disapprovePaymentTransactionsRequest, path, RequestOptions),
disapprovePaymentTransactionsRequest);
}

public CheckMasterpassUserResponse CheckMasterpassUser(CheckMasterpassUserRequest checkMasterpassUserRequest)
{
var path = "/payment/v1/masterpass-payments/check-user";
return RestClient.Post<CheckMasterpassUserResponse>(RequestOptions.BaseUrl + path,
CreateHeaders(checkMasterpassUserRequest, path, RequestOptions), checkMasterpassUserRequest);
}
}

*/
