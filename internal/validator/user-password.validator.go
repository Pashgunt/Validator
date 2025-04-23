package validatorprocess

import (
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/enum"
	"github.com/Pashgunt/Validator/internal/factory"
	"github.com/Pashgunt/Validator/pkg/interface"
	"reflect"
)

type UserPasswordValidator struct {
	assocStrength map[int]float64
}

func NewUserPasswordValidator() contract.Validator {
	return &UserPasswordValidator{}
}

func (v *UserPasswordValidator) Process(
	constraint contract.ConstraintInterface,
	value interface{},
	exception pkginterface.ValidationFailedExceptionInterface,
) {
	userPasswordConstraint := reflect.ValueOf(constraint).Interface().(contract.UserPasswordInterface)

	if value.(string) == userPasswordConstraint.PasswordHasher().GetPasswordHash() {
		return
	}

	exception.AppendMessageGeneral(userPasswordConstraint.Message())
	exception.AddViolations([]pkginterface.ConstraintViolationInterface{factory.ConstraintViolationFactory(
		userPasswordConstraint,
		value,
		enum.MessageMethod,
	)})
}
