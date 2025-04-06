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
	validatorNameNotBlank = "NotBlank"
)

type notBlankArgs struct {
	constraint contract.ConstraintInterface
	value      interface{}
	exception  contract.ValidationFailedExceptionInterface
}

func newNotBlankArgs(value interface{}) *notBlankArgs {
	return &notBlankArgs{
		value:      value,
		constraint: validator.NewNotBlank(testhelper.DefaultErrorMessage),
		exception:  &violation.ValidationFailedException{},
	}
}

func TestNotBlankValidator(t *testing.T) {
	tests := []struct {
		name string
		notBlankArgs
		resultValidator testhelper.ResultValidator
	}{
		{
			name:         "test isset error message " + validatorNameNotBlank + " validator",
			notBlankArgs: *newNotBlankArgs(testhelper.BlankString),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultIssetErrorCount,
				testhelper.DefaultErrorMessage,
			),
		},
		{
			name:         "test blank error message " + validatorNameNotBlank + " validator",
			notBlankArgs: *newNotBlankArgs("Value"),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultNotIssetErrorCount,
				testhelper.BlankString,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			(&validatorprocess.NotBlankValidator{}).Process(
				test.notBlankArgs.constraint,
				test.notBlankArgs.value,
				test.notBlankArgs.exception,
			)

			if len(test.exception.Violations()) != test.resultValidator.Count() {
				t.Errorf(
					validatorNameNotBlank+" error count exceptions %v, want %v",
					len(test.exception.Violations()),
					test.resultValidator.Count(),
				)
			}

			for _, viol := range test.exception.Violations() {
				if viol.Message().Message() == test.resultValidator.Message() {
					continue
				}

				t.Errorf(
					validatorNameNotBlank+" error message exceptions %v, want %v",
					viol.Message().Message(),
					test.resultValidator.Message(),
				)
			}
		})
	}
}
