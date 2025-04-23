package validatorprocess

import (
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/enum"
	"github.com/Pashgunt/Validator/internal/factory"
	"github.com/Pashgunt/Validator/pkg/interface"
	"reflect"
	"strings"
)

type WordCountValidator struct {
	lengthConstraintConverted contract.ConstraintLengthInterface
	exception                 pkginterface.ValidationFailedExceptionInterface
	value                     interface{}
}

func NewWordCountValidator() contract.Validator {
	return &WordCountValidator{}
}

func (l *WordCountValidator) Process(
	constraint contract.ConstraintInterface,
	value interface{},
	exception pkginterface.ValidationFailedExceptionInterface,
) {
	l.lengthConstraintConverted = reflect.ValueOf(constraint).Interface().(contract.ConstraintLengthInterface)
	l.exception = exception
	l.value = value
	countWords := len(strings.Fields(l.value.(string)))

	if countWords < l.lengthConstraintConverted.Min() {
		l.processMessage(l.lengthConstraintConverted.MinMessage(), enum.MinMessageMethod)
	}

	if countWords > l.lengthConstraintConverted.Max() {
		l.processMessage(l.lengthConstraintConverted.MaxMessage(), enum.MaxMessageMethod)
	}
}

func (l *WordCountValidator) processMessage(
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
