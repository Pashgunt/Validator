package validatorprocess

import (
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/enum"
	"github.com/Pashgunt/Validator/internal/factory"
	"reflect"
)

type ComparisonValidator struct {
	comparisonOperator enum.ComparisonOperator
	value              int
	constraint         contract.ComparisonInterface
}

func NewComparisonValidator(comparisonOperator enum.ComparisonOperator) *ComparisonValidator {
	return &ComparisonValidator{comparisonOperator: comparisonOperator}
}

func (c *ComparisonValidator) Process(
	constraint contract.ConstraintInterface,
	value interface{},
	exception contract.ValidationFailedExceptionInterface,
) {
	c.value = value.(int)
	comparisonConstraint := reflect.ValueOf(constraint).Interface().(contract.ComparisonInterface)
	c.constraint = comparisonConstraint

	if ok := c.correctComparisonOperatorProcess(); ok {
		return
	}

	exception.AppendMessageGeneral(comparisonConstraint.Message())
	exception.AddViolations([]contract.ConstraintViolationInterface{factory.ConstraintViolationFactory(
		comparisonConstraint,
		value,
		"Message",
	)})
}

func (c *ComparisonValidator) correctComparisonOperatorProcess() bool {
	switch c.comparisonOperator {
	case enum.LessThen:
		if c.value < c.constraint.Value() {
			return true
		}
	case enum.GreaterThen:
		if c.value > c.constraint.Value() {
			return true
		}
	case enum.LessThenOrEqual:
		if c.value <= c.constraint.Value() {
			return true
		}
	case enum.GreaterThenOrEqual:
		if c.value >= c.constraint.Value() {
			return true
		}
	case enum.NotEqual:
		if c.value != c.constraint.Value() {
			return true
		}
	case enum.Equal:
		if c.value == c.constraint.Value() {
			return true
		}
	}

	return false
}
