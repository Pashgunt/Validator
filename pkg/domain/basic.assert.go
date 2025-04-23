package domain

import (
	"github.com/Pashgunt/Validator/internal/contract"
	"reflect"
)

type MinBaseConstraint struct {
	min        int
	minMessage string
}

func (m *MinBaseConstraint) SetMin(min int) contract.MinConstraintLengthInterface {
	m.min = min

	return m
}

func (m *MinBaseConstraint) SetMinMessage(minMessage string) contract.MinConstraintLengthInterface {
	m.minMessage = minMessage

	return m
}

func (m *MinBaseConstraint) Min() int {
	return m.min
}

func (m *MinBaseConstraint) MinMessage() string {
	return m.minMessage
}

type MaxBaseConstraint struct {
	max        int
	maxMessage string
}

func (m *MaxBaseConstraint) SetMax(max int) contract.MaxConstraintLengthInterface {
	m.max = max

	return m
}

func (m *MaxBaseConstraint) SetMaxMessage(maxMessage string) contract.MaxConstraintLengthInterface {
	m.maxMessage = maxMessage

	return m
}

func (m *MaxBaseConstraint) MaxMessage() string {
	return m.maxMessage
}

func (m *MaxBaseConstraint) Max() int {
	return m.max
}

type BaseConstraint struct {
	message, propertyPath, root string
	processValidators           []contract.Validator
}

func (b *BaseConstraint) Message() string {
	return b.message
}

func (b *BaseConstraint) SetMessage(message string) {
	b.message = message
}

func (b *BaseConstraint) PropertyPath() string {
	return b.propertyPath
}

func (b *BaseConstraint) SetPropertyPath(propertyPath string) {
	b.propertyPath = propertyPath
}

func (b *BaseConstraint) Root() string {
	return b.root
}

func (b *BaseConstraint) SetRoot(root string) {
	b.root = root
}

func (b *BaseConstraint) ProcessValidators() []contract.Validator {
	return b.processValidators
}

func (b *BaseConstraint) AddProcessValidator(validators []contract.Validator) {
	b.processValidators = append(b.processValidators, validators...)
}

type NotBlankConstraint struct {
	contract.ConstraintInterface
}

type BlankConstraint struct {
	contract.ConstraintInterface
}

type IsFalseConstraint struct {
	contract.ConstraintInterface
}

type IsTrueConstraint struct {
	contract.ConstraintInterface
}

type IsTypeConstraint struct {
	dataType reflect.Kind
	contract.ConstraintInterface
}

func (n *IsTypeConstraint) DataType() reflect.Kind {
	return n.dataType
}

func (n *IsTypeConstraint) SetDataType(kind reflect.Kind) {
	n.dataType = kind
}
