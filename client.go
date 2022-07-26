package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	BaseURL                 = "https://sandbox-api.craftgate.io"
	ApiKeyHeaderName        = "x-api-key"
	RandomHeaderName        = "x-rnd-key"
	AuthVersionHeaderName   = "x-auth-version"
	ClientVersionHeaderName = "x-client-version"
	SignatureHeaderName     = "x-signature"
)

type ResponseData map[string]interface{}

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
	BaseURL    string
	apiKey     string
	secretKey  string
	HTTPClient *http.Client
}

func CraftgateClient(apiKey, secretKey string) *Client {
	return &Client{
		BaseURL:   BaseURL,
		apiKey:    apiKey,
		secretKey: secretKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	randomStr := GenerateRandomString()
	hashStr := GenerateHash(req.URL.String(), c.apiKey, c.secretKey, randomStr, "")

	req.Header.Set(ApiKeyHeaderName, c.apiKey)
	req.Header.Set(RandomHeaderName, randomStr)
	req.Header.Set(AuthVersionHeaderName, "1")
	req.Header.Set(ClientVersionHeaderName, "craftgate-go-client:1.0.0")
	req.Header.Set(SignatureHeaderName, hashStr)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes Response[any]
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Errors.ErrorGroup)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}

func GenerateHash(url, apiKey, secretKey, randomString, body string) string {
	hashStr := []string{url, apiKey, secretKey, randomString, body}
	hash := strings.Join(hashStr, "")

	fmt.Println(hash)
	hasher := sha256.New()
	hasher.Write([]byte(hash))
	hashResult := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	return hashResult

}

func GenerateRandomString() string {
	s := strconv.FormatInt(time.Now().UnixNano(), 16)
	fmt.Println(s[8:])
	return s[8:]
}
