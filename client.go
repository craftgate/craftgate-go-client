package main

const (
	BaseURL                 = "https://sandbox-api.craftgate.io"
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
	InstallmentApi InstallmentApi
}

func CraftgateClient(apiKey, secretKey string) *Client {

	return &Client{
		InstallmentApi: InstallmentApi{
			opts: RequestOptions{
				apiKey:    apiKey,
				secretKey: secretKey,
				baseURL:   BaseURL,
			},
		},
	}
}
