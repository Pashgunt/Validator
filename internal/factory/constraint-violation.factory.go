package factory

import (
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/violation"
	"github.com/Pashgunt/Validator/pkg/interface"
	"reflect"
)

func ConstraintViolationFactory(
	constraint contract.ConstraintMainDataInterface,
	value interface{},
	messageMethod string,
) pkginterface.ConstraintViolationInterface {
	constraintViolation := violation.ConstraintViolation{}
	constraintViolation.SetValue(value)
	constraintViolation.SetPropertyPathError(constraint.PropertyPath())
	constraintViolation.SetRootError(constraint.Root())

	constraintViolation.SetMessage(ConstraintViolationMessageFactory(constraint, messageMethod))

	return &constraintViolation
}

func ConstraintViolationMessageFactory(
	constraint contract.ConstraintMainDataInterface,
	messageMethod string,
) pkginterface.ConstraintViolationMessageInterface {
	constraintViolationMessage := violation.ConstraintViolationMessage{}
	message := reflect.ValueOf(constraint).MethodByName(messageMethod).Call(nil)[0].Interface().(string)
	constraintViolationMessage.SetMessage(message)
	constraintViolationMessage.SetTemplate(message)

	return &constraintViolationMessage
}
