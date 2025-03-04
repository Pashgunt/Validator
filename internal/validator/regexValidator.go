package validatorprocess

import (
	"fmt"
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/violation"
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

	constraintViolation := violation.ConstraintViolation{}
	constraintViolation.SetValue(value)
	constraintViolation.SetPropertyPath(regexConstraint.PropertyPath())
	constraintViolation.SetRoot(regexConstraint.Root())

	constraintViolationMessage := violation.ConstraintViolationMessage{}
	constraintViolationMessage.SetMessage(regexConstraint.Message())
	constraintViolationMessage.SetTemplate(regexConstraint.Message())

	constraintViolation.SetMessage(&constraintViolationMessage)
	exception.AppendMessageGeneral(regexConstraint.Message())
	exception.AddViolations([]contract.ConstraintViolationInterface{&constraintViolation})
}
