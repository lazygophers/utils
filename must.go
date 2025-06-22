package utils

import "github.com/lazygophers/log"

// MustOk 如果 ok 为 false 则触发 panic
// 返回 T 类型值用于链式调用
//
// 参数:
//
//	value: 需要验证的值
//	ok:    状态标识
//
// 返回:
//
//	T 类型值
func MustOk[T any](value T, ok bool) T {
	if !ok {
		log.Panic("is not ok")
	}
	return value
}

// MustSuccess 如果 error 不为 nil 则格式化 panic
// 用于验证函数执行状态
//
// 参数:
//
//	err: 错误对象
func MustSuccess(err error) {
	if err != nil {
		log.Panicf("err:%s", err)
	}
}

// Must 组合验证函数
// 先验证错误状态，成功后返回值对象
//
// 参数:
//
//	value: 准备返回的值对象
//	err:   需要验证的错误对象
//
// 返回:
//
//	T 类型值
func Must[T any](value T, err error) T {
	MustSuccess(err)
	return value
}

// Ignore 强制忽略任意参数
// 用于处理需要显式声明但实际不使用的返回值
//
// 参数:
//
//	value: 需要返回的值
//	_:     被忽略的任意参数
//
// 返回:
//
//	T 类型值
func Ignore[T any](value T, _ any) T {
	return value
}
