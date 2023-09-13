package adapter

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/schema"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var schemaEncoder = schema.NewEncoder()

const (
	timeEncodeLayout = "2006-01-02T15:04:05"
	timeDecodeLayout = "\"2006-01-02T15:04:05\""
)

type TimeResponse struct {
	time.Time
}

type Date struct {
	Year  int        // Year (e.g., 2014).
	Month time.Month // Month of the year (January = 1, ...).
	Day   int        // Day of the month, starting at 1.
}

func DateOf(t time.Time) Date {
	var d Date
	d.Year, d.Month, d.Day = t.Date()
	return d
}

func init() {
	timeConverter := func(value reflect.Value) string {
		return value.Interface().(time.Time).Format(timeEncodeLayout)
	}
	dateConverter := func(value reflect.Value) string {
		date := value.Interface().(Date)
		month := strconv.Itoa(int(date.Month))
		if date.Month < 10 {
			month = "0" + month
		}
		day := strconv.Itoa(date.Day)
		if date.Day < 10 {
			day = "0" + day
		}
		d := fmt.Sprintf("%d-%s-%s", date.Year, month, day)
		return d
	}

	schemaEncoder.RegisterEncoder(time.Time{}, timeConverter)
	schemaEncoder.RegisterEncoder(Date{}, dateConverter)
}

type ClientOption func(*Client) error

type Client struct {
	httpClient *http.Client
	baseURL    *url.URL
	apiKey     string
	secretKey  string
	headers    map[string]string

	Installment         *Installment
	Payment             *Payment
	PaymentReporting    *PaymentReporting
	Wallet              *Wallet
	Onboarding          *Onboarding
	PayByLink           *PayByLink
	Settlement          *Settlement
	SettlementReporting *SettlementReporting
	FileReporting       *FileReporting
	Fraud               *Fraud
	Hook                *Hook
	Masterpass          *Masterpass
}

func New(apiKey, apiSecret, baseURL string, opts ...ClientOption) (*Client, error) {
	client := newClient(apiKey, apiSecret)
	for _, option := range opts {
		if err := option(client); err != nil {
			return nil, err
		}
	}
	if client.httpClient == nil {
		client.httpClient = http.DefaultClient
		client.httpClient.Timeout = time.Minute * 120
	}
	if client.baseURL == nil {
		client.baseURL, _ = url.Parse(baseURL)
	}
	client.headers = make(map[string]string)
	return client, nil
}

func newClient(apiKey, secretKey string) *Client {
	client := &Client{apiKey: apiKey, secretKey: secretKey}

	client.Installment = &Installment{Client: client}
	client.Payment = &Payment{Client: client}
	client.Onboarding = &Onboarding{Client: client}
	client.PaymentReporting = &PaymentReporting{Client: client}
	client.PayByLink = &PayByLink{Client: client}
	client.Wallet = &Wallet{Client: client}
	client.Settlement = &Settlement{Client: client}
	client.SettlementReporting = &SettlementReporting{Client: client}
	client.FileReporting = &FileReporting{Client: client}
	client.Fraud = &Fraud{Client: client}
	client.Hook = &Hook{Client: client}
	client.Masterpass = &Masterpass{Client: client}

	return client
}

func (c *Client) NewRequest(ctx context.Context, method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.baseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var req *http.Request

	switch method {
	case http.MethodGet, http.MethodDelete, http.MethodHead, http.MethodOptions:
		req, err = http.NewRequestWithContext(ctx, method, u.String(), nil)
		if err != nil {
			return nil, err
		}

		if body != nil {
			req.URL.RawQuery, _ = QueryParams(body)
		}
	default:
		buf := new(bytes.Buffer)
		if body != nil {
			err = json.NewEncoder(buf).Encode(body)
			if err != nil {
				return nil, err
			}
		}

		req, err = http.NewRequest(method, u.String(), buf)
		if err != nil {
			return nil, err
		}
	}

	for k, v := range c.headers {
		req.Header.Add(k, v)
	}

	authorizationRequestBody := c.extractRequestBodyForAuthorization(body, method, req)
	randomStr := GenerateRandomString()
	hashStr := GenerateHash(req.URL.String(), c.apiKey, c.secretKey, randomStr, authorizationRequestBody)

	req.Header.Set(ApiKeyHeaderName, c.apiKey)
	req.Header.Set(RandomHeaderName, randomStr)
	req.Header.Set(AuthVersionHeaderName, AuthVersion)
	req.Header.Set(ClientVersionHeaderName, ClientVersion)
	req.Header.Set(SignatureHeaderName, hashStr)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	return req, nil
}

