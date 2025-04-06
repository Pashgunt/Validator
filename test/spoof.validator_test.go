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
	validatorNameSpoof = "Spoof"
)

type spoofArgs struct {
	constraint contract.ConstraintInterface
	value      interface{}
	exception  contract.ValidationFailedExceptionInterface
}

func newSpoofArgs(value interface{}) *spoofArgs {
	return &spoofArgs{
		value:      value,
		constraint: validator.NewNotBlank(testhelper.DefaultErrorMessage),
		exception:  &violation.ValidationFailedException{},
	}
}

func TestSpoofValidator(t *testing.T) {
	tests := []struct {
		name string
		spoofArgs
		resultValidator testhelper.ResultValidator
	}{
		{
			name:      "test isset error message " + validatorNameSpoof + " validator",
			spoofArgs: *newSpoofArgs("Hello\xA0World"),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultIssetErrorCount,
				testhelper.DefaultErrorMessage,
			),
		},
		{
			name:      "test blank error message " + validatorNameSpoof + " validator",
			spoofArgs: *newSpoofArgs("Hello World"),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultNotIssetErrorCount,
				testhelper.BlankString,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			(&validatorprocess.SpoofValidator{}).Process(
				test.spoofArgs.constraint,
				test.spoofArgs.value,
				test.spoofArgs.exception,
			)

			if len(test.exception.Violations()) != test.resultValidator.Count() {
				t.Errorf(
					validatorNameSpoof+" error count exceptions %v, want %v",
					len(test.exception.Violations()),
					test.resultValidator.Count(),
				)
			}

			for _, viol := range test.exception.Violations() {
				if viol.Message().Message() == test.resultValidator.Message() {
					continue
				}

				t.Errorf(
					validatorNameSpoof+" error message exceptions %v, want %v",
					viol.Message().Message(),
					test.resultValidator.Message(),
				)
			}
		})
	}
}
