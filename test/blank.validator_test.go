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
	validatorNameBlank = "Blank"
)

type blankArgs struct {
	constraint contract.ConstraintInterface
	value      interface{}
	exception  pkginterface.ValidationFailedExceptionInterface
}

func newBlankArgs(value interface{}) *blankArgs {
	return &blankArgs{
		value:      value,
		constraint: factory.NewNotBlank(testhelper.DefaultErrorMessage),
		exception:  &violation.ValidationFailedException{},
	}
}

func TestBlankValidator(t *testing.T) {
	tests := []struct {
		name string
		blankArgs
		resultValidator testhelper.ResultValidator
	}{
		{
			name:      "test isset error message " + validatorNameBlank + " validator",
			blankArgs: *newBlankArgs("Value"),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultIssetErrorCount,
				testhelper.DefaultErrorMessage,
			),
		},
		{
			name:      "test blank error message " + validatorNameBlank + " validator",
			blankArgs: *newBlankArgs(testhelper.BlankString),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultNotIssetErrorCount,
				testhelper.BlankString,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			(&validatorprocess.BlankValidator{}).Process(
				test.blankArgs.constraint,
				test.blankArgs.value,
				test.blankArgs.exception,
			)

			if len(test.exception.Violations()) != test.resultValidator.Count() {
				t.Errorf(
					validatorNameBlank+" error count exceptions %v, want %v",
					len(test.exception.Violations()),
					test.resultValidator.Count(),
				)
			}

			for _, viol := range test.exception.Violations() {
				if viol.Message().Message() == test.resultValidator.Message() {
					continue
				}

				t.Errorf(
					validatorNameBlank+" error message exceptions %v, want %v",
					viol.Message().Message(),
					test.resultValidator.Message(),
				)
			}
		})
	}
}
