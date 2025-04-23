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
	validatorNameCompromised = "CompromisedPassword"
)

type notCompromisedPasswordArgs struct {
	constraint contract.ConstraintInterface
	value      interface{}
	exception  pkginterface.ValidationFailedExceptionInterface
}

func newNotCompromisedPasswordArgs(value interface{}) *notCompromisedPasswordArgs {
	return &notCompromisedPasswordArgs{
		value:      value,
		constraint: factory.NewNotCompromisedPassword(testhelper.DefaultErrorMessage),
		exception:  &violation.ValidationFailedException{},
	}
}

func TestCompromisedPasswordValidator(t *testing.T) {
	tests := []struct {
		name string
		notCompromisedPasswordArgs
		resultValidator testhelper.ResultValidator
	}{
		{
			name:                       "test isset error message " + validatorNameCompromised + " validator",
			notCompromisedPasswordArgs: *newNotCompromisedPasswordArgs("test"),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultIssetErrorCount,
				testhelper.DefaultErrorMessage,
			),
		},
		{
			name:                       "test blank error message " + validatorNameCompromised + " validator",
			notCompromisedPasswordArgs: *newNotCompromisedPasswordArgs("The_Error12"),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultNotIssetErrorCount,
				testhelper.BlankString,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			(&validatorprocess.CompromisedPasswordValidator{}).Process(
				test.notCompromisedPasswordArgs.constraint,
				test.notCompromisedPasswordArgs.value,
				test.notCompromisedPasswordArgs.exception,
			)

			if len(test.exception.Violations()) != test.resultValidator.Count() {
				t.Errorf(
					validatorNameCompromised+" error count exceptions %v, want %v",
					len(test.exception.Violations()),
					test.resultValidator.Count(),
				)
			}

			for _, viol := range test.exception.Violations() {
				if viol.Message().Message() == test.resultValidator.Message() {
					continue
				}

				t.Errorf(
					validatorNameCompromised+" error message exceptions %v, want %v",
					viol.Message().Message(),
					test.resultValidator.Message(),
				)
			}
		})
	}
}
