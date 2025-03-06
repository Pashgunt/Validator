package validatorprocess

import (
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/factory"
)

type NotNilValidator struct {
}

func NewNotNilValidator() *NotNilValidator {
	return &NotNilValidator{}
}

func (v *NotNilValidator) Process(
	regexConstraint contract.ConstraintInterface,
	value interface{},
	exception contract.ValidationFailedExceptionInterface,
) {
	if value != nil {
		return
	}

	exception.AppendMessageGeneral(regexConstraint.Message())
	exception.AddViolations([]contract.ConstraintViolationInterface{factory.ConstraintViolationFactory(
		regexConstraint,
		value,
	)})
}
