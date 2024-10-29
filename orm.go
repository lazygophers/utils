package utils

import (
	"database/sql/driver"
	"fmt"
	"github.com/lazygophers/log"
	"github.com/lazygophers/utils/json"
	"reflect"

	"github.com/mcuadros/go-defaults"
	"github.com/pkg/errors"
)

func Scan(src interface{}, dst interface{}) (err error) {
	x := func(buf []byte) error {
		bufLen := len(buf)
		if bufLen >= 2 && ((buf[0] == '{' && buf[bufLen-1] == '}') || (buf[0] == '[' && buf[bufLen-1] == ']')) {
			err = json.Unmarshal(buf, dst)
			if err != nil {
				log.Errorf("err:%v", err)
				return err
			}
			return nil
		} else if bufLen > 0 {
			err = json.Unmarshal(buf, dst)
			if err != nil {
				log.Errorf("err:%v", err)
				return err
			}
			return nil
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

func Value(m interface{}) (value driver.Value, err error) {
	if m == nil {
		defaults.SetDefaults(m)
	}

	value, err = json.Marshal(m)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}
	return value, nil
}
