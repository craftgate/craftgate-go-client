package adapter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/schema"
	"net/url"
	"reflect"
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

func init() {
	timeConverter := func(value reflect.Value) string {
		timestamp := fmt.Sprintf("%s", value.Interface().(time.Time).Format(timeEncodeLayout))
		return timestamp
	}

	schemaEncoder.RegisterEncoder(time.Time{}, timeConverter)
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

func PrepareBody(req interface{}) (*bytes.Buffer, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(body), nil
}
