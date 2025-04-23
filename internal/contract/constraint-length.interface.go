package contract

type ConstraintLengthInterface interface {
	MinConstraintLengthInterface
	MaxConstraintLengthInterface
	ConstraintInterface
}

type MinConstraintLengthInterface interface {
	Min() int
	SetMin(min int) MinConstraintLengthInterface
	MinMessage() string
	SetMinMessage(minMessage string) MinConstraintLengthInterface
}

type MaxConstraintLengthInterface interface {
	Max() int
	SetMax(max int) MaxConstraintLengthInterface
	MaxMessage() string
	SetMaxMessage(maxMessage string) MaxConstraintLengthInterface
}
