package utils

import "github.com/lazygophers/log"

// MustOk 如果 ok 为 false 则触发 panic
func MustOk[T any](value T, ok bool) T {
	if !ok {
		log.Panic("is not ok")
	}
	return value
}

// MustSuccess 如果 error 不为 nil 则格式化 panic
func MustSuccess(err error) {
	if err != nil {
		log.Panicf("err:%s", err)
	}
}

// Must 组合验证函数，先验证错误状态，成功后返回值对象
func Must[T any](value T, err error) T {
	MustSuccess(err)
	return value
}

// Ignore 强制忽略任意参数
func Ignore[T any](value T, _ any) T {
	return value
}
