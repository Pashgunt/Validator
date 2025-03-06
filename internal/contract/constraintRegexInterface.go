package contract

import (
	"regexp"
)

type ConstraintRegexInterface interface {
	Pattern() regexp.Regexp
	ConstraintInterface
}
