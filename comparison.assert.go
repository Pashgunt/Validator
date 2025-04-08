package validator

import (
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/enum"
	validatorprocess "github.com/Pashgunt/Validator/internal/validator"
)

type comparisonBaseConstraint struct {
	value int
	baseConstraint
}

func (c comparisonBaseConstraint) Value() int {
	return c.value
}

type LessThenConstraint struct {
	comparisonBaseConstraint
}

func NewLessThenConstraint(message string, value int) *LessThenConstraint {
	return &LessThenConstraint{
		comparisonBaseConstraint{
			value: value,
			baseConstraint: baseConstraint{
				message:           message,
				processValidators: []contract.Validator{validatorprocess.NewComparisonValidator(enum.LessThen)},
			},
		},
	}
}

type GreaterThenConstraint struct {
	comparisonBaseConstraint
}

func NewGreaterThenConstraint(message string, value int) *GreaterThenConstraint {
	return &GreaterThenConstraint{
		comparisonBaseConstraint{
			value: value,
			baseConstraint: baseConstraint{
				message:           message,
				processValidators: []contract.Validator{validatorprocess.NewComparisonValidator(enum.GreaterThen)},
			},
		},
	}
}

type LessThanOrEqualConstraint struct {
	comparisonBaseConstraint
}

func NewLessThanOrEqualConstraint(message string, value int) *LessThenConstraint {
	return &LessThenConstraint{
		comparisonBaseConstraint{
			value: value,
			baseConstraint: baseConstraint{
				message:           message,
				processValidators: []contract.Validator{validatorprocess.NewComparisonValidator(enum.LessThenOrEqual)},
			},
		},
	}
}

type GreaterThanOrEqualConstraint struct {
	comparisonBaseConstraint
}

func NewGreaterThanOrEqualConstraint(message string, value int) *GreaterThenConstraint {
	return &GreaterThenConstraint{
		comparisonBaseConstraint{
			value: value,
			baseConstraint: baseConstraint{
				message:           message,
				processValidators: []contract.Validator{validatorprocess.NewComparisonValidator(enum.GreaterThenOrEqual)},
			},
		},
	}
}

type NotEqualToConstraint struct {
	comparisonBaseConstraint
}

func NewNotEqualToConstraint(message string, value int) *NotEqualToConstraint {
	return &NotEqualToConstraint{
		comparisonBaseConstraint{
			value: value,
			baseConstraint: baseConstraint{
				message:           message,
				processValidators: []contract.Validator{validatorprocess.NewComparisonValidator(enum.NotEqual)},
			},
		},
	}
}

type EqualToConstraint struct {
	comparisonBaseConstraint
}

func NewEqualToConstraint(message string, value int) *EqualToConstraint {
	return &EqualToConstraint{
		comparisonBaseConstraint{
			value: value,
			baseConstraint: baseConstraint{
				message:           message,
				processValidators: []contract.Validator{validatorprocess.NewComparisonValidator(enum.Equal)},
			},
		},
	}
}

type RangeConstraint struct {
	minBaseConstraint
	maxBaseConstraint
	baseConstraint
}

func NewRangeConstraint(min int, max int, minMessage string, maxMessage string) *RangeConstraint {
	return &RangeConstraint{
		minBaseConstraint: minBaseConstraint{min: min, minMessage: minMessage},
		maxBaseConstraint: maxBaseConstraint{max: max, maxMessage: maxMessage},
		baseConstraint: baseConstraint{
			processValidators: []contract.Validator{validatorprocess.NewComparisonRangeValidator()},
		},
	}
}
