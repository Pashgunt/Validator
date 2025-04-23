package validatorprocess

import (
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/enum"
	"github.com/Pashgunt/Validator/internal/factory"
	"github.com/Pashgunt/Validator/pkg/interface"
)

type NotNilValidator struct {
}

func NewNotNilValidator() contract.Validator {
	return &NotNilValidator{}
}

func (v *NotNilValidator) Process(
	constraint contract.ConstraintInterface,
	value interface{},
	exception pkginterface.ValidationFailedExceptionInterface,
) {
	if value != nil {
		return
	}

	exception.AppendMessageGeneral(constraint.Message())
	exception.AddViolations([]pkginterface.ConstraintViolationInterface{factory.ConstraintViolationFactory(
		constraint,
		value,
		enum.MessageMethod,
	)})
}
