package validatorprocess

import (
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/factory"
	"github.com/Pashgunt/Validator/pkg/interface"
	"regexp"
	"unicode"
)

const (
	PatternInvisibleSymbols = `^[\x20-\x7E]*$`
)

type SpoofValidator struct {
}

func NewSpoofValidator() contract.Validator {
	return &SpoofValidator{}
}

func (v *SpoofValidator) Process(
	constraint contract.ConstraintInterface,
	value interface{},
	exception pkginterface.ValidationFailedExceptionInterface,
) {
	checkSpoofing := func(value string) bool {
		if !regexp.MustCompile(PatternInvisibleSymbols).MatchString(value) {
			return true
		}

		for _, letter := range value {
			if unicode.Is(unicode.Latin, letter) && !unicode.Is(unicode.Latin, letter) {
				return true
			}
		}

		return false
	}

	if checkSpoofing(value.(string)) == false {
		return
	}

	exception.AppendMessageGeneral(constraint.Message())
	exception.AddViolations([]pkginterface.ConstraintViolationInterface{factory.ConstraintViolationFactory(
		constraint,
		value,
		"Message",
	)})
}
