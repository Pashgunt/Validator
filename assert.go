package validator

import (
	"github.com/Pashgunt/Validator/internal/contract"
	validatorprocess "github.com/Pashgunt/Validator/internal/validator"
	"reflect"
	"regexp"
)

type RegexConstraint struct {
	pattern           *regexp.Regexp
	message           string
	propertyPath      string
	root              string
	processValidators []contract.Validator
}

func (r *RegexConstraint) PropertyPath() string {
	return r.propertyPath
}

func (r *RegexConstraint) SetPropertyPath(propertyPath string) {
	r.propertyPath = propertyPath
}

func (r *RegexConstraint) Root() string {
	return r.root
}

func (r *RegexConstraint) SetRoot(root string) {
	r.root = root
}

func (r *RegexConstraint) ProcessValidators() []contract.Validator {
	return r.processValidators
}

func (r *RegexConstraint) Pattern() regexp.Regexp {
	return *r.pattern
}

func (r *RegexConstraint) Message() string {
	return r.message
}

func NewRegex(
	pattern string,
	message string,
) *RegexConstraint {
	regex := &RegexConstraint{
		pattern: regexp.MustCompile(pattern),
		message: message,
	}

	regex.processValidators = []contract.Validator{validatorprocess.NewRegexValidator()}

	return regex
}

type NotBlankConstraint struct {
	message           string
	propertyPath      string
	root              string
	processValidators []contract.Validator
}

func (n *NotBlankConstraint) Message() string {
	return n.message
}

func (n *NotBlankConstraint) SetMessage(message string) {
	n.message = message
}

func (n *NotBlankConstraint) PropertyPath() string {
	return n.propertyPath
}

func (n *NotBlankConstraint) SetPropertyPath(propertyPath string) {
	n.propertyPath = propertyPath
}

func (n *NotBlankConstraint) Root() string {
	return n.root
}

func (n *NotBlankConstraint) SetRoot(root string) {
	n.root = root
}

func (n *NotBlankConstraint) ProcessValidators() []contract.Validator {
	return n.processValidators
}

func NewNotBlank(message string) *NotBlankConstraint {
	notBlank := &NotBlankConstraint{message: message}
	notBlank.processValidators = []contract.Validator{
		validatorprocess.NewNotBlankValidator(),
		validatorprocess.NewNotNilValidator(),
	}

	return notBlank
}

type BlankConstraint struct {
	message           string
	propertyPath      string
	root              string
	processValidators []contract.Validator
}

func (n *BlankConstraint) Message() string {
	return n.message
}

func (n *BlankConstraint) SetMessage(message string) {
	n.message = message
}

func (n *BlankConstraint) PropertyPath() string {
	return n.propertyPath
}

func (n *BlankConstraint) SetPropertyPath(propertyPath string) {
	n.propertyPath = propertyPath
}

func (n *BlankConstraint) Root() string {
	return n.root
}

func (n *BlankConstraint) SetRoot(root string) {
	n.root = root
}

func (n *BlankConstraint) ProcessValidators() []contract.Validator {
	return n.processValidators
}

func NewBlank(message string) *BlankConstraint {
	blank := &BlankConstraint{message: message}
	blank.processValidators = []contract.Validator{
		validatorprocess.NewBlankValidator(),
		validatorprocess.NewIsNilValidator(),
	}

	return blank
}

type IsFalseConstraint struct {
	message           string
	propertyPath      string
	root              string
	processValidators []contract.Validator
}

func (n *IsFalseConstraint) Message() string {
	return n.message
}

func (n *IsFalseConstraint) SetMessage(message string) {
	n.message = message
}

func (n *IsFalseConstraint) PropertyPath() string {
	return n.propertyPath
}

func (n *IsFalseConstraint) SetPropertyPath(propertyPath string) {
	n.propertyPath = propertyPath
}

func (n *IsFalseConstraint) Root() string {
	return n.root
}

func (n *IsFalseConstraint) SetRoot(root string) {
	n.root = root
}

func (n *IsFalseConstraint) ProcessValidators() []contract.Validator {
	return n.processValidators
}

func NewIsFalse(message string) *IsFalseConstraint {
	isFalse := &IsFalseConstraint{message: message}
	isFalse.processValidators = []contract.Validator{validatorprocess.NewIsFalseValidator()}

	return isFalse
}

type IsTrueConstraint struct {
	message           string
	propertyPath      string
	root              string
	processValidators []contract.Validator
}

func (n *IsTrueConstraint) Message() string {
	return n.message
}

func (n *IsTrueConstraint) SetMessage(message string) {
	n.message = message
}

func (n *IsTrueConstraint) PropertyPath() string {
	return n.propertyPath
}

func (n *IsTrueConstraint) SetPropertyPath(propertyPath string) {
	n.propertyPath = propertyPath
}

func (n *IsTrueConstraint) Root() string {
	return n.root
}

func (n *IsTrueConstraint) SetRoot(root string) {
	n.root = root
}

func (n *IsTrueConstraint) ProcessValidators() []contract.Validator {
	return n.processValidators
}

func NewIsTrue(message string) *IsTrueConstraint {
	isTrue := &IsTrueConstraint{message: message}
	isTrue.processValidators = []contract.Validator{validatorprocess.NewIsTrueValidator()}

	return isTrue
}

type IsTypeConstraint struct {
	dataType          reflect.Kind
	message           string
	propertyPath      string
	root              string
	processValidators []contract.Validator
}

func (n *IsTypeConstraint) DataType() reflect.Kind {
	return n.dataType
}

func (n *IsTypeConstraint) Message() string {
	return n.message
}

func (n *IsTypeConstraint) SetMessage(message string) {
	n.message = message
}

func (n *IsTypeConstraint) PropertyPath() string {
	return n.propertyPath
}

func (n *IsTypeConstraint) SetPropertyPath(propertyPath string) {
	n.propertyPath = propertyPath
}

func (n *IsTypeConstraint) Root() string {
	return n.root
}

func (n *IsTypeConstraint) SetRoot(root string) {
	n.root = root
}

func (n *IsTypeConstraint) ProcessValidators() []contract.Validator {
	return n.processValidators
}

func NewIsType(dataType reflect.Kind, message string) *IsTypeConstraint {
	isType := &IsTypeConstraint{dataType: dataType, message: message}
	isType.processValidators = []contract.Validator{validatorprocess.NewIsTypeValidator()}

	return isType
}
