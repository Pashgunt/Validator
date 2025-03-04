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
	value        interface{}
	propertyPath string
	root         string
	message      contract.ConstraintViolationMessageInterface
}

func (c *ConstraintViolation) Value() interface{} {
	return c.value
}

func (c *ConstraintViolation) SetValue(value interface{}) {
	c.value = value
}

func (c *ConstraintViolation) PropertyPath() string {
	return c.propertyPath
}

func (c *ConstraintViolation) SetPropertyPath(propertyPath string) {
	c.propertyPath = propertyPath
}

func (c *ConstraintViolation) Root() string {
	return c.root
}

func (c *ConstraintViolation) SetRoot(root string) {
	c.root = root
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
