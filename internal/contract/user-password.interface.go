package contract

import "github.com/Pashgunt/Validator/pkg/interface"

type UserPasswordInterface interface {
	PasswordHasher() pkginterface.PasswordHasherInterface
	SetPasswordHasher(passwordHasher pkginterface.PasswordHasherInterface)
	ConstraintInterface
}
