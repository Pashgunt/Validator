package domain

import "github.com/Pashgunt/Validator/internal/contract"

type ComparisonBaseConstraint struct {
	value int
	contract.ConstraintInterface
}

func (c *ComparisonBaseConstraint) SetValue(value int) {
	c.value = value
}

func (c *ComparisonBaseConstraint) Value() int {
	return c.value
}

type LessThenConstraint struct {
	contract.ComparisonInterface
}

type GreaterThenConstraint struct {
	contract.ComparisonInterface
}

type LessThanOrEqualConstraint struct {
	contract.ComparisonInterface
}

type GreaterThanOrEqualConstraint struct {
	contract.ComparisonInterface
}

type NotEqualToConstraint struct {
	contract.ComparisonInterface
}

type EqualToConstraint struct {
	contract.ComparisonInterface
}

type RangeConstraint struct {
	contract.MinConstraintLengthInterface
	contract.MaxConstraintLengthInterface
	contract.ConstraintInterface
}
