package contract

type ConstraintInterface interface {
	ProcessValidators() []Validator
	ConstraintMainDataInterface
}

type ConstraintMainDataInterface interface {
	PropertyPath() string
	SetPropertyPath(propertyPath string)
	Root() string
	SetRoot(root string)
	Message() string
}

type Validator interface {
	Process(
		regexConstraint ConstraintInterface,
		value interface{},
		exception ValidationFailedExceptionInterface,
	)
}
