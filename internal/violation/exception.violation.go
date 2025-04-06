package violation

import "github.com/Pashgunt/Validator/internal/contract"

type ValidationFailedException struct {
	messageGeneral string
	violations     []contract.ConstraintViolationInterface
}

func (v *ValidationFailedException) MessageGeneral() string {
	return v.messageGeneral
}

func (v *ValidationFailedException) SetMessageGeneral(messageGeneral string) {
	v.messageGeneral = messageGeneral
}

func (v *ValidationFailedException) AppendMessageGeneral(message string) {
	v.messageGeneral += message + "\n"
}

func (v *ValidationFailedException) Violations() []contract.ConstraintViolationInterface {
	return v.violations
}

func (v *ValidationFailedException) SetViolations(violations []contract.ConstraintViolationInterface) {
	v.violations = violations
}

func (v *ValidationFailedException) AddViolations(violations []contract.ConstraintViolationInterface) {
	v.violations = append(v.violations, violations...)
}

type ConstraintViolation struct {
	value             interface{}
	propertyPathError string
	rootError         string
	message           contract.ConstraintViolationMessageInterface
}

func (c *ConstraintViolation) Value() interface{} {
	return c.value
}

func (c *ConstraintViolation) SetValue(value interface{}) {
	c.value = value
}

func (c *ConstraintViolation) PropertyPathError() string {
	return c.propertyPathError
}

func (c *ConstraintViolation) SetPropertyPathError(propertyPath string) {
	c.propertyPathError = propertyPath
}

func (c *ConstraintViolation) RootError() string {
	return c.rootError
}

func (c *ConstraintViolation) SetRootError(root string) {
	c.rootError = root
}

func (c *ConstraintViolation) Message() contract.ConstraintViolationMessageInterface {
	return c.message
}

func (c *ConstraintViolation) SetMessage(message contract.ConstraintViolationMessageInterface) {
	c.message = message
}

type ConstraintViolationMessage struct {
	template string
	message  string
}

func (c *ConstraintViolationMessage) Template() string {
	return c.template
}

func (c *ConstraintViolationMessage) SetTemplate(template string) {
	c.template = template
}

func (c *ConstraintViolationMessage) Message() string {
	return c.message
}

func (c *ConstraintViolationMessage) SetMessage(message string) {
	c.message = message
}
