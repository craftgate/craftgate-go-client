package adapter

import "time"

type PaymentType string
type ApmType string
type PaymentProvider string
type PaymentStatus string
type PaymentSource string
type PaymentGroup string
type PaymentPhase string
type PaymentMethod string
type CardType string
type CardAssociation string
type CardExpiryStatus string
type Currency string
type LoyaltyType string
type PaymentRefundStatus string
type RefundStatus string
type RefundType string
type ApprovalStatus string
type Status string
type MemberType string
type SettlementType string
type SettlementEarningsDestination string
type RefundDestinationType string
type TransactionStatus string
type TransactionPayoutStatus string
type WalletTransactionRefundCardTransactionType string
type FraudAction string
type FraudCheckStatus string
type ApmAdditionalAction string
type ReportFileType string
type WalletTransactionType string

const (
	ApiKeyHeaderName        = "x-api-key"
	RandomHeaderName        = "x-rnd-key"
	AuthVersionHeaderName   = "x-auth-version"
	ClientVersionHeaderName = "x-client-version"
	SignatureHeaderName     = "x-signature"
	AuthVersion             = "1"
	ClientVersion           = "craftgate-go-client:1.0.0"
)

// payment type declaration
const (
	_                       PaymentType = ""
	CARD_PAYMENT                        = "CARD_PAYMENT"
	DEPOSIT_PAYMENT                     = "DEPOSIT_PAYMENT"
	WALLET_PAYMENT                      = "WALLET_PAYMENT"
	CARD_AND_WALLET_PAYMENT             = "CARD_AND_WALLET_PAYMENT"
	BANK_TRANSFER                       = "BANK_TRANSFER"
	APM                                 = "APM"
)

const (
	_                   ApmType = ""
	ApmPAPARA                   = "PAPARA"
	ApmPAYONEER                 = "PAYONEER"
	ApmSODEXO                   = "SODEXO"
	ApmEDENRED                  = "EDENRED"
	ApmPAYPAL                   = "PAYPAL"
	ApmFUND_TRANSFER            = "FUND_TRANSFER"
	ApmCASH_ON_DELIVERY         = "CASH_ON_DELIVERY"
)

// payment provider declaration
const (
	_           PaymentProvider = ""
	BANK                        = "BANK"
	CG_WALLET                   = "CG_WALLET"
	MASTERPASS                  = "MASTERPASS"
	GARANTI_PAY                 = "GARANTI_PAY"
	PAPARA                      = "PAPARA"
	PAYONEER                    = "PAYONEER"
	SODEXO                      = "SODEXO"
	EDENRED                     = "EDENRED"
)

// payment status declaration
const (
	_                PaymentStatus = ""
	FAILURE                        = "FAILURE"
	SUCCESS                        = "SUCCESS"
	INIT_THREEDS                   = "INIT_THREEDS"
	CALLBACK_THREEDS               = "CALLBACK_THREEDS"
	WAITING                        = "WAITING"
)

// payment source declaration
const (
	_             PaymentSource = ""
	API                         = "API"
	CHECKOUT_FORM               = "CHECKOUT_FORM"
	PAY_BY_LINK                 = "PAY_BY_LINK"
)

// currency declaration
const (
	_   Currency = ""
	TRY          = "TRY"
	USD          = "USD"
	EUR          = "EUR"
	GBP          = "GBP"
)

// payment group declaration
const (
	_                       PaymentGroup = ""
	PRODUCT                              = "PRODUCT"
	LISTING_OR_SUBSCRIPTION              = "LISTING_OR_SUBSCRIPTION"
)

// payment phase declaration
const (
	_         PaymentPhase = ""
	AUTH                   = "AUTH"
	PRE_AUTH               = "PRE_AUTH"
	POST_AUTH              = "POST_AUTH"
)

// payment method declaration
const (
	_                        PaymentMethod = ""
	PaymentMethod_CARD                     = "CARD"
	PaymentMethod_MASTERPASS               = "MASTERPASS"
)

// card type declaration
const (
	_            CardType = ""
	CREDIT_CARD           = "CREDIT_CARD"
	DEBIT_CARD            = "DEBIT_CARD"
	PREPAID_CARD          = "PREPAID_CARD"
)

// card association declaration
const (
	_           CardAssociation = ""
	VISA                        = "VISA"
	MASTER_CARD                 = "MASTER_CARD"
	AMEX                        = "AMEX"
	TROY                        = "TROY"
	JCB                         = "JCB"
	UNION_PAY                   = "UNION_PAY"
	MAESTRO                     = "MAESTRO"
	DISCOVER                    = "DISCOVER"
	DINERS_CLUB                 = "DINERS_CLUB"
)

// card expiry status declaration
const (
	_                      CardExpiryStatus = ""
	EXPIRED                                 = "EXPIRED"
	WILL_EXPIRE_NEXT_MONTH                  = "WILL_EXPIRE_NEXT_MONTH"
	NOT_EXPIRED                             = "NOT_EXPIRED"
)

// loyalty type declaration
const (
	_                      LoyaltyType = ""
	REWARD_MONEY                       = "REWARD_MONEY"
	ADDITIONAL_INSTALLMENT             = "ADDITIONAL_INSTALLMENT"
	POSTPONING_INSTALLMENT             = "POSTPONING_INSTALLMENT"
	EXTRA_POINTS                       = "EXTRA_POINTS"
	GAINING_MINUTES                    = "GAINING_MINUTES"
	POSTPONING_STATEMENT               = "POSTPONING_STATEMENT"
)

// payment refund status declaration
const (
	_                PaymentRefundStatus = ""
	NO_REFUND                            = "NO_REFUND"
	NOT_REFUNDED                         = "NOT_REFUNDED"
	PARTIAL_REFUNDED                     = "PARTIAL_REFUNDED"
	FULLY_REFUNDED                       = "FULLY_REFUNDED"
)

// refund status declaration
const (
	_                   RefundStatus = ""
	SuccessRefundStatus              = "SUCCESS"
	FailureRefundStatus              = "FAILURE"
)

// refund type declaration
const (
	_      RefundType = ""
	CANCEL            = "CANCEL"
	REFUND            = "REFUND"
)

// approval status declaration
const (
	_                     ApprovalStatus = ""
	SuccessApprovalStatus                = "SUCCESS"
	FailureApprovalStatus                = "FAILURE"
)

// status declaration
const (
	_       Status = ""
	ACTIVE         = "ACTIVE"
	PASSIVE        = "PASSIVE"
)

// member type declaration
const (
	_                              MemberType = ""
	PERSONAL                                  = "PERSONAL"
	PRIVATE_COMPANY                           = "PRIVATE_COMPANY"
	LIMITED_OR_JOINT_STOCK_COMPANY            = "LIMITED_OR_JOINT_STOCK_COMPANY"
)

// settlementEarningsDestination type declaration
const (
	_                                   SettlementEarningsDestination = ""
	SettlementEarningsDestinationIBAN                                 = "IBAN"
	SettlementEarningsDestinationWALLET                               = "WALLET"
)

// refundDestinationType type declaration
const (
	_                             RefundDestinationType = ""
	RefundDestinationTypeCARD                           = "CARD"
	RefundDestinationTypePROVIDER                       = "PROVIDER"
	RefundDestinationTypeWALLET                         = "WALLET"
)

// transaction status declaration
const (
	_                    TransactionStatus = ""
	WAITING_FOR_APPROVAL                   = "WAITING_FOR_APPROVAL"
	APPROVED                               = "APPROVED"
	PAYOUT_STARTED                         = "PAYOUT_STARTED"
)

