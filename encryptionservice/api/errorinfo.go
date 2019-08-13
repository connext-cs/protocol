package api

import "errors"

const (
	ParameterError = "参数错误"
)

func Gerr(str string) error {
	return errors.New(str)
}
