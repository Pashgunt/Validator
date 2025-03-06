package contract

import (
	"regexp"
)

type ConstraintRegexInterface interface {
	Pattern() regexp.Regexp
	ConstraintInterface
}

type Validator interface {
	Process(
		regexConstraint ConstraintInterface,
		value interface{},
		exception ValidationFailedExceptionInterface,
	)
}
