package factory

import (
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/enum"
	validatorprocess "github.com/Pashgunt/Validator/internal/validator"
	"github.com/Pashgunt/Validator/pkg/domain"
)

func NewComparisonConstraint(message string, value int, operator enum.ComparisonOperator) contract.ComparisonInterface {
	comparison := domain.ComparisonBaseConstraint{
		ConstraintInterface: NewBaseConstraint(
			message,
			[]contract.Validator{validatorprocess.NewComparisonValidator(operator)},
		),
	}
	comparison.SetValue(value)

	return &domain.LessThenConstraint{ComparisonInterface: &comparison}
}

func NewRangeConstraint(min int, max int, minMessage string, maxMessage string) contract.ConstraintLengthInterface {
	return &domain.RangeConstraint{
		MinConstraintLengthInterface: NewMinRangeConstraint(min, minMessage),
		MaxConstraintLengthInterface: NewMaxRangeConstraint(max, maxMessage),
		ConstraintInterface: NewBaseConstraint(
			"",
			[]contract.Validator{validatorprocess.NewComparisonRangeValidator()},
		),
	}
}
