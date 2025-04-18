package validatorprocess

import (
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/factory"
	"reflect"
)

type IsTrueValidator struct {
}

func NewIsTrueValidator() *IsTrueValidator {
	return &IsTrueValidator{}
}

func (v *IsTrueValidator) Process(
	constraint contract.ConstraintInterface,
	value interface{},
	exception contract.ValidationFailedExceptionInterface,
) {
	if reflect.ValueOf(value).Kind() != reflect.Bool {
		return
	}

	if value == true {
		return
	}

	exception.AppendMessageGeneral(constraint.Message())
	exception.AddViolations([]contract.ConstraintViolationInterface{factory.ConstraintViolationFactory(
		constraint,
		value,
		"Message",
	)})
}
