package validatorprocess

import (
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/enum"
	"github.com/Pashgunt/Validator/internal/factory"
	stringhelper "github.com/Pashgunt/Validator/internal/helper/string"
	"github.com/Pashgunt/Validator/pkg"
	"github.com/Pashgunt/Validator/pkg/interface"
	"reflect"
)

const (
	weekAssoc       = 20
	mediumAssoc     = 40
	strongAssoc     = 60
	veryStrongAssoc = 80
)

type PasswordStrengthValidator struct {
	assocStrength map[int]float64
}

func NewPasswordStrengthValidator() contract.Validator {
	return &PasswordStrengthValidator{}
}

func (v *PasswordStrengthValidator) Process(
	constraint contract.ConstraintInterface,
	value interface{},
	exception pkginterface.ValidationFailedExceptionInterface,
) {
	v.assocStrength = map[int]float64{
		pkg.Weak:       weekAssoc,
		pkg.Medium:     mediumAssoc,
		pkg.Strong:     strongAssoc,
		pkg.VeryStrong: veryStrongAssoc,
	}

	passwordStrengthConstraint := reflect.ValueOf(constraint).Interface().(contract.PasswordStrengthInterface)

	if stringhelper.CalculateEntropy(value.(string)) >= v.assocStrength[passwordStrengthConstraint.MinScore()] {
		return
	}

	exception.AppendMessageGeneral(passwordStrengthConstraint.Message())
	exception.AddViolations([]pkginterface.ConstraintViolationInterface{factory.ConstraintViolationFactory(
		passwordStrengthConstraint,
		value,
		enum.MessageMethod,
	)})
}
