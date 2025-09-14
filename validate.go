package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/lazygophers/log"
)

var validate = validator.New()

// Validate 校验结构体对象

func Validate(m interface{}) (err error) {
	err = validate.Struct(m)
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}
	return nil
}
