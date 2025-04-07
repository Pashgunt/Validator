package validatorprocess

import (
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/factory"
	"reflect"
)

type UserPasswordValidator struct {
	assocStrength map[int]float64
}

func NewUserPasswordValidator() *UserPasswordValidator {
	return &UserPasswordValidator{}
}

func (v *UserPasswordValidator) Process(
	constraint contract.ConstraintInterface,
	value interface{},
	exception contract.ValidationFailedExceptionInterface,
) {
	userPasswordConstraint := reflect.ValueOf(constraint).Interface().(contract.UserPasswordInterface)

	if value.(string) == userPasswordConstraint.PasswordHasher().GetPasswordHash() {
		return
	}

	exception.AppendMessageGeneral(userPasswordConstraint.Message())
	exception.AddViolations([]contract.ConstraintViolationInterface{factory.ConstraintViolationFactory(
		userPasswordConstraint,
		value,
		"Message",
	)})
}
