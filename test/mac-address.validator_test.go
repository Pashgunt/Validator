package test

import (
	validator "github.com/Pashgunt/Validator"
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/validator"
	"github.com/Pashgunt/Validator/internal/violation"
	"github.com/Pashgunt/Validator/pkg"
	testhelper "github.com/Pashgunt/Validator/test/helper"
	"testing"
)

const (
	validatorNameMacAddress = "MacAddress"
)

type macAddressArgs struct {
	constraint contract.ConstraintInterface
	value      interface{}
	exception  contract.ValidationFailedExceptionInterface
}

func newMacAddressArgs(value interface{}) *macAddressArgs {
	return &macAddressArgs{
		value:      value,
		constraint: validator.NewRegex(pkg.MacAddress, testhelper.DefaultErrorMessage),
		exception:  &violation.ValidationFailedException{},
	}
}

func TestMacAddressValidator(t *testing.T) {
	tests := []struct {
		name string
		macAddressArgs
		resultValidator testhelper.ResultValidator
	}{
		{
			name:           "test isset error message " + validatorNameMacAddress + " validator",
			macAddressArgs: *newMacAddressArgs("mac address"),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultIssetErrorCount,
				testhelper.DefaultErrorMessage,
			),
		},
		{
			name:           "test blank error message " + validatorNameMacAddress + " validator",
			macAddressArgs: *newMacAddressArgs("af-14-b3-c2-14-45"),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultNotIssetErrorCount,
				testhelper.BlankString,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			(&validatorprocess.RegexValidator{}).Process(
				test.macAddressArgs.constraint,
				test.macAddressArgs.value,
				test.macAddressArgs.exception,
			)

			if len(test.exception.Violations()) != test.resultValidator.Count() {
				t.Errorf(
					validatorNameMacAddress+" error count exceptions %v, want %v",
					len(test.exception.Violations()),
					test.resultValidator.Count(),
				)
			}

			for _, viol := range test.exception.Violations() {
				if viol.Message().Message() == test.resultValidator.Message() {
					continue
				}

				t.Errorf(
					validatorNameMacAddress+" error message exceptions %v, want %v",
					viol.Message().Message(),
					test.resultValidator.Message(),
				)
			}
		})
	}
}
