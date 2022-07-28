package adapter

import (
	"net/url"
	"reflect"
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
	//removeNulls(queryParams)

	return queryParams.Encode(), nil
}

func removeNulls(m map[string]interface{}) {
	val := reflect.ValueOf(m)
	for _, e := range val.MapKeys() {
		v := val.MapIndex(e)
		if v.IsNil() {
			delete(m, e.String())
			continue
		}
	}
}
