package main

const (
	ApiKeyHeaderName        = "x-api-key"
	RandomHeaderName        = "x-rnd-key"
	AuthVersionHeaderName   = "x-auth-version"
	ClientVersionHeaderName = "x-client-version"
	SignatureHeaderName     = "x-signature"
)

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

type Client struct {
	Installment InstallmentApi
}

func CraftgateClient(apiKey, secretKey, baseURL string) *Client {

	return &Client{
		Installment: InstallmentApi{
			Opts: RequestOptions{
				ApiKey:    apiKey,
				SecretKey: secretKey,
				BaseURL:   baseURL,
			},
		},
	}
}
