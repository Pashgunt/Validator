package validatorprocess

import (
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/factory"
)

type BlankValidator struct {
}

func NewBlankValidator() *BlankValidator {
	return &BlankValidator{}
}

func (v *BlankValidator) Process(
	constraint contract.ConstraintInterface,
	value interface{},
	exception contract.ValidationFailedExceptionInterface,
) {
	if value == "" {
		return
	}

	exception.AppendMessageGeneral(constraint.Message())
	exception.AddViolations([]contract.ConstraintViolationInterface{factory.ConstraintViolationFactory(
		constraint,
		value,
	)})
}
