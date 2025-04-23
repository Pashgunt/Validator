package validatorprocess

import (
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/enum"
	"github.com/Pashgunt/Validator/internal/factory"
	"github.com/Pashgunt/Validator/pkg/interface"
	"reflect"
)

type ComparisonRangeValidator struct {
	comparisonRangeConstraint contract.ConstraintLengthInterface
	exception                 pkginterface.ValidationFailedExceptionInterface
	value                     int
}

func NewComparisonRangeValidator() contract.Validator {
	return &ComparisonRangeValidator{}
}

func (c *ComparisonRangeValidator) Process(
	constraint contract.ConstraintInterface,
	value interface{},
	exception pkginterface.ValidationFailedExceptionInterface,
) {
	c.comparisonRangeConstraint = reflect.ValueOf(constraint).Interface().(contract.ConstraintLengthInterface)
	c.exception = exception
	c.value = value.(int)

	if c.value < c.comparisonRangeConstraint.Min() {
		c.processMessage(c.comparisonRangeConstraint.MinMessage(), enum.MinMessageMethod)
	}

	if c.value > c.comparisonRangeConstraint.Max() {
		c.processMessage(c.comparisonRangeConstraint.MaxMessage(), enum.MaxMessageMethod)
	}
}

func (c *ComparisonRangeValidator) processMessage(
	message string,
	messageMethod string,
) {
	c.exception.AppendMessageGeneral(message)
	c.exception.AddViolations([]pkginterface.ConstraintViolationInterface{factory.ConstraintViolationFactory(
		c.comparisonRangeConstraint,
		c.value,
		messageMethod,
	)})
}
