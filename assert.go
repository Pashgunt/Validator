package validator

import (
	"github.com/Pashgunt/Validator/internal/contract"
	validatorprocess "github.com/Pashgunt/Validator/internal/validator"
	"regexp"
)

type RegexConstraint struct {
	pattern           *regexp.Regexp
	message           string
	propertyPath      string
	root              string
	processValidators []contract.Validator
}

func (r *RegexConstraint) PropertyPath() string {
	return r.propertyPath
}

func (r *RegexConstraint) SetPropertyPath(propertyPath string) {
	r.propertyPath = propertyPath
}

func (r *RegexConstraint) Root() string {
	return r.root
}

func (r *RegexConstraint) SetRoot(root string) {
	r.root = root
}

func (r *RegexConstraint) ProcessValidators() []contract.Validator {
	return r.processValidators
}

func (r *RegexConstraint) Pattern() regexp.Regexp {
	return *r.pattern
}

func (r *RegexConstraint) Message() string {
	return r.message
}

func NewRegex(
	pattern string,
	message string,
) *RegexConstraint {
	regex := &RegexConstraint{
		pattern: regexp.MustCompile(pattern),
		message: message,
	}

	regex.processValidators = []contract.Validator{validatorprocess.NewRegexValidator()}

	return regex
}
