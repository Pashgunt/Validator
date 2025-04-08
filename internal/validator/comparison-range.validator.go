package validatorprocess

import (
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/factory"
	"reflect"
)

type ComparisonRangeValidator struct {
	comparisonRangeConstraint contract.ConstraintLengthInterface
	exception                 contract.ValidationFailedExceptionInterface
	value                     int
}

func NewComparisonRangeValidator() *ComparisonRangeValidator {
	return &ComparisonRangeValidator{}
}

func (c *ComparisonRangeValidator) Process(
	constraint contract.ConstraintInterface,
	value interface{},
	exception contract.ValidationFailedExceptionInterface,
) {
	c.comparisonRangeConstraint = reflect.ValueOf(constraint).Interface().(contract.ConstraintLengthInterface)
	c.exception = exception
	c.value = value.(int)

	if c.value < c.comparisonRangeConstraint.Min() {
		c.processMessage(c.comparisonRangeConstraint.MinMessage(), "MinMessage")
	}

	if c.value > c.comparisonRangeConstraint.Max() {
		c.processMessage(c.comparisonRangeConstraint.MaxMessage(), "MaxMessage")
	}
}

func (c *ComparisonRangeValidator) processMessage(
	message string,
	messageMethod string,
) {
	c.exception.AppendMessageGeneral(message)
	c.exception.AddViolations([]contract.ConstraintViolationInterface{factory.ConstraintViolationFactory(
		c.comparisonRangeConstraint,
		c.value,
		messageMethod,
	)})
}
