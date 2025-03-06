package validatorprocess

import (
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/factory"
	"reflect"
)

type IsFalseValidator struct {
}

func NewIsFalseValidator() *IsFalseValidator {
	return &IsFalseValidator{}
}

func (v *IsFalseValidator) Process(
	constraint contract.ConstraintInterface,
	value interface{},
	exception contract.ValidationFailedExceptionInterface,
) {
	if reflect.ValueOf(value).Kind() != reflect.Bool {
		return
	}

	if value == false {
		return
	}

	exception.AppendMessageGeneral(constraint.Message())
	exception.AddViolations([]contract.ConstraintViolationInterface{factory.ConstraintViolationFactory(
		constraint,
		value,
	)})
}
