package contract

type ComparisonInterface interface {
	Value() int
	SetValue(value int)
	ConstraintInterface
}