// transaction payout status declaration
const (
	_                                         TransactionPayoutStatus = ""
	TransactionPayoutStatusCANCELLED                                  = "CANCELLED"
	TransactionPayoutStatusNO_PAYOUT                                  = "NO_PAYOUT"
	TransactionPayoutStatusWAITING_FOR_PAYOUT                         = "WAITING_FOR_PAYOUT"
	TransactionPayoutStatusPAYOUT_STARTED                             = "PAYOUT_STARTED"
	TransactionPayoutStatusPAYOUT_COMPLETED                           = "PAYOUT_COMPLETED"
)

// settlement type declaration
const (
	_                                SettlementType = ""
	SettlementTypeSETTLEMENT                        = "SETTLEMENT"
	SettlementTypeBOUNCED_SETTLEMENT                = "BOUNCED_SETTLEMENT"
	SettlementTypeWITHDRAW                          = "WITHDRAW"
)

// wallet transaction refund type declaration
const (
	_          WalletTransactionRefundCardTransactionType = ""
	PAYMENT                                               = "PAYMENT"
	PAYMENT_TX                                            = "PAYMENT_TX"
)

// fraud action type declaration
const (
	_      FraudAction = ""
	BLOCK              = "BLOCK"
	REVIEW             = "REVIEW"
)

// fraud check status type declaration
const (
	_                         FraudCheckStatus = ""
	FraudCheckStatusWAITING                    = "WAITING"
	FraudCheckStatusNOT_FRAUD                  = "NOT_FRAUD"
	FraudCheckStatusFRAUD                      = "FRAUD"
)

// apm additional action type declaration
const (
	_               ApmAdditionalAction = ""
	REDIRECT_TO_URL                     = "REDIRECT_TO_URL"
	OTP_REQUIRED                        = "OTP_REQUIRED"
	NONE                                = "NONE"
)

// report file type declaration
const (
	_    ReportFileType = ""
	CSV                 = "CSV"
	XLSX                = "XLSX"
)

// wallet transaction type declaration
const (
	_                            WalletTransactionType = ""
	PAYMENT_REDEEM                                     = "PAYMENT_REDEEM"
	REFUND_DEPOSIT                                     = "REFUND_DEPOSIT"
	REFUND_TX_DEPOSIT                                  = "REFUND_TX_DEPOSIT"
	WITHDRAW                                           = "WITHDRAW"
	CANCEL_REFUND_WALLET_TO_CARD                       = "CANCEL_REFUND_WALLET_TO_CARD"
	REFUND_WALLET_TX_TO_CARD                           = "REFUND_WALLET_TX_TO_CARD"
	CANCEL_REFUND_TO_WALLET                            = "CANCEL_REFUND_TO_WALLET"
	REFUND_TX_TO_WALLET                                = "REFUND_TX_TO_WALLET"
	MANUAL_REFUND_TX_TO_WALLET                         = "MANUAL_REFUND_TX_TO_WALLET"
	SETTLEMENT_EARNINGS                                = "SETTLEMENT_EARNINGS"
	DEPOSIT_FROM_CARD                                  = "DEPOSIT_FROM_CARD"
	REMITTANCE                                         = "REMITTANCE"
	LOYALTY                                            = "LOYALTY"
	WITHDRAW_CANCEL                                    = "WITHDRAW_CANCEL"
	MERCHANT_BALANCE_RESET                             = "MERCHANT_BALANCE_RESET"
	DEPOSIT_FROM_FUND_TRANSFER                         = "DEPOSIT_FROM_FUND_TRANSFER"
)

// requests
type CreatePaymentRequest struct {
	Price            float64                `json:"price,omitempty"`
	PaidPrice        float64                `json:"paidPrice,omitempty"`
	WalletPrice      float64                `json:"walletPrice,omitempty"`
	PosAlias         string                 `json:"posAlias,omitempty"`
	Installment      int                    `json:"installment,omitempty"`
	Currency         Currency               `json:"currency,omitempty"`
	PaymentGroup     PaymentGroup           `json:"paymentGroup,omitempty"`
	ConversationId   string                 `json:"conversationId,omitempty"`
	ExternalId       string                 `json:"externalId,omitempty"`
	ClientIp         string                 `json:"clientIp,omitempty"`
	PaymentPhase     PaymentPhase           `json:"paymentPhase,omitempty"`
	PaymentChannel   string                 `json:"paymentChannel,omitempty"`
	BuyerMemberId    int64                  `json:"buyerMemberId,omitempty"`
	BankOrderId      string                 `json:"bankOrderId,omitempty"`
	Card             Card                   `json:"card,omitempty"`
	Items            []PaymentItem          `json:"items"`
	AdditionalParams map[string]interface{} `json:"additionalParams,omitempty"`
}

type CreateApmPaymentRequest struct {
	ApmType        ApmType       `json:"apmType,omitempty"`
	Price          float64       `json:"price,omitempty"`
	PaidPrice      float64       `json:"paidPrice,omitempty"`
	Currency       Currency      `json:"currency,omitempty"`
	PaymentGroup   PaymentGroup  `json:"paymentGroup,omitempty"`
	PaymentChannel string        `json:"paymentChannel,omitempty"`
	ConversationId string        `json:"conversationId,omitempty"`
	ExternalId     string        `json:"externalId,omitempty"`
	BuyerMemberId  int64         `json:"buyerMemberId,omitempty"`
	ApmOrderId     string        `json:"apmOrderId,omitempty"`
	ClientIp       string        `json:"clientIp,omitempty"`
	Items          []PaymentItem `json:"items"`
}

type Init3DSPaymentRequest struct {
	Price            float64                `json:"price,omitempty"`
	PaidPrice        float64                `json:"paidPrice,omitempty"`
	WalletPrice      float64                `json:"walletPrice,omitempty"`
	PosAlias         string                 `json:"posAlias,omitempty"`
	Installment      int                    `json:"installment,omitempty"`
	Currency         Currency               `json:"currency,omitempty"`
	PaymentGroup     PaymentGroup           `json:"paymentGroup,omitempty"`
	ConversationId   string                 `json:"conversationId,omitempty"`
	ExternalId       string                 `json:"externalId,omitempty"`
	ClientIp         string                 `json:"clientIp,omitempty"`
	PaymentPhase     PaymentPhase           `json:"paymentPhase,omitempty"`
	PaymentChannel   string                 `json:"paymentChannel,omitempty"`
	BuyerMemberId    int64                  `json:"buyerMemberId,omitempty"`
	BankOrderId      string                 `json:"bankOrderId,omitempty"`
	Card             Card                   `json:"card,omitempty"`
	CallbackUrl      string                 `json:"callbackUrl,omitempty"`
	Items            []PaymentItem          `json:"items"`
	AdditionalParams map[string]interface{} `json:"additionalParams"`
}

type InitCheckoutPaymentRequest struct {
	Price                       float64         `json:"price,omitempty"`
	PaidPrice                   float64         `json:"paidPrice,omitempty"`
	Currency                    Currency        `json:"currency,omitempty"`
	PaymentGroup                PaymentGroup    `json:"paymentGroup,omitempty"`
	ConversationId              string          `json:"conversationId,omitempty"`
	ExternalId                  string          `json:"externalId,omitempty"`
	BankOrderId                 string          `json:"bankOrderId,omitempty"`
	CallbackUrl                 string          `json:"callbackUrl,omitempty"`
	ClientIp                    string          `json:"clientIp,omitempty"`
	PaymentPhase                PaymentPhase    `json:"paymentPhase,omitempty"`
	PaymentChannel              string          `json:"paymentChannel,omitempty"`
	EnabledPaymentMethods       []PaymentMethod `json:"enabledPaymentMethods,omitempty"`
	MasterpassGsmNumber         string          `json:"masterpassGsmNumber,omitempty"`
	MasterpassUserId            string          `json:"masterpassUserId,omitempty"`
	CardUserKey                 string          `json:"cardUserKey,omitempty"`
	BuyerMemberId               int64           `json:"buyerMemberId,omitempty"`
	EnabledInstallments         []int           `json:"enabledInstallments,omitempty"`
	AlwaysStoreCardAfterPayment bool            `json:"alwaysStoreCardAfterPayment,omitempty"`
	AllowOnlyStoredCards        bool            `json:"allowOnlyStoredCards,omitempty"`
	AllowOnlyCreditCard         bool            `json:"allowOnlyCreditCard,omitempty"`
	ForceThreeDS                bool            `json:"forceThreeDS,omitempty"`
	Items                       []PaymentItem   `json:"items"`
}

