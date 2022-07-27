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
