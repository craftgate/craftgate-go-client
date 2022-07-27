package rest

import (
	"craftgate-go-client/model"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func SendRequest(req *http.Request, v interface{}, opts model.RequestOptions) error {
	client := &http.Client{
		Timeout: time.Minute,
	}

	randomStr := GenerateRandomString()
	hashStr := GenerateHash(req.URL.String(), opts.ApiKey, opts.SecretKey, randomStr, "")
	fmt.Println(hashStr)

	req.Header.Set(model.ApiKeyHeaderName, opts.ApiKey)
	req.Header.Set(model.RandomHeaderName, randomStr)
	req.Header.Set(model.AuthVersionHeaderName, model.AuthVersion)
	req.Header.Set(model.ClientVersionHeaderName, model.ClientVersion)
	req.Header.Set(model.SignatureHeaderName, hashStr)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	res, err := client.Do(req)
	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes model.Response[any]
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
	hashStr := []string{url, apiKey, secretKey, randomString, body}
	hash := strings.Join(hashStr, "")
	fmt.Println(hash)
	hasher := sha256.New()
	hasher.Write([]byte(hash))
	hashResult := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return hashResult
}

func GenerateRandomString() string {
	/*s := strconv.FormatInt(time.Now().UnixNano(), 16)
	fmt.Println(s[8:])
	return s[8:]*/
	return "1234567890"
}
