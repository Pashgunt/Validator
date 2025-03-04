package contract

import (
	"regexp"
)

type ConstraintRegexInterface interface {
	Pattern() regexp.Regexp
	Message() string
	ConstraintInterface
}

type Validator interface {
	Process(
		regexConstraint ConstraintRegexInterface,
		value interface{},
		exception ValidationFailedExceptionInterface,
	)
}
