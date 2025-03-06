package validatorprocess

import (
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/factory"
)

type NotBlankValidator struct {
}

func NewNotBlankValidator() *NotBlankValidator {
	return &NotBlankValidator{}
}

func (v *NotBlankValidator) Process(
	regexConstraint contract.ConstraintInterface,
	value interface{},
	exception contract.ValidationFailedExceptionInterface,
) {
	if value != "" {
		return
	}

	exception.AppendMessageGeneral(regexConstraint.Message())
	exception.AddViolations([]contract.ConstraintViolationInterface{factory.ConstraintViolationFactory(
		regexConstraint,
		value,
	)})
}
