package contract

import "github.com/Pashgunt/Validator/pkg"

type UserPasswordInterface interface {
	PasswordHasher() pkg.PasswordHasherInterface
	ConstraintInterface
}
