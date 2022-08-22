package craftgate

import (
	"bytes"
	"context"
	"craftgate-go-client/adapter"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type ClientOption func(*Client) error

type Client struct {
	httpClient *http.Client
	baseURL    *url.URL
	apiKey     string
	secretKey  string
	headers    map[string]string

	Installment         *adapter.Installment
	Payment             adapter.Payment
	PaymentReporting    adapter.PaymentReporting
	Wallet              adapter.Wallet
	Onboarding          adapter.Onboarding
	PayByLink           adapter.PayByLink
	Settlement          adapter.Settlement
	SettlementReporting adapter.SettlementReporting
}

func New(apiKey, apiSecret string, opts ...ClientOption) (*Client, error) {
	client := newClient(apiKey, apiSecret)
	for _, option := range opts {
		if err := option(client); err != nil {
			return nil, err
		}
	}
	if client.httpClient == nil {
		client.httpClient = http.DefaultClient
	}
	if client.baseURL == nil {
		client.baseURL, _ = url.Parse("https://.....")
	}
	client.headers = make(map[string]string)
	return client, nil
}

func newClient(apiKey, secretKey string) *Client {
	cgClient := &Client{apiKey: apiKey, secretKey: secretKey}
	cgClient.Installment = &adapter.Installment{&client}
	return &Client{
		Installment:         adapter.Installment{Opts: options},
		Payment:             adapter.Payment{Opts: options},
		PaymentReporting:    adapter.PaymentReporting{Opts: options},
		Wallet:              adapter.Wallet{Opts: options},
		Onboarding:          adapter.Onboarding{Opts: options},
		PayByLink:           adapter.PayByLink{Opts: options},
		Settlement:          adapter.Settlement{Opts: options},
		SettlementReporting: adapter.SettlementReporting{Opts: options},
	}
}

func (c *Client) NewRequest(ctx context.Context, method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.baseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var req *http.Request
	switch method {
	case http.MethodGet, http.MethodHead, http.MethodOptions:
		req, err = http.NewRequest(method, u.String(), nil)
		if err != nil {
			return nil, err
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
		req.Header.Set("Content-Type", "application/json")
	}

	for k, v := range c.headers {
		req.Header.Add(k, v)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Go/1.19")

	return req, nil
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	resp, err := DoRequestWithClient(ctx, c.httpClient, req)
	if err != nil {
		return nil, err
	}

	defer func() {

		if rerr := resp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	response := newResponse(resp)

	err = CheckResponse(resp)
	if err != nil {
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
			if err != nil {
				return nil, err
			}
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err != nil {
				return nil, err
			}
		}
	}

	return response, err
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := io.ReadAll(r.Body)
	if err == nil && len(data) > 0 {
		err := json.Unmarshal(data, errorResponse)
		if err != nil {
			errorResponse.Message = string(data)
		}
	}

	return errorResponse
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Message)
}

type ErrorResponse struct {
	Response  *http.Response
	Message   string
	RequestID string
}

func DoRequestWithClient(
	ctx context.Context,
	client *http.Client,
	req *http.Request) (*http.Response, error) {
	req = req.WithContext(ctx)
	return client.Do(req)
}

func newResponse(r *http.Response) *Response {
	response := Response{Response: r}

	return &response
}

type Response struct {
	*http.Response
	//additional fields like rate limiting etc
}

func SetRequestHeaders(headers map[string]string) ClientOption {
	return func(c *Client) error {
		for k, v := range headers {
			c.headers[k] = v
		}
		return nil
	}
}
