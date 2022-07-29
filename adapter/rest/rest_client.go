package rest

import (
	"craftgate-go-client/adapter"
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

func SendRequest(req *http.Request, v interface{}, opts adapter.RequestOptions) error {
	client := &http.Client{
		Timeout: time.Minute,
	}

	randomStr := GenerateRandomString()
	hashStr := GenerateHash(req.URL.String(), opts.ApiKey, opts.SecretKey, randomStr, "")
	fmt.Println(req.URL.String())

	req.Header.Set(adapter.ApiKeyHeaderName, opts.ApiKey)
	req.Header.Set(adapter.RandomHeaderName, randomStr)
	req.Header.Set(adapter.AuthVersionHeaderName, adapter.AuthVersion)
	req.Header.Set(adapter.ClientVersionHeaderName, adapter.ClientVersion)
	req.Header.Set(adapter.SignatureHeaderName, hashStr)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	res, err := client.Do(req)
	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes adapter.Response[any]
		if err = json.NewDecoder(res.Body).Decode(&v); err == nil {
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
	hashStr := strings.Join([]string{url, apiKey, secretKey, randomString, body}, "")
	hash := sha256.New()
	hash.Write([]byte(hashStr))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

func GenerateRandomString() string {
	s := strconv.FormatInt(time.Now().UnixNano(), 16)
	fmt.Println(s[8:])
	return s[8:]
}