type InitApmPaymentRequest struct {
	ApmType          ApmType           `json:"apmType,omitempty"`
	Price            float64           `json:"price,omitempty"`
	PaidPrice        float64           `json:"paidPrice,omitempty"`
	Currency         Currency          `json:"currency,omitempty"`
	PaymentGroup     PaymentGroup      `json:"paymentGroup,omitempty"`
	PaymentChannel   string            `json:"paymentChannel,omitempty"`
	ConversationId   string            `json:"conversationId,omitempty"`
	ExternalId       string            `json:"externalId,omitempty"`
	CallbackUrl      string            `json:"callbackUrl,omitempty"`
	BuyerMemberId    int64             `json:"buyerMemberId,omitempty"`
	ApmOrderId       string            `json:"apmOrderId,omitempty"`
	ApmUserIdentity  string            `json:"apmUserIdentity,omitempty"`
	ClientIp         string            `json:"clientIp,omitempty"`
	AdditionalParams map[string]string `json:"additionalParams,omitempty"`
	Items            []PaymentItem     `json:"items"`
}

type CompleteApmPaymentRequest struct {
	PaymentId        int64             `json:"paymentId,omitempty"`
	AdditionalParams map[string]string `json:"additionalParams,omitempty"`
}

type Complete3DSPaymentRequest struct {
	PaymentId int64 `json:"paymentId"`
}

type PostAuthPaymentRequest struct {
	PaidPrice float64 `json:"paidPrice"`
}

type DepositPaymentRequest struct {
	Price          float64 `json:"price,omitempty"`
	BuyerMemberId  int64   `json:"buyerMemberId,omitempty"`
	ConversationId string  `json:"conversationId,omitempty"`
	CallbackUrl    string  `json:"callbackUrl,omitempty"`
	PosAlias       string  `json:"posAlias,omitempty"`
	ClientIp       string  `json:"clientIp,omitempty"`
	Card           Card    `json:"card"`
}

type CreateFundTransferDepositPaymentRequest struct {
	Price          float64 `json:"price,omitempty"`
	BuyerMemberId  int64   `json:"buyerMemberId,omitempty"`
	ConversationId string  `json:"conversationId,omitempty"`
	ClientIp       string  `json:"clientIp,omitempty"`
}

type RetrieveLoyaltiesRequest struct {
	CardNumber  string `json:"cardNumber,omitempty"`
	ExpireYear  string `json:"expireYear,omitempty"`
	ExpireMonth string `json:"expireMonth,omitempty"`
	Cvc         string `json:"cvc,omitempty"`
	CardUserKey string `json:"cardUserKey,omitempty"`
	CardToken   string `json:"cardToken,omitempty"`
}

type InitGarantiPayPaymentRequest struct {
	Price          float64                 `json:"price,omitempty"`
	PaidPrice      float64                 `json:"paidPrice,omitempty"`
	Currency       Currency                `json:"currency,omitempty"`
	PosAlias       string                  `json:"posAlias,omitempty"`
	PaymentGroup   PaymentGroup            `json:"paymentGroup,omitempty"`
	ConversationId string                  `json:"conversationId,omitempty"`
	ExternalId     string                  `json:"externalId,omitempty"`
	CallbackUrl    string                  `json:"callbackUrl,omitempty"`
	ClientIp       string                  `json:"clientIp,omitempty"`
	PaymentChannel string                  `json:"paymentChannel,omitempty"`
	BuyerMemberId  int64                   `json:"buyerMemberId,omitempty"`
	BankOrderId    string                  `json:"bankOrderId,omitempty"`
	Items          []PaymentItem           `json:"items"`
	Installments   []GarantiPayInstallment `json:"installments,omitempty"`
}

type RefundPaymentTransactionRequest struct {
	PaymentTransactionId  int64                 `json:"paymentTransactionId"`
	ConversationId        string                `json:"conversationId"`
	RefundPrice           float64               `json:"refundPrice"`
	RefundDestinationType RefundDestinationType `json:"refundDestinationType"`
	ChargeFromMe          bool                  `json:"chargeFromMe"`
}

type UpdatePaymentTransactionRequest struct {
	SubMerchantMemberId    int64   `json:"subMerchantMemberId,omitempty"`
	SubMerchantMemberPrice float64 `json:"subMerchantMemberPrice,omitempty"`
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
	CardAlias       string          `schema:"cardAlias,omitempty"`
	CardBrand       string          `schema:"cardBrand,omitempty"`
	CardType        CardType        `schema:"cardType,omitempty"`
	CardUserKey     string          `schema:"cardUserKey,omitempty"`
	CardToken       string          `schema:"cardToken,omitempty"`
	CardBankName    string          `schema:"cardBankName,omitempty"`
	CardAssociation CardAssociation `schema:"cardAssociation,omitempty"`
	Page            int             `schema:"page,omitempty"`
	Size            int             `schema:"size,omitempty"`
}

type PaymentTransactionsApprovalRequest struct {
	PaymentTransactionIds []int64 `json:"paymentTransactionIds,omitempty"`
	IsTransactional       bool    `json:"isTransactional,omitempty"`
}

type RefundPaymentRequest struct {
	PaymentId             int64                 `json:"paymentId,omitempty"`
	ConversationId        string                `json:"conversationId,omitempty"`
	RefundDestinationType RefundDestinationType `json:"refundDestinationType,omitempty"`
	ChargeFromMe          bool                  `json:"chargeFromMe,omitempty"`
}

type StoreCardRequest struct {
	CardHolderName string `json:"cardHolderName,omitempty"`
	CardNumber     string `json:"cardNumber,omitempty"`
	ExpireYear     string `json:"expireYear,omitempty"`
	ExpireMonth    string `json:"expireMonth,omitempty"`
	CardAlias      string `json:"cardAlias,omitempty"`
	CardUserKey    string `json:"cardUserKey,omitempty"`
}

type CheckMasterpassUserRequest struct {
	MasterpassGsmNumber string `json:"masterpassGsmNumber"`
}

// responses
type PaymentResponse struct {
	Id                           *int64                       `json:"id"`
	CreatedDate                  *TimeResponse                `json:"createdDate"`
	Price                        *float64                     `json:"price"`
	PaidPrice                    *float64                     `json:"paidPrice"`
	WalletPrice                  *float64                     `json:"walletPrice"`
	Currency                     *string                      `json:"currency"`
	BuyerMemberId                *int64                       `json:"buyerMemberId"`
	FraudId                      *int64                       `json:"fraudId"`
	Installment                  *int                         `json:"installment"`
	ConversationId               *string                      `json:"conversationId"`
	ExternalId                   *string                      `json:"externalId"`
	PaymentType                  *PaymentType                 `json:"paymentType"`
	PaymentProvider              *PaymentProvider             `json:"paymentProvider"`
	PaymentSource                *PaymentSource               `json:"paymentSource"`
	PaymentGroup                 *PaymentGroup                `json:"paymentGroup"`
	PaymentStatus                *PaymentStatus               `json:"paymentStatus"`
	PaymentPhase                 *PaymentPhase                `json:"paymentPhase"`
	FraudAction                  *FraudAction                 `json:"fraudAction"`
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
	Pos                          *MerchantPos                 `json:"pos"`
	Loyalty                      *Loyalty                     `json:"loyalty"`
	PaymentError                 *PaymentError                `json:"paymentError"`
	CardUserKey                  *string                      `json:"cardUserKey"`
	CardToken                    *string                      `json:"cardToken"`
	PaymentTransactions          []PaymentTransactionResponse `json:"paymentTransactions"`
	AdditionalData               map[string]any               `json:"additionalData"`
}

