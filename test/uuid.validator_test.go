package test

import (
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/validator"
	"github.com/Pashgunt/Validator/internal/violation"
	"github.com/Pashgunt/Validator/pkg"
	"github.com/Pashgunt/Validator/pkg/factory"
	"github.com/Pashgunt/Validator/pkg/interface"
	testhelper "github.com/Pashgunt/Validator/test/helper"
	"github.com/google/uuid"
	"testing"
)

const (
	validatorNameUuid = "Uuid"
)

type uuidArgs struct {
	constraint contract.ConstraintInterface
	value      interface{}
	exception  pkginterface.ValidationFailedExceptionInterface
}

func newUuidArgs(value interface{}) *uuidArgs {
	return &uuidArgs{
		value:      value,
		constraint: factory.NewSpecialRegex(pkg.Uuid, testhelper.DefaultErrorMessage),
		exception:  &violation.ValidationFailedException{},
	}
}

func TestUuidValidator(t *testing.T) {
	uuidGen, _ := uuid.NewUUID()

	tests := []struct {
		name string
		uuidArgs
		resultValidator testhelper.ResultValidator
	}{
		{
			name:     "test isset error message " + validatorNameUuid + " validator",
			uuidArgs: *newUuidArgs("uuid"),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultIssetErrorCount,
				testhelper.DefaultErrorMessage,
			),
		},
		{
			name:     "test blank error message " + validatorNameUuid + " validator",
			uuidArgs: *newUuidArgs(uuidGen.String()),
			resultValidator: *testhelper.NewResultValidator(
				testhelper.DefaultNotIssetErrorCount,
				testhelper.BlankString,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			(&validatorprocess.RegexValidator{}).Process(
				test.uuidArgs.constraint,
				test.uuidArgs.value,
				test.uuidArgs.exception,
			)

			if len(test.exception.Violations()) != test.resultValidator.Count() {
				t.Errorf(
					validatorNameUuid+" error count exceptions %v, want %v",
					len(test.exception.Violations()),
					test.resultValidator.Count(),
				)
			}

			for _, viol := range test.exception.Violations() {
				if viol.Message().Message() == test.resultValidator.Message() {
					continue
				}

				t.Errorf(
					validatorNameUuid+" error message exceptions %v, want %v",
					viol.Message().Message(),
					test.resultValidator.Message(),
				)
			}
		})
	}
}
