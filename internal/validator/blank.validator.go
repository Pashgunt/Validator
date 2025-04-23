package validatorprocess

import (
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/enum"
	"github.com/Pashgunt/Validator/internal/factory"
	"github.com/Pashgunt/Validator/pkg/interface"
)

const (
	blankValue = ""
)

type BlankValidator struct {
}

func NewBlankValidator() contract.Validator {
	return &BlankValidator{}
}

func (v *BlankValidator) Process(
	constraint contract.ConstraintInterface,
	value interface{},
	exception pkginterface.ValidationFailedExceptionInterface,
) {
	if value == blankValue {
		return
	}

	exception.AppendMessageGeneral(constraint.Message())
	exception.AddViolations([]pkginterface.ConstraintViolationInterface{factory.ConstraintViolationFactory(
		constraint,
		value,
		enum.MessageMethod,
	)})
}