type PaymentTransactionResponse struct {
	ID                            *int64   `json:"id"`
	ExternalID                    *string  `json:"externalId"`
	Name                          *string  `json:"name"`
	TransactionStatus             *string  `json:"transactionStatus"`
	BlockageResolvedDate          *string  `json:"blockageResolvedDate"`
	Price                         *float64 `json:"price"`
	PaidPrice                     *float64 `json:"paidPrice"`
	WalletPrice                   *float64 `json:"walletPrice"`
	MerchantCommissionRate        *float64 `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount  *float64 `json:"merchantCommissionRateAmount"`
	MerchantPayoutAmount          *float64 `json:"merchantPayoutAmount"`
	SubMerchantMemberID           *int64   `json:"subMerchantMemberId"`
	SubMerchantMemberPrice        *float64 `json:"subMerchantMemberPrice"`
	SubMerchantMemberPayoutRate   *float64 `json:"subMerchantMemberPayoutRate"`
	SubMerchantMemberPayoutAmount *float64 `json:"subMerchantMemberPayoutAmount"`
	Payout                        *Payout  `json:"payout"`
}

type Init3DSPaymentResponse struct {
	HtmlContent *string `json:"htmlContent"`
}

type InitCheckoutPaymentResponse struct {
	Token   *string `json:"token"`
	PageUrl *string `json:"pageUrl"`
}

type InitApmPaymentResponse struct {
	PaymentId           *int64               `json:"paymentId"`
	RedirectUrl         *string              `json:"redirectUrl"`
	PaymentStatus       *PaymentStatus       `json:"paymentStatus"`
	ApmAdditionalAction *ApmAdditionalAction `json:"additionalAction"`
}

type CompleteApmPaymentResponse struct {
	PaymentId     *int64         `json:"paymentId"`
	PaymentStatus *PaymentStatus `json:"paymentStatus"`
}

type DepositPaymentResponse struct {
	Id                       *int64             `json:"id"`
	CreatedDate              *TimeResponse      `json:"createdDate"`
	Price                    *float64           `json:"price"`
	Currency                 *string            `json:"currency"`
	BuyerMemberId            *int64             `json:"buyerMemberId"`
	ConversationId           *string            `json:"conversationId"`
	BankCommissionRate       *float64           `json:"bankCommissionRate"`
	BankCommissionRateAmount *float64           `json:"bankCommissionRateAmount"`
	AuthCode                 *string            `json:"authCode"`
	HostReference            *string            `json:"hostReference"`
	TransId                  *string            `json:"transId"`
	OrderId                  *string            `json:"orderId"`
	PaymentType              *PaymentType       `json:"paymentType"`
	PaymentStatus            *PaymentStatus     `json:"paymentStatus"`
	CardUserKey              *string            `json:"cardUserKey"`
	CardToken                *string            `json:"cardToken"`
	WalletTransaction        *WalletTransaction `json:"walletTransaction"`
}

type RefundWalletTransactionToCardRequest struct {
	RefundPrice float64 `json:"refundPrice"`
}

type RemittanceRequest struct {
	MemberId             int64   `json:"memberId"`
	Price                float64 `json:"price"`
	Description          string  `json:"description"`
	RemittanceReasonType string  `json:"remittanceReasonType"`
}

type CreateWithdrawRequest struct {
	MemberId    int64    `json:"memberId"`
	Price       float64  `json:"price"`
	Description string   `json:"description"`
	Currency    Currency `json:"currency"`
}

type SearchWalletTransactionsRequest struct {
	WalletTransactionType WalletTransactionType `schema:"walletTransactionType,omitempty"`
	Page                  int                   `schema:"page,omitempty"`
	Size                  int                   `schema:"size,omitempty"`
}

type SearchWithdrawsRequest struct {
	MemberId         int64     `schema:"walletId,omitempty"`
	PayoutStatus     string    `schema:"payoutStatus,omitempty"`
	Currency         Currency  `schema:"currency,omitempty"`
	MinWithdrawPrice float64   `schema:"minWithdrawPrice,omitempty"`
	MaxWithdrawPrice float64   `schema:"maxWithdrawPrice,omitempty"`
	MinCreatedDate   time.Time `schema:"minCreatedDate,omitempty"`
	MaxCreatedDate   time.Time `schema:"maxCreatedDate,omitempty"`
	Page             int       `schema:"page,omitempty"`
	Size             int       `schema:"size,omitempty"`
}

type MemberWalletResponse struct {
	Id               *int64        `json:"id"`
	CreatedDate      *TimeResponse `json:"createdDate"`
	UpdatedDate      *TimeResponse `json:"updatedDate"`
	Amount           *float64      `json:"amount"`
	WithdrawalAmount *float64      `json:"withdrawalAmount"`
	Currency         *Currency     `json:"currency"`
	MemberId         *int64        `json:"memberId"`
}

type RemittanceResponse struct {
	Id                   *int64        `json:"id"`
	CreatedDate          *TimeResponse `json:"createdDate"`
	Status               *Status       `json:"status"`
	Price                *float64      `json:"price"`
	MemberId             *int64        `json:"memberId"`
	RemittanceType       *string       `json:"remittanceType"`
	RemittanceReasonType *string       `json:"remittanceReasonType"`
	Description          *string       `json:"description"`
}

type WithdrawResponse struct {
	Id           *int64                   `json:"id"`
	CreatedDate  *TimeResponse            `json:"createdDate"`
	Status       *Status                  `json:"status"`
	MemberId     *int64                   `json:"memberId"`
	PayoutId     *int64                   `json:"payoutId"`
	Price        *float64                 `json:"price"`
	Description  *string                  `json:"description"`
	Currency     *Currency                `json:"currency"`
	PayoutStatus *TransactionPayoutStatus `json:"payoutStatus"`
}

type RefundWalletTransactionToCardResponse struct {
	Id                  *int64                                      `json:"id"`
	CreatedDate         *TimeResponse                               `json:"createdDate"`
	RefundStatus        *string                                     `json:"refundStatus"`
	RefundPrice         *float64                                    `json:"refundPrice"`
	AuthCode            *string                                     `json:"authCode"`
	HostReference       *string                                     `json:"hostReference"`
	TransId             *string                                     `json:"transId"`
	TransactionId       *int64                                      `json:"transactionId"`
	WalletTransactionId *int64                                      `json:"walletTransactionId"`
	PaymentError        *PaymentError                               `json:"paymentError"`
	TransactionType     *WalletTransactionRefundCardTransactionType `json:"transactionType"`
}

type SearchWalletTransactionsResponse struct {
	ID                    *int64        `json:"id"`
	CreatedDate           *TimeResponse `json:"createdDate"`
	WalletTransactionType *string       `json:"walletTransactionType"`
	Amount                *float64      `json:"amount"`
	TransactionID         *int64        `json:"transactionId"`
	WalletID              *int64        `json:"walletId"`
}

type ResetMerchantMemberWalletBalanceRequest struct {
	WalletAmount float64 `json:"walletAmount"`
}

type RetrieveWalletTransactionRefundableAmountResponse struct {
	RefundableAmount *float64 `json:"refundableAmount"`
}

type FundTransferDepositPaymentResponse struct {
	Price             *float64           `json:"price"`
	Currency          *string            `json:"currency"`
	ConversationId    *string            `json:"conversationId"`
	BuyerMemberId     *int64             `json:"buyerMemberId"`
	WalletTransaction *WalletTransaction `json:"walletTransaction"`
}

type WalletTransaction struct {
	Id                    *int64                 `json:"id"`
	WalletTransactionType *WalletTransactionType `json:"walletTransactionType"`
	Amount                *float64               `json:"amount"`
	WalletId              *int64                 `json:"walletId"`
}

type GarantiPayInstallment struct {
	Number     int     `json:"number,omitempty"`
	TotalPrice float64 `json:"totalPrice,omitempty"`
}

type InitGarantiPayPaymentResponse struct {
	HtmlContent *string `json:"htmlContent"`
}

type RetrieveLoyaltiesResponse struct {
	CardBrand *string      `json:"cardBrand"`
	Force3ds  *bool        `json:"force3ds"`
	Pos       *MerchantPos `json:"pos"`
	Loyalties []Loyalty    `json:"loyalties"`
}

type PaymentTransactionRefundResponse struct {
	Id                    *int64                 `json:"id"`
	CreatedDate           *TimeResponse          `json:"createdDate"`
	Status                *RefundStatus          `json:"status"`
	RefundDestinationType *RefundDestinationType `json:"refundDestinationType"`
	RefundPrice           *float64               `json:"refundPrice"`
	RefundBankPrice       *float64               `json:"refundBankPrice"`
	RefundWalletPrice     *float64               `json:"refundWalletPrice"`
	ConversationId        *string                `json:"conversationId"`
	AuthCode              *string                `json:"authCode"`
	HostReference         *string                `json:"hostReference"`
	TransId               *string                `json:"transId"`
	IsAfterSettlement     *bool                  `json:"isAfterSettlement"`
	PaymentTransactionId  *int64                 `json:"paymentTransactionId"`
	Currency              *Currency              `json:"currency"`
	PaymentId             *int64                 `json:"paymentId"`
}

type PaymentRefundResponse struct {
	Id                        *int64                             `json:"id"`
	CreatedDate               *TimeResponse                      `json:"createdDate"`
	Status                    *RefundStatus                      `json:"status"`
	RefundDestinationType     *RefundDestinationType             `json:"refundDestinationType"`
	RefundPrice               *float64                           `json:"refundPrice"`
	RefundBankPrice           *float64                           `json:"refundBankPrice"`
	RefundWalletPrice         *float64                           `json:"refundWalletPrice"`
	ConversationId            *string                            `json:"conversationId"`
	AuthCode                  *string                            `json:"authCode"`
	HostReference             *string                            `json:"hostReference"`
	TransId                   *string                            `json:"transId"`
	PaymentId                 *int64                             `json:"paymentId"`
	RefundType                *RefundType                        `json:"refundType"`
	Currency                  *Currency                          `json:"currency"`
	PaymentTransactionRefunds []PaymentTransactionRefundResponse `json:"paymentTransactionRefunds"`
}

type StoredCardResponse struct {
	BinNumber        *string           `json:"binNumber"`
	LastFourDigits   *string           `json:"lastFourDigits"`
	CardUserKey      *string           `json:"cardUserKey"`
	CardToken        *string           `json:"cardToken"`
	CardAlias        *string           `json:"cardAlias"`
	CardType         *CardType         `json:"cardType"`
	CardAssociation  *CardAssociation  `json:"cardAssociation"`
	CardExpiryStatus *CardExpiryStatus `json:"cardExpiryStatus"`
	CardBrand        *string           `json:"cardBrand"`
	CardBankName     *string           `json:"cardBankName"`
	CardBankId       *int64            `json:"cardBankId"`
}

type PaymentTransactionsApprovalResponse struct {
	PaymentTransactionId *int64          `json:"paymentTransactionId"`
	ApprovalStatus       *ApprovalStatus `json:"approvalStatus"`
	FailedReason         *string         `json:"failedReason"`
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

type InstallmentPrice struct {
	InstallmentPrice  *float64 `json:"installmentPrice"`
	TotalPrice        *float64 `json:"totalPrice"`
	InstallmentNumber *int     `json:"installmentNumber"`
	InstallmentLabel  *string  `json:"installmentLabel"`
}

type SearchInstallmentsRequest struct {
	BinNumber                               string   `schema:"binNumber"`
	Price                                   float64  `schema:"price"`
	Currency                                Currency `schema:"currency"`
	DistinctCardBrandsWithLowestCommissions bool     `schema:"distinctCardBrandsWithLowestCommissions"`
}

type InstallmentListResponse struct {
	Items []InstallmentResponse `json:"items"`
}

type InstallmentResponse struct {
	BinNumber         *string            `json:"binNumber"`
	Price             *float64           `json:"price"`
	CardType          *string            `json:"cardType"`
	CardAssociation   *string            `json:"cardAssociation"`
	CardBrand         *string            `json:"cardBrand"`
	BankName          *string            `json:"bankName"`
	BankCode          *int               `json:"bankCode"`
	Force3Ds          *bool              `json:"force3ds"`
	CvcRequired       *bool              `json:"cvcRequired"`
	Commercial        *bool              `json:"commercial"`
	PosAlias          *string            `json:"posAlias"`
	InstallmentPrices []InstallmentPrice `json:"installmentPrices"`
}

type RetrieveBinNumberResponse struct {
	BinNumber       *string `json:"binNumber"`
	CardType        *string `json:"cardType"`
	CardAssociation *string `json:"cardAssociation"`
	CardBrand       *string `json:"cardBrand"`
	BankName        *string `json:"bankName"`
	BankCode        *int    `json:"bankCode"`
	Commercial      *bool   `json:"commercial"`
}

type CreateMemberRequest struct {
	MemberExternalId              string                        `json:"memberExternalId,omitempty"`
	Name                          string                        `json:"name,omitempty"`
	Address                       string                        `json:"address,omitempty"`
	Email                         string                        `json:"email,omitempty"`
	PhoneNumber                   string                        `json:"phoneNumber,omitempty"`
	IdentityNumber                string                        `json:"identityNumber,omitempty"`
	ContactName                   string                        `json:"contactName,omitempty"`
	ContactSurname                string                        `json:"contactSurname,omitempty"`
	MemberType                    MemberType                    `json:"memberType,omitempty"`
	LegalCompanyTitle             string                        `json:"legalCompanyTitle,omitempty"`
	TaxOffice                     string                        `json:"taxOffice,omitempty"`
	TaxNumber                     string                        `json:"taxNumber,omitempty"`
	Iban                          string                        `json:"iban,omitempty"`
	SettlementEarningsDestination SettlementEarningsDestination `json:"settlementEarningsDestination,omitempty"`
	NegativeWalletAmountLimit     float64                       `json:"negativeWalletAmountLimit,omitempty"`
	IsBuyer                       bool                          `json:"isBuyer,omitempty"`
	IsSubMerchant                 bool                          `json:"isSubMerchant,omitempty"`
}

type UpdateMemberRequest struct {
	Name                          string                        `json:"name,omitempty"`
	Address                       string                        `json:"address,omitempty"`
	Email                         string                        `json:"email,omitempty"`
	PhoneNumber                   string                        `json:"phoneNumber,omitempty"`
	IdentityNumber                string                        `json:"identityNumber,omitempty"`
	ContactName                   string                        `json:"contactName,omitempty"`
	ContactSurname                string                        `json:"contactSurname,omitempty"`
	MemberType                    MemberType                    `json:"memberType,omitempty"`
	LegalCompanyTitle             string                        `json:"legalCompanyTitle,omitempty"`
	TaxOffice                     string                        `json:"taxOffice,omitempty"`
	TaxNumber                     string                        `json:"taxNumber,omitempty"`
	Iban                          string                        `json:"iban,omitempty"`
	SettlementEarningsDestination SettlementEarningsDestination `json:"settlementEarningsDestination,omitempty"`
	NegativeWalletAmountLimit     float64                       `json:"negativeWalletAmountLimit,omitempty"`
	IsBuyer                       bool                          `json:"isBuyer,omitempty"`
	IsSubMerchant                 bool                          `json:"isSubMerchant,omitempty"`
}

type SearchMembersRequest struct {
	Page             int        `schema:"page,omitempty"`
	Size             int        `schema:"size,omitempty"`
	IsBuyer          bool       `schema:"isBuyer,omitempty"`
	IsSubMerchant    bool       `schema:"isSubMerchant,omitempty"`
	Name             string     `schema:"name,omitempty"`
	MemberIds        []int64    `schema:"memberIds,omitempty"`
	MemberType       MemberType `schema:"memberType,omitempty"`
	MemberExternalId string     `schema:"memberExternalId,omitempty"`
}

type MemberResponse struct {
	Id                            *int64                         `json:"id"`
	CreatedDate                   *TimeResponse                  `json:"createdDate"`
	UpdatedDate                   *TimeResponse                  `json:"updatedDate"`
	Status                        *Status                        `json:"status"`
	IsBuyer                       *bool                          `json:"isBuyer"`
	IsSubMerchant                 *bool                          `json:"isSubMerchant"`
	MemberType                    *MemberType                    `json:"memberType"`
	MemberExternalId              *string                        `json:"memberExternalId"`
	Name                          *string                        `json:"name"`
	Email                         *string                        `json:"email"`
	Address                       *string                        `json:"address"`
	PhoneNumber                   *string                        `json:"phoneNumber"`
	IdentityNumber                *string                        `json:"identityNumber"`
	ContactName                   *string                        `json:"contactName"`
	ContactSurname                *string                        `json:"contactSurname"`
	LegalCompanyTitle             *string                        `json:"legalCompanyTitle"`
	TaxOffice                     *string                        `json:"taxOffice"`
	TaxNumber                     *string                        `json:"taxNumber"`
	SettlementEarningsDestination *SettlementEarningsDestination `json:"settlementEarningsDestination"`
	NegativeWalletAmountLimit     *float64                       `json:"negativeWalletAmountLimit"`
	Iban                          *string                        `json:"iban"`
}

type CreateProductRequest struct {
	Name                string   `json:"name"`
	Channel             string   `json:"channel,omitempty"`
	OrderId             string   `json:"orderId,omitempty"`
	Stock               int      `json:"stock,omitempty"`
	Price               float64  `json:"price"`
	Currency            Currency `json:"currency"`
	Description         string   `json:"description,omitempty"`
	EnabledInstallments []int    `json:"enabledInstallments"`
}

type UpdateProductRequest struct {
	Name                string   `json:"name"`
	Channel             string   `json:"channel,omitempty"`
	OrderId             string   `json:"orderId,omitempty"`
	Stock               int      `json:"stock,omitempty"`
	Status              Status   `json:"status"`
	Price               float64  `json:"price"`
	Currency            Currency `json:"currency"`
	Description         string   `json:"description,omitempty"`
	EnabledInstallments []int    `json:"enabledInstallments"`
}

type SearchProductsRequest struct {
	Name     string   `schema:"name,omitempty"`
	MinPrice float64  `schema:"minPrice,omitempty"`
	MaxPrice float64  `schema:"maxPrice,omitempty"`
	Currency Currency `schema:"currency,omitempty"`
	Channel  string   `schema:"channel,omitempty"`
	Page     int      `schema:"page,omitempty"`
	Size     int      `schema:"size,omitempty"`
}

type ProductResponse struct {
	Id                  *int64    `json:"id"`
	Name                *string   `json:"name"`
	Description         *string   `json:"description"`
	OrderId             *string   `json:"orderId,omitempty"`
	Status              *Status   `json:"status"`
	Price               *float64  `json:"price"`
	Currency            *Currency `json:"currency"`
	Stock               *int      `json:"stock"`
	SoldCount           *int      `json:"soldCount"`
	Token               *string   `json:"token"`
	EnabledInstallments []int     `json:"enabledInstallments"`
	Url                 *string   `json:"url"`
	QrCodeUrl           *string   `json:"qrCodeUrl"`
	Channel             *string   `json:"channel"`
}

type SearchPaymentsRequest struct {
	Page                 int             `schema:"page,omitempty"`
	Size                 int             `schema:"size,omitempty"`
	PaymentId            int64           `schema:"paymentId,omitempty"`
	PaymentTransactionId int64           `schema:"paymentTransactionId,omitempty"`
	BuyerMemberId        int64           `schema:"buyerMemberId,omitempty"`
	SubMerchantMemberId  int64           `schema:"subMerchantMemberId,omitempty"`
	ConversationId       string          `schema:"conversationId,omitempty"`
	ExternalId           string          `schema:"externalId,omitempty"`
	OrderId              string          `schema:"orderId,omitempty"`
	PaymentType          PaymentType     `schema:"paymentType,omitempty"`
	PaymentProvider      PaymentProvider `schema:"paymentProvider,omitempty"`
	PaymentStatus        PaymentStatus   `schema:"paymentStatus,omitempty"`
	PaymentSource        PaymentSource   `schema:"paymentSource,omitempty"`
	PaymentChannel       string          `schema:"paymentChannel,omitempty"`
	BinNumber            string          `schema:"binNumber,omitempty"`
	LastFourDigits       string          `schema:"lastFourDigits,omitempty"`
	Currency             Currency        `schema:"currency,omitempty"`
	MinPaidPrice         float64         `schema:"minPaidPrice,omitempty"`
	MaxPaidPrice         float64         `schema:"maxPaidPrice,omitempty"`
	Installment          int             `schema:"installment,omitempty"`
	IsThreeDS            bool            `schema:"isThreeDS,omitempty"`
	MinCreatedDate       time.Time       `schema:"minCreatedDate,omitempty"`
	MaxCreatedDate       time.Time       `schema:"maxCreatedDate,omitempty"`
}

type SearchPaymentRefundsRequest struct {
	Page           int          `schema:"page,omitempty"`
	Size           int          `schema:"size,omitempty"`
	Id             int64        `schema:"id,omitempty"`
	PaymentId      int64        `schema:"paymentId,omitempty"`
	BuyerMemberId  int64        `schema:"buyerMemberId,omitempty"`
	ConversationId string       `schema:"conversationId,omitempty"`
	Status         RefundStatus `schema:"status,omitempty"`
	Currency       Currency     `schema:"currency,omitempty"`
	MinRefundPrice float64      `schema:"minRefundPrice,omitempty"`
	MaxRefundPrice float64      `schema:"maxRefundPrice,omitempty"`
	MinCreatedDate time.Time    `schema:"minCreatedDate,omitempty"`
	MaxCreatedDate time.Time    `schema:"maxCreatedDate,omitempty"`
}

type SearchPaymentTransactionRefundsRequest struct {
	Page                 int          `schema:"page,omitempty"`
	Size                 int          `schema:"size,omitempty"`
	Id                   int64        `schema:"id,omitempty"`
	PaymentId            int64        `schema:"paymentId,omitempty"`
	PaymentTransactionId int64        `schema:"paymentTransactionId,omitempty"`
	BuyerMemberId        int64        `schema:"buyerMemberId,omitempty"`
	ConversationId       string       `schema:"conversationId,omitempty"`
	Status               RefundStatus `schema:"status,omitempty"`
	Currency             Currency     `schema:"currency,omitempty"`
	IsAfterSettlement    bool         `schema:"isAfterSettlement,omitempty"`
	MinRefundPrice       float64      `schema:"minRefundPrice,omitempty"`
	MaxRefundPrice       float64      `schema:"maxRefundPrice,omitempty"`
	MinCreatedDate       time.Time    `schema:"minCreatedDate,omitempty"`
	MaxCreatedDate       time.Time    `schema:"maxCreatedDate,omitempty"`
}

type ReportingPaymentResponse struct {
	Id                           *int64                            `json:"id"`
	CreatedDate                  *TimeResponse                     `json:"createdDate"`
	Price                        *float64                          `json:"price"`
	PaidPrice                    *float64                          `json:"paidPrice"`
	WalletPrice                  *float64                          `json:"walletPrice"`
	Currency                     *Currency                         `json:"currency"`
	BuyerMemberId                *int64                            `json:"buyerMemberId"`
	Installment                  *int                              `json:"installment"`
	ConversationId               *string                           `json:"conversationId"`
	ExternalId                   *string                           `json:"externalId"`
	PaymentType                  *PaymentType                      `json:"paymentType"`
	PaymentProvider              *PaymentProvider                  `json:"paymentProvider"`
	PaymentSource                *PaymentSource                    `json:"paymentSource"`
	PaymentGroup                 *PaymentGroup                     `json:"paymentGroup"`
	PaymentStatus                *PaymentStatus                    `json:"paymentStatus"`
	PaymentPhase                 *PaymentPhase                     `json:"paymentPhase"`
	PaymentChannel               *string                           `json:"paymentChannel"`
	IsThreeDS                    *bool                             `json:"isThreeDS"`
	MerchantCommissionRate       *float64                          `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount *float64                          `json:"merchantCommissionRateAmount"`
	BankCommissionRate           *float64                          `json:"bankCommissionRate"`
	BankCommissionRateAmount     *float64                          `json:"bankCommissionRateAmount"`
	PaidWithStoredCard           *bool                             `json:"paidWithStoredCard"`
	BinNumber                    *string                           `json:"binNumber"`
	LastFourDigits               *string                           `json:"lastFourDigits"`
	AuthCode                     *string                           `json:"authCode"`
	HostReference                *string                           `json:"hostReference"`
	OrderId                      *string                           `json:"orderId"`
	TransId                      *string                           `json:"transId"`
	CardHolderName               *string                           `json:"cardHolderName"`
	BankCardHolderName           *string                           `json:"bankCardHolderName"`
	CardType                     *CardType                         `json:"cardType"`
	CardAssociation              *CardAssociation                  `json:"cardAssociation"`
	CardBrand                    *string                           `json:"cardBrand"`
	RequestedPosAlias            *string                           `json:"requestedPosAlias"`
	RetryCount                   *int                              `json:"retryCount"`
	RefundablePrice              *float64                          `json:"refundablePrice"`
	RefundStatus                 *PaymentRefundStatus              `json:"refundStatus"`
	CardIssuerBankName           *string                           `json:"cardIssuerBankName"`
	MdStatus                     *int                              `json:"mdStatus"`
	BuyerMember                  *MemberResponse                   `json:"buyerMember"`
	Refunds                      *[]ReportingPaymentRefundResponse `json:"refunds"`
	Pos                          *MerchantPos                      `json:"pos"`
	Loyalty                      *Loyalty                          `json:"loyalty"`
	PaymentError                 *PaymentError                     `json:"paymentError"`
}

