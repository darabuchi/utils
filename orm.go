package utils

import (
	"database/sql/driver"
	"fmt"
	"reflect"

	"github.com/bytedance/sonic"
	"github.com/mcuadros/go-defaults"
	"github.com/pkg/errors"
)

func Scan(src interface{}, dst interface{}) error {
	x := func(buf []byte) error {
		bufLen := len(buf)
		if bufLen >= 2 && buf[0] == '{' && buf[bufLen-1] == '}' {
			return sonic.Unmarshal(buf, dst)
		} else if bufLen > 0 {
			return sonic.Unmarshal(buf, dst)
		} else {
			defaults.SetDefaults(dst)
		}
		return nil
	}

	switch r := src.(type) {
	case []byte:
		buf := src.([]byte)
		return x(buf)
	case string:
		buf := []byte(src.(string))
		return x(buf)
	default:
		return errors.New(
			fmt.Sprintf("unknown type %v %s to scan", r, reflect.ValueOf(src).String()))
	}
}

func Value(m interface{}) (driver.Value, error) {
	if m == nil {
		defaults.SetDefaults(m)
	}

	err := Validate(m)
	if err != nil {
		return nil, err
	}

	return sonic.Marshal(m)
}
