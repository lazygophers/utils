package utils

import "github.com/lazygophers/log"

func MustOk[T any](value T, ok bool) T {
	if !ok {
		log.Panic("is not ok")
	}
	return value
}

func MustSuccess(err error) {
	if err != nil {
		log.Panicf("err:%s", err)
	}
}

func Must[T any](value T, err error) T {
	MustSuccess(err)
	return value
}
