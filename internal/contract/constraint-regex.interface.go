package contract

import (
	"regexp"
)

type ConstraintRegexInterface interface {
	Pattern() *regexp.Regexp
	SetPattern(pattern *regexp.Regexp)
	ConstraintInterface
}