type ReportingPaymentRefundResponse struct {
	Id                    *int64                 `json:"id"`
	CreatedDate           *TimeResponse          `json:"createdDate"`
	Status                *RefundStatus          `json:"status"`
	RefundDestinationType *RefundDestinationType `json:"refundDestinationType"`
	RefundPrice           *float64               `json:"refundPrice"`
	RefundBankPrice       *float64               `json:"refundBankPrice"`
	RefundWalletPrice     *float64               `json:"refundWalletPrice"`
	ConversationId        *string                `json:"conversationId"`
	AuthCode              *string                `json:"authCode"`
	HostReference         *string                `json:"hostReference"`
	PaymentType           *PaymentType           `json:"paymentType"`
	TransId               *string                `json:"transId"`
	PaymentId             *int64                 `json:"paymentId"`
	PaymentError          *PaymentError          `json:"paymentError"`
}

type ReportingPaymentTransactionRefundResponse struct {
	Id                    *int64                 `json:"id"`
	CreatedDate           *TimeResponse          `json:"createdDate"`
	Status                *RefundStatus          `json:"status"`
	RefundDestinationType *RefundDestinationType `json:"refundDestinationType"`
	RefundPrice           *float64               `json:"refundPrice"`
	RefundBankPrice       *float64               `json:"refundBankPrice"`
	RefundWalletPrice     *float64               `json:"refundWalletPrice"`
	ConversationId        *string                `json:"conversationId"`
	AuthCode              *string                `json:"authCode"`
	HostReference         *string                `json:"hostReference"`
	TransId               *string                `json:"transId"`
	IsAfterSettlement     *bool                  `json:"isAfterSettlement"`
	PaymentTransactionId  *int64                 `json:"paymentTransactionId"`
	PaymentError          *PaymentError          `json:"paymentError"`
}