func (c *Client) NewRequestForByteResponse(ctx context.Context, method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.baseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var req *http.Request

	switch method {
	case http.MethodGet, http.MethodDelete, http.MethodHead, http.MethodOptions:
		req, err = http.NewRequestWithContext(ctx, method, u.String(), nil)
		if err != nil {
			return nil, err
		}

		if body != nil {
			req.URL.RawQuery, _ = QueryParams(body)
		}
	default:
		buf := new(bytes.Buffer)
		if body != nil {
			err = json.NewEncoder(buf).Encode(body)
			if err != nil {
				return nil, err
			}
		}

		req, err = http.NewRequest(method, u.String(), buf)
		if err != nil {
			return nil, err
		}
	}

	for k, v := range c.headers {
		req.Header.Add(k, v)
	}

	authorizationRequestBody := c.extractRequestBodyForAuthorization(body, method, req)
	randomStr := GenerateRandomString()
	hashStr := GenerateHash(req.URL.String(), c.apiKey, c.secretKey, randomStr, authorizationRequestBody)

	req.Header.Set(ApiKeyHeaderName, c.apiKey)
	req.Header.Set(RandomHeaderName, randomStr)
	req.Header.Set(AuthVersionHeaderName, AuthVersion)
	req.Header.Set(ClientVersionHeaderName, ClientVersion)
	req.Header.Set(SignatureHeaderName, hashStr)
	req.Header.Set("Content-Type", "application/octet-stream; charset=utf-8")
	req.Header.Set("Accept", "application/octet-stream; charset=utf-8")

	return req, nil
}

func (c *Client) extractRequestBodyForAuthorization(body interface{}, method string, req *http.Request) string {
	authorizationRequestBody := ""
	if body != nil && method != http.MethodGet && method != http.MethodDelete {
		var buf bytes.Buffer
		tee := io.TeeReader(req.Body, &buf)
		req.Body = io.NopCloser(&buf)
		bodyBytes, bodyErr := io.ReadAll(tee)
		if bodyErr == nil {
			authorizationRequestBody = string(bodyBytes)
		}
	}
	return authorizationRequestBody
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) error {
	resp, err := DoRequestWithClient(ctx, c.httpClient, req)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
	}(resp.Body)

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		errRes := &Response[ErrorResponse]{}
		if err = json.NewDecoder(resp.Body).Decode(errRes); nil == errRes {
			return errors.New(*errRes.Errors.ErrorDescription)
		}

		return errRes
	}

	switch v.(type) {
	case *Void:
	default:
		if err = json.NewDecoder(resp.Body).Decode(&v); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) DoForByteResponse(ctx context.Context, req *http.Request) ([]byte, error) {
	resp, err := DoRequestWithClient(ctx, c.httpClient, req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
	}(resp.Body)

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		errRes := &Response[ErrorResponse]{}
		if err = json.NewDecoder(resp.Body).Decode(errRes); nil == errRes {
			return nil, errors.New(*errRes.Errors.ErrorDescription)
		}

		return nil, errRes
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v %v", r.ErrorGroup, r.ErrorCode, r.ErrorDescription)
}

func DoRequestWithClient(
	ctx context.Context,
	client *http.Client,
	req *http.Request) (*http.Response, error) {
	req = req.WithContext(ctx)
	return client.Do(req)
}

func GenerateHash(url, apiKey, secretKey, randomString, body string) string {
	hashStr := strings.Join([]string{url, apiKey, secretKey, randomString, body}, "")
	hash := sha256.New()
	hash.Write([]byte(hashStr))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

func GenerateRandomString() string {
	s := strconv.FormatInt(time.Now().UnixNano(), 16)
	return s[8:]
}

func (v *TimeResponse) UnmarshalJSON(b []byte) error {
	parse, err := time.Parse(timeDecodeLayout, string(b))
	if err != nil {
		return err
	}
	v.Time = parse
	return nil
}

func QueryParams(req interface{}) (string, error) {
	queryParams := url.Values{}
	err := schemaEncoder.Encode(req, queryParams)
	if err != nil {
		return "", err
	}
	encoded := queryParams.Encode()
	encoded = strings.Replace(encoded, "%3A", ":", -1)
	return encoded, nil
}

func (c *Client) Info() string {
	return "Craftgate Go Client. Api Key: " + c.apiKey
}
