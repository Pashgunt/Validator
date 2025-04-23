package validatorprocess

import (
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/enum"
	"github.com/Pashgunt/Validator/internal/factory"
	"github.com/Pashgunt/Validator/pkg/interface"
	"regexp"
	"unicode"
)

const (
	patternInvisibleSymbols = `^[\x20-\x7E]*$`
	isSpoof                 = true
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
		if !regexp.MustCompile(patternInvisibleSymbols).MatchString(value) {
			return isSpoof
		}

		for _, letter := range value {
			if unicode.Is(unicode.Latin, letter) && !unicode.Is(unicode.Latin, letter) {
				return isSpoof
			}
		}

		return !isSpoof
	}

	if !checkSpoofing(value.(string)) {
		return
	}

	exception.AppendMessageGeneral(constraint.Message())
	exception.AddViolations([]pkginterface.ConstraintViolationInterface{factory.ConstraintViolationFactory(
		constraint,
		value,
		enum.MessageMethod,
	)})
}
