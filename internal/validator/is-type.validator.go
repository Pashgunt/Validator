package validatorprocess

import (
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/factory"
	"reflect"
)

type IsTypeValidator struct {
}

func NewIsTypeValidator() *IsTypeValidator {
	return &IsTypeValidator{}
}

func (v *IsTypeValidator) Process(
	constraint contract.ConstraintInterface,
	value interface{},
	exception contract.ValidationFailedExceptionInterface,
) {
	constraintIsType := reflect.ValueOf(constraint).Interface().(contract.ConstraintIsTypeInterface)

	if reflect.ValueOf(value).Kind() == constraintIsType.DataType() {
		return
	}

	exception.AppendMessageGeneral(constraintIsType.Message())
	exception.AddViolations([]contract.ConstraintViolationInterface{factory.ConstraintViolationFactory(
		constraintIsType,
		value,
		"Message",
	)})
}
