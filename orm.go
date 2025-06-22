// utils 包提供数据库操作的实用工具
// 主要功能包括：
// 1. 扫描数据库字段到结构体(Scan)
// 2. 将结构体转换为数据库值(Value)
// 3. 自动填充默认值和错误处理
package utils

import (
	"database/sql/driver"
	"fmt"
	"reflect"

	"github.com/lazygophers/log"
	"github.com/lazygophers/utils/json"

	"github.com/mcuadros/go-defaults"
	"github.com/pkg/errors"
)

// Scan 实现数据库字段到目标结构体的扫描逻辑
// 支持[]byte和string类型的源数据，自动调用defaults.SetDefaults填充默认值
// 参数:
//   - src: 数据源（[]byte或string）
//   - dst: 目标结构体指针
//
// 返回:
//   - error: 反序列化失败时返回错误
//
// 特殊处理:
// 1. 当src为JSON数组或对象时调用json.Unmarshal
// 2. 空字符串时调用defaults.SetDefaults
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
