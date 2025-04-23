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
	validatorNameUserPassword = "UserPassword"
)

type UserPasswordArgs struct {
	constraint contract.ConstraintInterface
	value      interface{}
	exception  pkginterface.ValidationFailedExceptionInterface
}

func newUserPasswordArgs(value interface{}, passwordHasher pkginterface.PasswordHasherInterface) *UserPasswordArgs {
	return &UserPasswordArgs{
		value:      value,
		constraint: factory.NewUserPassword(testhelper.DefaultErrorMessage, passwordHasher),
		exception:  &violation.ValidationFailedException{},
	}
}

type NativePasswordHasherInterface struct {
}

func (n NativePasswordHasherInterface) GetPasswordHash() string {
	return "PasswordHash"
}

func TestUserPasswordValidator(t *testing.T) {
	tests := []struct {
		name string
		UserPasswordArgs
		resultValidator testhelper.ResultValidator
	}{
		{
			name:             "test isset error message " + validatorNameUserPassword + " validator",
			UserPasswordArgs: *newUserPasswordArgs("test", &NativePasswordHasherInterface{}),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultIssetErrorCount,
				testhelper.DefaultErrorMessage,
			),
		},
		{
			name:             "test blank error message " + validatorNameUserPassword + " validator",
			UserPasswordArgs: *newUserPasswordArgs("PasswordHash", &NativePasswordHasherInterface{}),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultNotIssetErrorCount,
				testhelper.BlankString,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			(&validatorprocess.UserPasswordValidator{}).Process(
				test.UserPasswordArgs.constraint,
				test.UserPasswordArgs.value,
				test.UserPasswordArgs.exception,
			)

			if len(test.exception.Violations()) != test.resultValidator.Count() {
				t.Errorf(
					validatorNameUserPassword+" error count exceptions %v, want %v",
					len(test.exception.Violations()),
					test.resultValidator.Count(),
				)
			}

			for _, viol := range test.exception.Violations() {
				if viol.Message().Message() == test.resultValidator.Message() {
					continue
				}

				t.Errorf(
					validatorNameUserPassword+" error message exceptions %v, want %v",
					viol.Message().Message(),
					test.resultValidator.Message(),
				)
			}
		})
	}
}
