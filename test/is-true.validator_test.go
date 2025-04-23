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
	validatorNameIsTrue = "IsTrue"
)

type isTrueArgs struct {
	constraint contract.ConstraintInterface
	value      interface{}
	exception  pkginterface.ValidationFailedExceptionInterface
}

func newIsTrueArgs(value interface{}) *isTrueArgs {
	return &isTrueArgs{
		value:      value,
		constraint: factory.NewIsTrue(testhelper.DefaultErrorMessage),
		exception:  &violation.ValidationFailedException{},
	}
}

func TestIsTrueValidator(t *testing.T) {
	tests := []struct {
		name string
		isTrueArgs
		resultValidator testhelper.ResultValidator
	}{
		{
			name:       "test isset error message " + validatorNameIsTrue + " validator",
			isTrueArgs: *newIsTrueArgs(false),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultIssetErrorCount,
				testhelper.DefaultErrorMessage,
			),
		},
		{
			name:       "test blank error message " + validatorNameIsTrue + " validator",
			isTrueArgs: *newIsTrueArgs(true),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultNotIssetErrorCount,
				testhelper.BlankString,
			),
		},
		{
			name:       "test blank error message " + validatorNameIsTrue + " validator with not bool data type",
			isTrueArgs: *newIsTrueArgs("str"),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultNotIssetErrorCount,
				testhelper.BlankString,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			(&validatorprocess.IsTrueValidator{}).Process(
				test.isTrueArgs.constraint,
				test.isTrueArgs.value,
				test.isTrueArgs.exception,
			)

			if len(test.exception.Violations()) != test.resultValidator.Count() {
				t.Errorf(
					validatorNameIsTrue+" error count exceptions %v, want %v",
					len(test.exception.Violations()),
					test.resultValidator.Count(),
				)
			}

			for _, viol := range test.exception.Violations() {
				if viol.Message().Message() == test.resultValidator.Message() {
					continue
				}

				t.Errorf(
					validatorNameIsTrue+" error message exceptions %v, want %v",
					viol.Message().Message(),
					test.resultValidator.Message(),
				)
			}
		})
	}
}
