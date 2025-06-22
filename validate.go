// validate 包提供了结构体校验功能，基于go-playground/validator库实现
package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/lazygophers/log"
)

// Validate函数用于校验结构体对象
//
// 参数:
//   - m: 需要校验的结构体对象
//
// 返回:
//   - error: 校验失败时返回错误信息
//
// 错误处理会自动记录到日志系统
var validate = validator.New()

func Validate(m interface{}) (err error) {
	err = validate.Struct(m)
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}
	return nil
}
