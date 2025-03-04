package contract

import (
	"regexp"
)

type ConstraintRegexInterface interface {
	Pattern() regexp.Regexp
	Message() string
	ProcessValidators() []Validator
}

type Validator interface {
	Process(regexConstraint ConstraintRegexInterface, value interface{})
}
