package contract

type ValidationFailedExceptionInterface interface {
	MessageGeneral() string
	SetMessageGeneral(messageGeneral string)
	AppendMessageGeneral(message string)
	Violations() []ConstraintViolationInterface
	SetViolations(violations []ConstraintViolationInterface)
	AddViolations(violations []ConstraintViolationInterface)
}

type ConstraintViolationInterface interface {
	Value() interface{}
	SetValue(value interface{})
	PropertyPath() string
	SetPropertyPath(propertyPath string)
	Root() string
	SetRoot(root string)
	Message() ConstraintViolationMessageInterface
	SetMessage(message ConstraintViolationMessageInterface)
}

type ConstraintViolationMessageInterface interface {
	Template() string
	SetTemplate(template string)
	Message() string
	SetMessage(message string)
}
