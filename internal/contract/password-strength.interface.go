package contract

type PasswordStrengthInterface interface {
	MinScore() int
	ConstraintInterface
}
