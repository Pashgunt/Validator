package validatorprocess

import (
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/factory"
	stringhelper "github.com/Pashgunt/Validator/internal/helper/string"
	"github.com/Pashgunt/Validator/strength"
	"reflect"
)

type PasswordStrengthValidator struct {
	assocStrength map[int]float64
}

func NewPasswordStrengthValidator() *PasswordStrengthValidator {
	return &PasswordStrengthValidator{}
}

func (v *PasswordStrengthValidator) Process(
	constraint contract.ConstraintInterface,
	value interface{},
	exception contract.ValidationFailedExceptionInterface,
) {
	v.assocStrength = map[int]float64{
		strength.Weak:       20,
		strength.Medium:     40,
		strength.Strong:     60,
		strength.VeryStrong: 80,
	}

	passwordStrengthConstraint := reflect.ValueOf(constraint).Interface().(contract.PasswordStrengthInterface)

	if stringhelper.CalculateEntropy(value.(string)) >= v.assocStrength[passwordStrengthConstraint.MinScore()] {
		return
	}

	exception.AppendMessageGeneral(passwordStrengthConstraint.Message())
	exception.AddViolations([]contract.ConstraintViolationInterface{factory.ConstraintViolationFactory(
		passwordStrengthConstraint,
		value,
		"Message",
	)})
}
