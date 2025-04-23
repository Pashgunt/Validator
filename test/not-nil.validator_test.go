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
	validatorNameNotNil = "NotNil"
)

type notNilArgs struct {
	constraint contract.ConstraintInterface
	value      interface{}
	exception  pkginterface.ValidationFailedExceptionInterface
}

func newNotNilArgs(value interface{}) *notNilArgs {
	return &notNilArgs{
		value:      value,
		constraint: factory.NewNotBlank(testhelper.DefaultErrorMessage),
		exception:  &violation.ValidationFailedException{},
	}
}

func TestNotNilValidator(t *testing.T) {
	tests := []struct {
		name string
		notNilArgs
		resultValidator testhelper.ResultValidator
	}{
		{
			name:       "test isset error message " + validatorNameNotNil + " validator",
			notNilArgs: *newNotNilArgs(nil),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultIssetErrorCount,
				testhelper.DefaultErrorMessage,
			),
		},
		{
			name:       "test blank error message " + validatorNameNotNil + " validator",
			notNilArgs: *newNotNilArgs("str"),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultNotIssetErrorCount,
				testhelper.BlankString,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			(&validatorprocess.NotNilValidator{}).Process(
				test.notNilArgs.constraint,
				test.notNilArgs.value,
				test.notNilArgs.exception,
			)

			if len(test.exception.Violations()) != test.resultValidator.Count() {
				t.Errorf(
					validatorNameNotNil+" error count exceptions %v, want %v",
					len(test.exception.Violations()),
					test.resultValidator.Count(),
				)
			}

			for _, viol := range test.exception.Violations() {
				if viol.Message().Message() == test.resultValidator.Message() {
					continue
				}

				t.Errorf(
					validatorNameNotNil+" error message exceptions %v, want %v",
					viol.Message().Message(),
					test.resultValidator.Message(),
				)
			}
		})
	}
}
