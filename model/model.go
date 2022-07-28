package model

type PaymentType string
type PaymentProvider string
type PaymentStatus string
type PaymentSource string
type PaymentGroup string
type PaymentPhase string
type CardType string
type CardAssociation string
type Currency string
type LoyaltyType string
type PaymentRefundStatus string
type RefundStatus string
type Status string
type MemberType string
type SettlementEarningsDestination string
type RefundDestinationType string

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
)

// payment provider declaration
const (
	_           PaymentProvider = ""
	BANK                        = "BANK"
	CG_WALLET                   = "CG_WALLET"
	MASTERPASS                  = "MASTERPASS"
	GARANTI_PAY                 = "GARANTI_PAY"
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

type RequestOptions struct {
	BaseURL   string
	ApiKey    string
	SecretKey string
}

type Response[T any] struct {
	Response T             `json:"data"`
	Errors   ErrorResponse `json:"errors"`
}

type ErrorResponse struct {
	ErrorGroup       string `json:"errorGroup"`
	ErrorDescription string `json:"errorDescription"`
	ErrorCode        string `json:"errorCode"`
}

type DataResponse[T any] struct {
	Items     []T   `json:"items"`
	Page      int   `json:"page"`
	Size      int   `json:"size"`
	TotalSize int64 `json:"totalSize"`
}

type MerchantPos struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Alias  string `json:"alias"`
	BankId int64  `json:"bankId"`
}

type Reward struct {
	CardRewardMoney float64 `json:"cardRewardMoney"`
	FirmRewardMoney float64 `json:"firmRewardMoney"`
}

type Loyalty struct {
	LoyaltyType LoyaltyType `json:"type"`
	Reward      Reward      `json:"reward"`
}

type PaymentError struct {
	ErrorGroup       string `json:"errorGroup"`
	ErrorDescription string `json:"errorDescription"`
	ErrorCode        string `json:"errorCode"`
}

type MemberResponse struct {
	Id int64 `json:"id"`
	//CreatedDate                   Time                          `json:"createdDate"`
	//UpdatedDate                   Time                          `json:"updatedDate"`
	Status                        Status                        `json:"status"`
	IsBuyer                       bool                          `json:"isBuyer"`
	IsSubMerchant                 bool                          `json:"isSubMerchant"`
	MemberType                    MemberType                    `json:"memberType"`
	MemberExternalId              string                        `json:"memberExternalId"`
	Name                          string                        `json:"name"`
	Email                         string                        `json:"email"`
	Address                       string                        `json:"address"`
	PhoneNumber                   string                        `json:"phoneNumber"`
	IdentityNumber                string                        `json:"identityNumber"`
	ContactName                   string                        `json:"contactName"`
	ContactSurname                string                        `json:"contactSurname"`
	LegalCompanyTitle             string                        `json:"legalCompanyTitle"`
	TaxOffice                     string                        `json:"taxOffice"`
	TaxNumber                     string                        `json:"taxNumber"`
	SettlementEarningsDestination SettlementEarningsDestination `json:"settlementEarningsDestination"`
	NegativeWalletAmountLimit     float64                       `json:"negativeWalletAmountLimit"`
	Iban                          string                        `json:"iban"`
}
