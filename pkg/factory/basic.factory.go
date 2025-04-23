package factory

import (
	"github.com/Pashgunt/Validator/internal/contract"
	validatorprocess "github.com/Pashgunt/Validator/internal/validator"
	"github.com/Pashgunt/Validator/pkg/domain"
	"reflect"
)

func NewBaseConstraint(message string, validators []contract.Validator) contract.ConstraintInterface {
	constraint := domain.BaseConstraint{}
	constraint.SetMessage(message)
	constraint.AddProcessValidator(validators)

	return &constraint
}

func NewMinRangeConstraint(
	min int,
	minMessage string,
) contract.MinConstraintLengthInterface {
	minBaseConstraint := domain.MinBaseConstraint{}
	minBaseConstraint.SetMin(min)
	minBaseConstraint.SetMinMessage(minMessage)

	return &minBaseConstraint
}

func NewMaxRangeConstraint(
	max int,
	maxMessage string,
) contract.MaxConstraintLengthInterface {

	maxBaseConstraint := domain.MaxBaseConstraint{}
	maxBaseConstraint.SetMax(max)
	maxBaseConstraint.SetMaxMessage(maxMessage)

	return &maxBaseConstraint
}

func NewNotBlank(message string) contract.ConstraintInterface {
	return &domain.NotBlankConstraint{ConstraintInterface: NewBaseConstraint(message, []contract.Validator{
		validatorprocess.NewNotBlankValidator(),
		validatorprocess.NewNotNilValidator(),
	})}
}

func NewBlank(message string) contract.ConstraintInterface {
	return &domain.BlankConstraint{ConstraintInterface: NewBaseConstraint(message, []contract.Validator{
		validatorprocess.NewBlankValidator(),
		validatorprocess.NewIsNilValidator(),
	})}
}

func NewIsFalse(message string) contract.ConstraintInterface {
	return &domain.IsFalseConstraint{ConstraintInterface: NewBaseConstraint(message, []contract.Validator{
		validatorprocess.NewIsFalseValidator(),
	})}
}

func NewIsTrue(message string) contract.ConstraintInterface {
	return &domain.IsTrueConstraint{ConstraintInterface: NewBaseConstraint(message, []contract.Validator{
		validatorprocess.NewIsTrueValidator(),
	})}
}

func NewIsType(message string, dataType reflect.Kind) contract.ConstraintInterface {
	isTypeConstraint := &domain.IsTypeConstraint{ConstraintInterface: NewBaseConstraint(message, []contract.Validator{
		validatorprocess.NewIsTypeValidator(),
	})}
	isTypeConstraint.SetDataType(dataType)

	return isTypeConstraint
}
