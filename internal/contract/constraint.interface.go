package contract

import "github.com/Pashgunt/Validator/pkg/interface"

type ConstraintInterface interface {
	ProcessValidators() []Validator
	AddProcessValidator(validators []Validator)
	ConstraintMainDataInterface
}

type ConstraintMainDataInterface interface {
	PropertyPath() string
	SetPropertyPath(propertyPath string)
	Root() string
	SetRoot(root string)
	Message() string
	SetMessage(message string)
}

type Validator interface {
	Process(
		regexConstraint ConstraintInterface,
		value interface{},
		exception pkginterface.ValidationFailedExceptionInterface,
	)
}
