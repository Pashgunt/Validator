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
	validatorNameRegex = "Regex"
)

type regexArgs struct {
	constraint contract.ConstraintInterface
	value      interface{}
	exception  contract.ValidationFailedExceptionInterface
}

func newRegexArgs(value interface{}, pattern string) *regexArgs {
	return &regexArgs{
		value:      value,
		constraint: validator.NewRegex(pattern, testhelper.DefaultErrorMessage),
		exception:  &violation.ValidationFailedException{},
	}
}

func TestRegexValidator(t *testing.T) {
	tests := []struct {
		name string
		regexArgs
		resultValidator testhelper.ResultValidator
	}{
		{
			name:      "test isset error message " + validatorNameRegex + " validator",
			regexArgs: *newRegexArgs("", `[a-zA-Z]`),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultIssetErrorCount,
				testhelper.DefaultErrorMessage,
			),
		},
		{
			name:      "test blank error message " + validatorNameRegex + " validator",
			regexArgs: *newRegexArgs("Value", `[a-zA-Z]`),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultNotIssetErrorCount,
				testhelper.BlankString,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			(&validatorprocess.RegexValidator{}).Process(
				test.regexArgs.constraint,
				test.regexArgs.value,
				test.regexArgs.exception,
			)

			if len(test.exception.Violations()) != test.resultValidator.Count() {
				t.Errorf(
					validatorNameRegex+" error count exceptions %v, want %v",
					len(test.exception.Violations()),
					test.resultValidator.Count(),
				)
			}

			for _, viol := range test.exception.Violations() {
				if viol.Message().Message() == test.resultValidator.Message() {
					continue
				}

				t.Errorf(
					validatorNameRegex+" error message exceptions %v, want %v",
					viol.Message().Message(),
					test.resultValidator.Message(),
				)
			}
		})
	}
}
