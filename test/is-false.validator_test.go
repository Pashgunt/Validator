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
	validatorNameIsFalse = "IsFalse"
)

type isFalseArgs struct {
	constraint contract.ConstraintInterface
	value      interface{}
	exception  pkginterface.ValidationFailedExceptionInterface
}

func newIsFalseArgs(value interface{}) *isFalseArgs {
	return &isFalseArgs{
		value:      value,
		constraint: factory.NewIsFalse(testhelper.DefaultErrorMessage),
		exception:  &violation.ValidationFailedException{},
	}
}

func TestIsFalseValidator(t *testing.T) {
	tests := []struct {
		name string
		isFalseArgs
		resultValidator testhelper.ResultValidator
	}{
		{
			name:        "test isset error message " + validatorNameIsFalse + " validator",
			isFalseArgs: *newIsFalseArgs(true),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultIssetErrorCount,
				testhelper.DefaultErrorMessage,
			),
		},
		{
			name:        "test blank error message " + validatorNameIsFalse + " validator",
			isFalseArgs: *newIsFalseArgs(false),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultNotIssetErrorCount,
				testhelper.BlankString,
			),
		},
		{
			name:        "test blank error message " + validatorNameIsFalse + " validator with not bool data type",
			isFalseArgs: *newIsFalseArgs("str"),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultNotIssetErrorCount,
				testhelper.BlankString,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			(&validatorprocess.IsFalseValidator{}).Process(
				test.isFalseArgs.constraint,
				test.isFalseArgs.value,
				test.isFalseArgs.exception,
			)

			if len(test.exception.Violations()) != test.resultValidator.Count() {
				t.Errorf(
					validatorNameIsFalse+" error count exceptions %v, want %v",
					len(test.exception.Violations()),
					test.resultValidator.Count(),
				)
			}

			for _, viol := range test.exception.Violations() {
				if viol.Message().Message() == test.resultValidator.Message() {
					continue
				}

				t.Errorf(
					validatorNameIsFalse+" error message exceptions %v, want %v",
					viol.Message().Message(),
					test.resultValidator.Message(),
				)
			}
		})
	}
}
