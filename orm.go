package utils

import (
	"database/sql/driver"
	"fmt"
	"reflect"

	"github.com/lazygophers/log"
	"github.com/lazygophers/utils/defaults"
	"github.com/lazygophers/utils/json"
	"github.com/pkg/errors"
)

// Scan 扫描数据库字段到结构体
func Scan(src interface{}, dst interface{}) (err error) {
	x := func(buf []byte) error {
		if len(buf) > 0 {
			err = json.Unmarshal(buf, dst)
			if err != nil {
				log.Errorf("err:%v", err)
				return err
			}
			return nil
		}
		defaults.SetDefaults(dst)
		return nil
	}

	switch r := src.(type) {
	case []byte:
		return x(r)
	case string:
		return x([]byte(r))
	default:
		return errors.New(
			fmt.Sprintf("unknown type %v %s to scan", r, reflect.ValueOf(src).String()))
	}
}

// Value 将结构体转换为数据库值
func Value(m interface{}) (value driver.Value, err error) {
	if m == nil {
		return []byte("null"), nil
	}

	// Apply defaults only for non-nil structs and struct pointers
	v := reflect.ValueOf(m)
	if v.Kind() == reflect.Struct || (v.Kind() == reflect.Ptr && !v.IsNil() && v.Elem().Kind() == reflect.Struct) {
		defaults.SetDefaults(m)
	}

	value, err = json.Marshal(m)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}
	return value, nil
}
