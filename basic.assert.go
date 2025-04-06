package validator

import (
	"github.com/Pashgunt/Validator/internal/contract"
	validatorprocess "github.com/Pashgunt/Validator/internal/validator"
	"reflect"
)

type minBaseConstraint struct {
	min        int
	minMessage string
}

func (m minBaseConstraint) Min() int {
	return m.min
}

func (m minBaseConstraint) MinMessage() string {
	return m.minMessage
}

type maxBaseConstraint struct {
	max        int
	maxMessage string
}

func (m maxBaseConstraint) MaxMessage() string {
	return m.maxMessage
}

func (m maxBaseConstraint) Max() int {
	return m.max
}

type baseConstraint struct {
	message, propertyPath, root string
	processValidators           []contract.Validator
}

func (b *baseConstraint) Message() string {
	return b.message
}

func (b *baseConstraint) SetMessage(message string) {
	b.message = message
}

func (b *baseConstraint) PropertyPath() string {
	return b.propertyPath
}

func (b *baseConstraint) SetPropertyPath(propertyPath string) {
	b.propertyPath = propertyPath
}

func (b *baseConstraint) Root() string {
	return b.root
}

func (b *baseConstraint) SetRoot(root string) {
	b.root = root
}

func (b *baseConstraint) ProcessValidators() []contract.Validator {
	return b.processValidators
}

type NotBlankConstraint struct {
	baseConstraint
}

func NewNotBlank(message string) *NotBlankConstraint {
	return &NotBlankConstraint{baseConstraint: baseConstraint{
		message: message,
		processValidators: []contract.Validator{
			validatorprocess.NewNotBlankValidator(),
			validatorprocess.NewNotNilValidator(),
		},
	}}
}

type BlankConstraint struct {
	baseConstraint
}

func NewBlank(message string) *BlankConstraint {
	return &BlankConstraint{baseConstraint: baseConstraint{
		message: message,
		processValidators: []contract.Validator{
			validatorprocess.NewBlankValidator(),
			validatorprocess.NewIsNilValidator(),
		},
	}}
}

type IsFalseConstraint struct {
	baseConstraint
}

func NewIsFalse(message string) *IsFalseConstraint {
	return &IsFalseConstraint{baseConstraint: baseConstraint{
		message:           message,
		processValidators: []contract.Validator{validatorprocess.NewIsFalseValidator()},
	}}
}

type IsTrueConstraint struct {
	baseConstraint
}

func NewIsTrue(message string) *IsTrueConstraint {
	return &IsTrueConstraint{baseConstraint: baseConstraint{
		message:           message,
		processValidators: []contract.Validator{validatorprocess.NewIsTrueValidator()},
	}}
}

type IsTypeConstraint struct {
	dataType reflect.Kind
	baseConstraint
}

func (n *IsTypeConstraint) DataType() reflect.Kind {
	return n.dataType
}

func NewIsType(dataType reflect.Kind, message string) *IsTypeConstraint {
	return &IsTypeConstraint{
		dataType: dataType,
		baseConstraint: baseConstraint{
			message:           message,
			processValidators: []contract.Validator{validatorprocess.NewIsTypeValidator()},
		},
	}
}
