package adapter

import (
	"net/url"

	"github.com/gorilla/schema"
)

var encoder = schema.NewEncoder()

func QueryParams(req interface{}) (string, error) {
	queryParams := url.Values{}
	err := encoder.Encode(req, queryParams)
	if err != nil {
		return "", err
	}
	return queryParams.Encode(), nil
}
