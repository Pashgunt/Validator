package validatorprocess

import (
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/factory"
	"reflect"
)

type LengthValidator struct {
	lengthConstraintConverted contract.ConstraintLengthInterface
	exception                 contract.ValidationFailedExceptionInterface
	value                     interface{}
}

func NewLengthValidator() *LengthValidator {
	return &LengthValidator{}
}

func (l *LengthValidator) Process(
	constraint contract.ConstraintInterface,
	value interface{},
	exception contract.ValidationFailedExceptionInterface,
) {
	l.lengthConstraintConverted = reflect.ValueOf(constraint).Interface().(contract.ConstraintLengthInterface)
	l.exception = exception
	l.value = value

	if len(l.value.(string)) < l.lengthConstraintConverted.Min() {
		l.processMessage(l.lengthConstraintConverted.MinMessage(), "MinMessage")
	}

	if len(l.value.(string)) > l.lengthConstraintConverted.Max() {
		l.processMessage(l.lengthConstraintConverted.MaxMessage(), "MaxMessage")
	}
}

func (l *LengthValidator) processMessage(
	message string,
	messageMethod string,
) {
	l.exception.AppendMessageGeneral(message)
	l.exception.AddViolations([]contract.ConstraintViolationInterface{factory.ConstraintViolationFactory(
		l.lengthConstraintConverted,
		l.value,
		messageMethod,
	)})
}
