package contract

type PasswordStrengthInterface interface {
	MinScore() int
	SetMinScore(minScore int)
	ConstraintInterface
}
