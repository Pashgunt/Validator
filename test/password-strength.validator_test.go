package test

import (
	validator "github.com/Pashgunt/Validator"
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/validator"
	"github.com/Pashgunt/Validator/internal/violation"
	"github.com/Pashgunt/Validator/strength"
	testhelper "github.com/Pashgunt/Validator/test/helper"
	"testing"
)

const (
	validatorNamePasswordStrength = "PasswordStrength"
)

type PasswordStrengthArgs struct {
	constraint contract.ConstraintInterface
	value      interface{}
	exception  contract.ValidationFailedExceptionInterface
}

func newPasswordStrengthArgs(value interface{}, minScore int) *PasswordStrengthArgs {
	return &PasswordStrengthArgs{
		value:      value,
		constraint: validator.NewPasswordStrength(testhelper.DefaultErrorMessage, minScore),
		exception:  &violation.ValidationFailedException{},
	}
}

func TestPasswordStrengthValidator(t *testing.T) {
	tests := []struct {
		name string
		PasswordStrengthArgs
		resultValidator testhelper.ResultValidator
	}{
		{
			name:                 "test isset error message " + validatorNamePasswordStrength + " validator",
			PasswordStrengthArgs: *newPasswordStrengthArgs("test", strength.Strong),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultIssetErrorCount,
				testhelper.DefaultErrorMessage,
			),
		},
		{
			name:                 "test blank error message " + validatorNamePasswordStrength + " validator",
			PasswordStrengthArgs: *newPasswordStrengthArgs("The_Error12", strength.Strong),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultNotIssetErrorCount,
				testhelper.BlankString,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			(&validatorprocess.PasswordStrengthValidator{}).Process(
				test.PasswordStrengthArgs.constraint,
				test.PasswordStrengthArgs.value,
				test.PasswordStrengthArgs.exception,
			)

			if len(test.exception.Violations()) != test.resultValidator.Count() {
				t.Errorf(
					validatorNamePasswordStrength+" error count exceptions %v, want %v",
					len(test.exception.Violations()),
					test.resultValidator.Count(),
				)
			}

			for _, viol := range test.exception.Violations() {
				if viol.Message().Message() == test.resultValidator.Message() {
					continue
				}

				t.Errorf(
					validatorNamePasswordStrength+" error message exceptions %v, want %v",
					viol.Message().Message(),
					test.resultValidator.Message(),
				)
			}
		})
	}
}