type ReportingPaymentTransactionResponse struct {
	Id                            *int64               `json:"id"`
	Name                          *string              `json:"name"`
	ExternalId                    *string              `json:"externalId"`
	Price                         *float64             `json:"price"`
	PaidPrice                     *float64             `json:"paidPrice"`
	WalletPrice                   *float64             `json:"walletPrice"`
	MerchantCommissionRate        *float64             `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount  *float64             `json:"merchantCommissionRateAmount"`
	MerchantPayoutAmount          *float64             `json:"merchantPayoutAmount"`
	SubMerchantMemberId           *int64               `json:"subMerchantMemberId"`
	SubMerchantMemberPrice        *float64             `json:"subMerchantMemberPrice"`
	SubMerchantMemberPayoutRate   *float64             `json:"subMerchantMemberPayoutRate"`
	SubMerchantMemberPayoutAmount *float64             `json:"subMerchantMemberPayoutAmount"`
	TransactionStatus             *TransactionStatus   `json:"transactionStatus"`
	BlockageResolvedDate          *TimeResponse        `json:"blockageResolvedDate"`
	CreatedDate                   *TimeResponse        `json:"createdDate"`
	TransactionStatusDate         *TimeResponse        `json:"transactionStatusDate"`
	RefundablePrice               *float64             `json:"refundablePrice"`
	BankCommissionRate            *float64             `json:"bankCommissionRate"`
	BankCommissionRateAmount      *float64             `json:"bankCommissionRateAmount"`
	Payout                        *Payout              `json:"payout"`
	SubMerchantMember             *MemberResponse      `json:"subMerchantMember"`
	RefundStatus                  *PaymentRefundStatus `json:"refundStatus"`
	PayoutStatus                  *PayoutStatus        `json:"payoutStatus"`
}

type Payout struct {
	PaidPrice                     *float64  `json:"paidPrice"`
	Parity                        *float64  `json:"parity"`
	Currency                      *Currency `json:"currency"`
	MerchantPayoutAmount          *float64  `json:"merchantPayoutAmount"`
	SubMerchantMemberPayoutAmount *float64  `json:"subMerchantMemberPayoutAmount"`
}

type PayoutStatus struct {
	MerchantStatus              *TransactionPayoutStatus `json:"merchantStatus"`
	MerchantStatusDate          *TimeResponse            `json:"merchantStatusDate"`
	SubMerchantMemberStatus     *TransactionPayoutStatus `json:"subMerchantMemberStatus"`
	SubMerchantMemberStatusDate *TimeResponse            `json:"subMerchantMemberStatusDate"`
}

type CreateInstantWalletSettlementRequest struct {
	ExcludedSubMerchantMemberIds []int64
}

type CreateInstantWalletSettlementResponse struct {
	SettlementResultStatus *string `json:"settlementResultStatus"`
}

type SearchPayoutCompletedTransactionsRequest struct {
	SettlementFileId int64          `schema:"settlementFileId,omitempty"`
	SettlementType   SettlementType `schema:"settlementType,omitempty"`
	StartDate        time.Time      `schema:"startDate,omitempty"`
	EndDate          time.Time      `schema:"endDate,omitempty"`
}

type SearchPayoutBouncedTransactionsRequest struct {
	StartDate time.Time `schema:"startDate,omitempty"`
	EndDate   time.Time `schema:"endDate,omitempty"`
}

type RetrievePayoutDetailsRequest struct {
	PayoutDetailId int64
}

type SearchPayoutCompletedTransactionsResponse struct {
	PayoutId                      *int64    `json:"payoutId"`
	TransactionId                 *int64    `json:"transactionId"`
	TransactionType               *string   `json:"transactionType"`
	PayoutAmount                  *float64  `json:"payoutAmount"`
	Currency                      *Currency `json:"currency"`
	MerchantId                    *int64    `json:"merchantId"`
	MerchantType                  *string   `json:"merchantType"`
	SettlementEarningsDestination *string   `json:"settlementEarningsDestination"`
	SettlementSource              *string   `json:"settlementSource"`
}

type SearchPayoutBouncedTransactionsResponse struct {
	Id                *int64        `json:"id"`
	Iban              *string       `json:"iban"`
	CreatedDate       *TimeResponse `json:"createdDate"`
	UpdatedDate       *TimeResponse `json:"updatedDate"`
	PayoutId          *int64        `json:"payoutId"`
	PayoutAmount      *float64      `json:"payoutAmount"`
	ContactName       *string       `json:"contactName"`
	ContactSurname    *string       `json:"contactSurname"`
	LegalCompanyTitle *string       `json:"legalCompanyTitle"`
	RowDescription    *string       `json:"rowDescription"`
}

type PayoutDetailResponse struct {
	RowDescription                *string                           `json:"rowDescription"`
	PayoutDate                    *TimeResponse                     `json:"payoutDate"`
	Name                          *string                           `json:"name"`
	Iban                          *string                           `json:"iban"`
	PayoutAmount                  *float64                          `json:"payoutAmount"`
	Currency                      *string                           `json:"currency"`
	MerchantId                    *int64                            `json:"merchantId"`
	MerchantType                  *string                           `json:"merchantType"`
	SettlementEarningsDestination *string                           `json:"settlementEarningsDestination"`
	SettlementSource              *string                           `json:"settlementSource"`
	BounceStatus                  *string                           `json:"bounceStatus"`
	PayoutTransactions            []PayoutDetailTransactionResponse `json:"payoutTransactions"`
}

type PayoutDetailTransactionResponse struct {
	TransactionId   *int64   `json:"transactionId"`
	TransactionType *string  `json:"transactionType"`
	PayoutAmount    *float64 `json:"payoutAmount"`
}

type RetrieveDailyTransactionReportRequest struct {
	ReportDate Date           `schema:"reportDate,omitempty"`
	FileType   ReportFileType `schema:"fileType,omitempty"`
}

type SearchFraudChecksRequest struct {
	Page           int              `schema:"page,omitempty"`
	Size           int              `schema:"size,omitempty"`
	Action         FraudAction      `schema:"action,omitempty"`
	CheckStatus    FraudCheckStatus `schema:"checkStatus,omitempty"`
	MinCreatedDate time.Time        `schema:"minCreatedDate,omitempty"`
	MaxCreatedDate time.Time        `schema:"maxCreatedDate,omitempty"`
	RuleId         int64            `schema:"ruleId,omitempty"`
	PaymentId      int64            `schema:"paymentId,omitempty"`
	PaymentStatus  PaymentStatus    `schema:"paymentStatus,omitempty"`
}

type FraudCheckResponse struct {
	Id             *int64            `json:"id"`
	Status         *Status           `json:"status"`
	Action         *FraudAction      `json:"action"`
	CheckStatus    *FraudCheckStatus `json:"checkStatus"`
	PaymentData    *FraudPaymentData `json:"paymentData"`
	RuleId         *int64            `json:"ruleId"`
	RuleName       *string           `json:"ruleName"`
	RuleConditions *string           `json:"ruleConditions"`
	PaymentId      *int64            `json:"paymentId"`
	PaymentStatus  *PaymentStatus    `json:"paymentStatus"`
}

type FraudPaymentData struct {
	PaymentDate                   *time.Time `json:"paymentDate"`
	ConversationId                *string    `json:"conversationId"`
	PaidPrice                     *float64   `json:"paidPrice"`
	Currency                      *Currency  `json:"currency"`
	CardFingerprintId             *string    `json:"cardFingerprintId"`
	CardFingerprintExpirationDate *time.Time `json:"cardFingerprintExpirationDate"`
	BuyerId                       *int64     `json:"buyerId"`
	ClientIp                      *string    `json:"clientIp"`
}

type FraudValueListRequest struct {
	ListName          string `json:"listName,omitempty"`
	Value             string `json:"value,omitempty"`
	DurationInSeconds int    `json:"durationInSeconds,omitempty"`
}

type FraudValuesResponse struct {
	Name   string       `json:"name"`
	Values []FraudValue `json:"values"`
}

type FraudValue struct {
	Value           *string `json:"value"`
	ExpireInSeconds *int    `json:"expireInSeconds"`
}

type RequestOptions struct {
	BaseURL   string
	ApiKey    string
	SecretKey string
}

type Response[T any] struct {
	Data   *T             `json:"data"`
	Errors *ErrorResponse `json:"errors"`
}

func (r Response[ErrorResponse]) Error() string {
	if r.Errors.ErrorGroup != nil {
		return *r.Errors.ErrorGroup + "-" + *r.Errors.ErrorCode + "-" + *r.Errors.ErrorDescription
	}

	return *r.Errors.ErrorCode + "-" + *r.Errors.ErrorDescription
}

type ErrorResponse struct {
	ErrorGroup       *string `json:"errorGroup"`
	ErrorDescription *string `json:"errorDescription"`
	ErrorCode        *string `json:"errorCode"`
}

type DataResponse[T any] struct {
	Items     []T   `json:"items"`
	Page      int   `json:"page"`
	Size      int   `json:"size"`
	TotalSize int64 `json:"totalSize"`
}

type MerchantPos struct {
	Id     *int64  `json:"id"`
	Name   *string `json:"name"`
	Alias  *string `json:"alias"`
	BankId *int64  `json:"bankId"`
}

type Reward struct {
	CardRewardMoney *float64 `json:"cardRewardMoney,omitempty"`
	FirmRewardMoney *float64 `json:"firmRewardMoney,omitempty"`
}

type Loyalty struct {
	LoyaltyType *LoyaltyType `json:"type,omitempty"`
	Reward      *Reward      `json:"reward,omitempty"`
	Message     *string      `json:"message,omitempty"`
}

type Card struct {
	CardHolderName               string   `json:"cardHolderName,omitempty"`
	CardNumber                   string   `json:"cardNumber,omitempty"`
	ExpireYear                   string   `json:"expireYear,omitempty"`
	ExpireMonth                  string   `json:"expireMonth,omitempty"`
	Cvc                          string   `json:"cvc,omitempty"`
	CardAlias                    string   `json:"cardAlias,omitempty"`
	CardUserKey                  string   `json:"cardUserKey,omitempty"`
	CardToken                    string   `json:"cardToken,omitempty"`
	BinNumber                    string   `json:"binNumber,omitempty"`
	LastFourDigits               string   `json:"lastFourDigits,omitempty"`
	CardHolderIdentityNumber     string   `json:"cardHolderIdentityNumber,omitempty"`
	Loyalty                      *Loyalty `json:"loyalty,omitempty"`
	StoreCardAfterSuccessPayment bool     `json:"storeCardAfterSuccessPayment,omitempty"`
}

type PaymentItem struct {
	Name                   string  `json:"name,omitempty"`
	Price                  float64 `json:"price,omitempty"`
	ExternalId             string  `json:"externalId,omitempty"`
	SubMerchantMemberId    int64   `json:"subMerchantMemberId,omitempty"`
	SubMerchantMemberPrice float64 `json:"subMerchantMemberPrice,omitempty"`
}

type PaymentError ErrorResponse

type Void struct {
}
