package enum

type ComparisonOperator string

const (
	LessThen           ComparisonOperator = "lt"
	GreaterThen        ComparisonOperator = "gt"
	LessThenOrEqual    ComparisonOperator = "lte"
	GreaterThenOrEqual ComparisonOperator = "gte"
	NotEqual           ComparisonOperator = "neq"
	Equal              ComparisonOperator = "eq"
)
