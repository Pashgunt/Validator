package validatorprocess

import (
	"fmt"
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/factory"
)

type RegexValidator struct {
}

func NewRegexValidator() *RegexValidator {
	return &RegexValidator{}
}

func (v *RegexValidator) Process(
	regexConstraint contract.ConstraintRegexInterface,
	value interface{},
	exception contract.ValidationFailedExceptionInterface,
) {
	pattern := regexConstraint.Pattern()

	if pattern.MatchString(fmt.Sprintf("%v", value)) {
		return
	}

	exception.AppendMessageGeneral(regexConstraint.Message())
	exception.AddViolations([]contract.ConstraintViolationInterface{factory.ConstraintViolationFactory(
		regexConstraint,
		value,
	)})
}
