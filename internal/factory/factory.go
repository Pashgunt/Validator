package factory

import (
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/violation"
)

func ConstraintViolationFactory(
	constraint contract.ConstraintMainDataInterface,
	value interface{},
) contract.ConstraintViolationInterface {
	constraintViolation := violation.ConstraintViolation{}
	constraintViolation.SetValue(value)
	constraintViolation.SetPropertyPathError(constraint.PropertyPath())
	constraintViolation.SetRootError(constraint.Root())

	constraintViolation.SetMessage(ConstraintViolationMessageFactory(constraint))

	return &constraintViolation
}

func ConstraintViolationMessageFactory(
	constraint contract.ConstraintMainDataInterface,
) contract.ConstraintViolationMessageInterface {
	constraintViolationMessage := violation.ConstraintViolationMessage{}
	constraintViolationMessage.SetMessage(constraint.Message())
	constraintViolationMessage.SetTemplate(constraint.Message())

	return &constraintViolationMessage
}
