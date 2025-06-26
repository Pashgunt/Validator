package enum

type ConstraintTag string

const (
	KeyAssert string        = "assert"
	NotBlank  ConstraintTag = "not_blank"
	Blank     ConstraintTag = "blank"
	IsFalse   ConstraintTag = "is_false"
	IsTrue    ConstraintTag = "is_true"
	IsType    ConstraintTag = "is_type"
)
