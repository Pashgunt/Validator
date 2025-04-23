package validatorprocess

import (
	"fmt"
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/factory"
	"github.com/Pashgunt/Validator/pkg/interface"
	"reflect"
)

type RegexValidator struct {
}

func NewRegexValidator() contract.Validator {
	return &RegexValidator{}
}

func (v *RegexValidator) Process(
	constraint contract.ConstraintInterface,
	value interface{},
	exception pkginterface.ValidationFailedExceptionInterface,
) {
	regexConstraintConverted := reflect.ValueOf(constraint).Interface().(contract.ConstraintRegexInterface)
	pattern := regexConstraintConverted.Pattern()

	if pattern.MatchString(fmt.Sprintf("%v", value)) {
		return
	}

	exception.AppendMessageGeneral(regexConstraintConverted.Message())
	exception.AddViolations([]pkginterface.ConstraintViolationInterface{factory.ConstraintViolationFactory(
		regexConstraintConverted,
		value,
		"Message",
	)})
}
