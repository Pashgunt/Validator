package validatorprocess

import (
	"fmt"
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/factory"
	"reflect"
)

type RegexValidator struct {
}

func NewRegexValidator() *RegexValidator {
	return &RegexValidator{}
}

func (v *RegexValidator) Process(
	constraint contract.ConstraintInterface,
	value interface{},
	exception contract.ValidationFailedExceptionInterface,
) {
	regexConstraintConverted := reflect.ValueOf(constraint).Interface().(contract.ConstraintRegexInterface)
	pattern := regexConstraintConverted.Pattern()

	if pattern.MatchString(fmt.Sprintf("%v", value)) {
		return
	}

	exception.AppendMessageGeneral(regexConstraintConverted.Message())
	exception.AddViolations([]contract.ConstraintViolationInterface{factory.ConstraintViolationFactory(
		regexConstraintConverted,
		value,
		"Message",
	)})
}
