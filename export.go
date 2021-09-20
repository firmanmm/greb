package greb

import (
	"os"

	"github.com/go-playground/validator/v10"
)

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
