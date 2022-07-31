package adapter

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/url"
	"reflect"
	"time"

	"github.com/gorilla/schema"
)

var ErrNull = errors.New("can't be both IsNull and have value")
var Null = []byte("null")

var schemaEncoder = schema.NewEncoder()

const (
	TimeLayout = "\"2006-01-02T15:04:05\""
)

type Boolean struct {
	bool
	IsNull bool
}

type Integer struct {
	int
	IsNull bool
}

type Long struct {
	int64
	IsNull bool
}

type BigDecimal struct {
	float64
	IsNull bool
}

type String struct {
	string
	IsNull bool
}

type Time struct {
	time.Time
	IsNull bool
}

func NewBoolean(v bool) *Boolean {
	return &Boolean{v, false}
}

func NewInteger(v int) *Integer {
	return &Integer{v, false}
}

func NewLong(v int64) *Long {
	return &Long{v, false}
}

func NewBigDecimal(v float64) *BigDecimal {
	return &BigDecimal{v, false}
}

func NewString(v string) *String {
	return &String{v, false}
}

func NewTime(v time.Time) *Time {
	return &Time{v, false}
}

func (v Boolean) MarshalJSON() ([]byte, error) {
	switch {
	case v.IsNull && v.bool:
		return nil, ErrNull
	case v.IsNull:
		return Null, nil
	default:
		return json.Marshal(v.bool)
	}
}

func (v *Boolean) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, Null) {
		v.IsNull = true
		return nil
	}
	return json.Unmarshal(src, &v.bool)
}

func (v Integer) MarshalJSON() ([]byte, error) {
	switch {
	case v.IsNull && v.int != 0:
		return nil, ErrNull
	case v.IsNull:
		return Null, nil
	default:
		return json.Marshal(v.int)
	}
}

func (v *Integer) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, Null) {
		v.IsNull = true
		return nil
	}
	return json.Unmarshal(src, &v.int)
}

func (v Long) MarshalJSON() ([]byte, error) {
	switch {
	case v.IsNull && v.int64 != 0:
		return nil, ErrNull
	case v.IsNull:
		return Null, nil
	default:
		return json.Marshal(v.int64)
	}
}

func (v *Long) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, Null) {
		v.IsNull = true
		return nil
	}
	return json.Unmarshal(src, &v.int64)
}

func (v BigDecimal) MarshalJSON() ([]byte, error) {
	switch {
	case v.IsNull && v.float64 != 0.0:
		return nil, ErrNull
	case v.IsNull:
		return Null, nil
	default:
		return json.Marshal(v.float64)
	}
}

func (v *BigDecimal) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, Null) {
		v.IsNull = true
		return nil
	}
	return json.Unmarshal(src, &v.float64)
}

func (v String) MarshalJSON() ([]byte, error) {
	switch {
	case v.IsNull && v.string != "":
		return nil, ErrNull
	case v.IsNull:
		return Null, nil
	default:
		return json.Marshal(v.string)
	}
}

func (v *String) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, Null) {
		v.IsNull = true
		return nil
	}
	return json.Unmarshal(src, &v.string)
}

func (v Time) MarshalJSON() ([]byte, error) {
	switch {
	case v.IsNull && !v.Time.IsZero():
		return nil, ErrNull
	case v.IsNull:
		return Null, nil
	default:
		return v.Time.MarshalJSON()
	}
}

func (v *Time) UnmarshalJSON(src []byte) error {
	parse, err := time.Parse(TimeLayout, string(src))
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
	//removeNulls(queryParams)

	for k, v := range queryParams {
		if v == nil {
			queryParams.Del(k)
		}
	}
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

func PrepareBody(req interface{}) (*bytes.Buffer, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(body), nil
}
