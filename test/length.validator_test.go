package test

import (
	validator "github.com/Pashgunt/Validator"
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/validator"
	"github.com/Pashgunt/Validator/internal/violation"
	testhelper "github.com/Pashgunt/Validator/test/helper"
	"testing"
)

const (
	validatorNameLength = "Length"
)

type lengthArgs struct {
	constraint contract.ConstraintInterface
	value      interface{}
	exception  contract.ValidationFailedExceptionInterface
}

func newLengthArgs(value interface{}, min int, max int, minMessage string, maxMessage string) *lengthArgs {
	return &lengthArgs{
		value:      value,
		constraint: validator.NewLength(min, max, minMessage, maxMessage),
		exception:  &violation.ValidationFailedException{},
	}
}

func TestLengthValidator(t *testing.T) {
	tests := []struct {
		name string
		lengthArgs
		resultValidator testhelper.ResultValidator
	}{
		{
			name: "test isset error message " + validatorNameLength + " validator",
			lengthArgs: *newLengthArgs(
				"str",
				5,
				10,
				testhelper.DefaultErrorMessage,
				testhelper.DefaultErrorMessage,
			),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultIssetErrorCount,
				testhelper.DefaultErrorMessage,
			),
		},
		{
			name: "test blank error message " + validatorNameLength + " validator",
			lengthArgs: *newLengthArgs(
				"string",
				0,
				10,
				testhelper.DefaultErrorMessage,
				testhelper.DefaultErrorMessage,
			),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultNotIssetErrorCount,
				testhelper.BlankString,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			(&validatorprocess.LengthValidator{}).Process(
				test.lengthArgs.constraint,
				test.lengthArgs.value,
				test.lengthArgs.exception,
			)

			if len(test.exception.Violations()) != test.resultValidator.Count() {
				t.Errorf(
					validatorNameLength+" error count exceptions %v, want %v",
					len(test.exception.Violations()),
					test.resultValidator.Count(),
				)
			}

			for _, viol := range test.exception.Violations() {
				if viol.Message().Message() == test.resultValidator.Message() {
					continue
				}

				t.Errorf(
					validatorNameLength+" error message exceptions %v, want %v",
					viol.Message().Message(),
					test.resultValidator.Message(),
				)
			}
		})
	}
}
