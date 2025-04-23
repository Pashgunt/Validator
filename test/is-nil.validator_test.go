package test

import (
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/validator"
	"github.com/Pashgunt/Validator/internal/violation"
	"github.com/Pashgunt/Validator/pkg/factory"
	"github.com/Pashgunt/Validator/pkg/interface"
	testhelper "github.com/Pashgunt/Validator/test/helper"
	"testing"
)

const (
	validatorNameIsNil = "IsNil"
)

type isNilArgs struct {
	constraint contract.ConstraintInterface
	value      interface{}
	exception  pkginterface.ValidationFailedExceptionInterface
}

func newIsNilArgs(value interface{}) *isNilArgs {
	return &isNilArgs{
		value:      value,
		constraint: factory.NewBlank(testhelper.DefaultErrorMessage),
		exception:  &violation.ValidationFailedException{},
	}
}

func TestIsNilValidator(t *testing.T) {
	tests := []struct {
		name string
		isNilArgs
		resultValidator testhelper.ResultValidator
	}{
		{
			name:      "test isset error message " + validatorNameIsNil + " validator",
			isNilArgs: *newIsNilArgs(false),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultIssetErrorCount,
				testhelper.DefaultErrorMessage,
			),
		},
		{
			name:      "test blank error message " + validatorNameIsNil + " validator",
			isNilArgs: *newIsNilArgs(nil),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultNotIssetErrorCount,
				testhelper.BlankString,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			(&validatorprocess.IsNilValidator{}).Process(
				test.isNilArgs.constraint,
				test.isNilArgs.value,
				test.isNilArgs.exception,
			)

			if len(test.exception.Violations()) != test.resultValidator.Count() {
				t.Errorf(
					validatorNameIsNil+" error count exceptions %v, want %v",
					len(test.exception.Violations()),
					test.resultValidator.Count(),
				)
			}

			for _, viol := range test.exception.Violations() {
				if viol.Message().Message() == test.resultValidator.Message() {
					continue
				}

				t.Errorf(
					validatorNameIsNil+" error message exceptions %v, want %v",
					viol.Message().Message(),
					test.resultValidator.Message(),
				)
			}
		})
	}
}
