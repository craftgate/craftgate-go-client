package adapter

import (
	"net/url"
	"time"

	"github.com/gorilla/schema"
)

const (
	TimeLayout = "\"2006-01-02T15:04:05\""
)

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(b []byte) error {
	parse, err := time.Parse(TimeLayout, string(b))
	if err != nil {
		return err
	}
	t.Time = parse
	return nil
}

var encoder = schema.NewEncoder()

func QueryParams(req interface{}) (string, error) {
	queryParams := url.Values{}
	err := encoder.Encode(req, queryParams)
	if err != nil {
		return "", err
	}
	return queryParams.Encode(), nil
}
