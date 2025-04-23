package test

import (
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/validator"
	"github.com/Pashgunt/Validator/internal/violation"
	"github.com/Pashgunt/Validator/pkg/factory"
	"github.com/Pashgunt/Validator/pkg/interface"
	testhelper "github.com/Pashgunt/Validator/test/helper"
	"reflect"
	"testing"
)

const (
	validatorNameIsType = "IsType"
)

type isTypeArgs struct {
	constraint contract.ConstraintInterface
	value      interface{}
	exception  pkginterface.ValidationFailedExceptionInterface
}

func newIsTypeArgs(value interface{}, dataType reflect.Kind) *isTypeArgs {
	return &isTypeArgs{
		value:      value,
		constraint: factory.NewIsType(testhelper.DefaultErrorMessage, dataType),
		exception:  &violation.ValidationFailedException{},
	}
}

func TestIsTypeValidator(t *testing.T) {
	tests := []struct {
		name string
		isTypeArgs
		resultValidator testhelper.ResultValidator
	}{
		{
			name:       "test isset error message " + validatorNameIsType + " validator",
			isTypeArgs: *newIsTypeArgs("str", reflect.Int),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultIssetErrorCount,
				testhelper.DefaultErrorMessage,
			),
		},
		{
			name:       "test blank error message " + validatorNameIsType + " validator",
			isTypeArgs: *newIsTypeArgs(123, reflect.Int),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultNotIssetErrorCount,
				testhelper.BlankString,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			(&validatorprocess.IsTypeValidator{}).Process(
				test.isTypeArgs.constraint,
				test.isTypeArgs.value,
				test.isTypeArgs.exception,
			)

			if len(test.exception.Violations()) != test.resultValidator.Count() {
				t.Errorf(
					validatorNameIsType+" error count exceptions %v, want %v",
					len(test.exception.Violations()),
					test.resultValidator.Count(),
				)
			}

			for _, viol := range test.exception.Violations() {
				if viol.Message().Message() == test.resultValidator.Message() {
					continue
				}

				t.Errorf(
					validatorNameIsType+" error message exceptions %v, want %v",
					viol.Message().Message(),
					test.resultValidator.Message(),
				)
			}
		})
	}
}
