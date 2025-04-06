package contract

type ConstraintLengthInterface interface {
	Min() int
	Max() int
	MinMessage() string
	MaxMessage() string
	ConstraintInterface
}
