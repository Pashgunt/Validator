package validatorprocess

import (
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/enum"
	"github.com/Pashgunt/Validator/internal/factory"
	"github.com/Pashgunt/Validator/pkg/interface"
	"reflect"
)

const (
	isTrueValue = true
)

type IsTrueValidator struct {
}

func NewIsTrueValidator() contract.Validator {
	return &IsTrueValidator{}
}

func (v *IsTrueValidator) Process(
	constraint contract.ConstraintInterface,
	value interface{},
	exception pkginterface.ValidationFailedExceptionInterface,
) {
	if reflect.ValueOf(value).Kind() != reflect.Bool {
		return
	}

	if value == isTrueValue {
		return
	}

	exception.AppendMessageGeneral(constraint.Message())
	exception.AddViolations([]pkginterface.ConstraintViolationInterface{factory.ConstraintViolationFactory(
		constraint,
		value,
		enum.MessageMethod,
	)})
}
