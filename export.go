package greb

import (
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
)

type URLParamBind func(req *http.Request, key string) string

var URLParamBindFunc URLParamBind
var _InternalValidator = validator.New()

func Generate(fileName string) (string, error) {
	rawByte, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	ast, err := _InternalParser.Parse(string(rawByte))
	if err != nil {
		return "", err
	}
	return _InternalGenerator.Generate(ast)
}

func Validate(data IBindable) error {
	return _InternalValidator.Struct(data)
}
