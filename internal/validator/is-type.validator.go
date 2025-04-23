package validatorprocess

import (
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/factory"
	"github.com/Pashgunt/Validator/pkg/interface"
	"reflect"
)

type IsTypeValidator struct {
}

func NewIsTypeValidator() contract.Validator {
	return &IsTypeValidator{}
}

func (v *IsTypeValidator) Process(
	constraint contract.ConstraintInterface,
	value interface{},
	exception pkginterface.ValidationFailedExceptionInterface,
) {
	constraintIsType := reflect.ValueOf(constraint).Interface().(contract.ConstraintIsTypeInterface)

	if reflect.ValueOf(value).Kind() == constraintIsType.DataType() {
		return
	}

	exception.AppendMessageGeneral(constraintIsType.Message())
	exception.AddViolations([]pkginterface.ConstraintViolationInterface{factory.ConstraintViolationFactory(
		constraintIsType,
		value,
		"Message",
	)})
}
