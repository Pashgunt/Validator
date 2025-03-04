package validatorprocess

import (
	"fmt"
	"github.com/Pashgunt/Validator/internal/contract"
)

type RegexValidator struct {
}

func NewRegexValidator() *RegexValidator {
	return &RegexValidator{}
}

func (v *RegexValidator) Process(regexConstraint contract.ConstraintRegexInterface, value interface{}) {
	pattern := regexConstraint.Pattern()

	if !pattern.MatchString(fmt.Sprintf("%v", value)) {
		fmt.Println(regexConstraint.Message())
	}
}
