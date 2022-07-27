package model

type PaymentType string
type PaymentProvider string
type PaymentStatus string
type PaymentSource string
type Currency string

type RequestOptions struct {
	BaseURL   string
	ApiKey    string
	SecretKey string
}

type Response[T any] struct {
	DataResponse DataResponse[T] `json:"data"`
	Errors       ErrorResponse   `json:"errors"`
}

type ErrorResponse struct {
	ErrorGroup       string `json:"errorGroup"`
	ErrorDescription string `json:"errorDescription"`
	ErrorCode        string `json:"errorCode"`
}

type DataResponse[T any] struct {
	Items []T `json:"items"`
}

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
	CARD_PAYMENT            PaymentType = "CARD_PAYMENT"
	DEPOSIT_PAYMENT         PaymentType = "DEPOSIT_PAYMENT"
	WALLET_PAYMENT          PaymentType = "WALLET_PAYMENT"
	CARD_AND_WALLET_PAYMENT PaymentType = "CARD_AND_WALLET_PAYMENT"
	BANK_TRANSFER           PaymentType = "BANK_TRANSFER"
)

// payment provider declaration
const (
	BANK        PaymentProvider = "BANK"
	CG_WALLET   PaymentProvider = "CG_WALLET"
	MASTERPASS  PaymentProvider = "MASTERPASS"
	GARANTI_PAY PaymentProvider = "GARANTI_PAY"
)

// payment status declaration
const (
	FAILURE          PaymentStatus = "FAILURE"
	SUCCESS          PaymentStatus = "SUCCESS"
	INIT_THREEDS     PaymentStatus = "INIT_THREEDS"
	CALLBACK_THREEDS PaymentStatus = "CALLBACK_THREEDS"
	WAITING          PaymentStatus = "WAITING"
)

// payment source declaration
const (
	API           PaymentSource = "API"
	CHECKOUT_FORM PaymentSource = "CHECKOUT_FORM"
	PAY_BY_LINK   PaymentSource = "PAY_BY_LINK"
)

// currency declaration
const (
	TRY Currency = "TRY"
	USD Currency = "USD"
	EUR Currency = "EUR"
	GBP Currency = "GBP"
)
