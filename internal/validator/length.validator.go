package validatorprocess

import (
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/enum"
	"github.com/Pashgunt/Validator/internal/factory"
	"github.com/Pashgunt/Validator/pkg/interface"
	"reflect"
)

type LengthValidator struct {
	lengthConstraintConverted contract.ConstraintLengthInterface
	exception                 pkginterface.ValidationFailedExceptionInterface
	value                     string
}

func NewLengthValidator() contract.Validator {
	return &LengthValidator{}
}

func (l *LengthValidator) Process(
	constraint contract.ConstraintInterface,
	value interface{},
	exception pkginterface.ValidationFailedExceptionInterface,
) {
	l.lengthConstraintConverted = reflect.ValueOf(constraint).Interface().(contract.ConstraintLengthInterface)
	l.exception = exception
	l.value = value.(string)

	if len(l.value) < l.lengthConstraintConverted.Min() {
		l.processMessage(l.lengthConstraintConverted.MinMessage(), enum.MinMessageMethod)
	}

	if len(l.value) > l.lengthConstraintConverted.Max() {
		l.processMessage(l.lengthConstraintConverted.MaxMessage(), enum.MaxMessageMethod)
	}
}

func (l *LengthValidator) processMessage(
	message string,
	messageMethod string,
) {
	l.exception.AppendMessageGeneral(message)
	l.exception.AddViolations([]pkginterface.ConstraintViolationInterface{factory.ConstraintViolationFactory(
		l.lengthConstraintConverted,
		l.value,
		messageMethod,
	)})
}
