package validator

import (
	"github.com/Pashgunt/Validator/internal/contract"
	validatorprocess "github.com/Pashgunt/Validator/internal/validator"
	"reflect"
	"regexp"
)

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

type RegexConstraint struct {
	pattern *regexp.Regexp
	baseConstraint
}

func (r *RegexConstraint) Pattern() regexp.Regexp {
	return *r.pattern
}

func NewRegex(
	pattern string,
	message string,
) *RegexConstraint {
	regex := &RegexConstraint{
		pattern: regexp.MustCompile(pattern),
		baseConstraint: baseConstraint{
			message:           message,
			processValidators: []contract.Validator{validatorprocess.NewRegexValidator()},
		},
	}

	return regex
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

type LengthConstraint struct {
	min, max               int
	minMessage, maxMessage string
	baseConstraint
}

func NewLength(min int, max int, minMessage string, maxMessage string) *LengthConstraint {
	return &LengthConstraint{
		min:        min,
		max:        max,
		minMessage: minMessage,
		maxMessage: maxMessage,
		baseConstraint: baseConstraint{
			processValidators: []contract.Validator{validatorprocess.NewLengthValidator()},
		},
	}
}

func (l *LengthConstraint) Min() int {
	return l.min
}

func (l *LengthConstraint) Max() int {
	return l.max
}

func (l *LengthConstraint) MinMessage() string {
	return l.minMessage
}

func (l *LengthConstraint) MaxMessage() string {
	return l.maxMessage
}

type UrlConstraint struct {
	RegexConstraint
}

func NewUrl(message string) *UrlConstraint {
	return &UrlConstraint{
		RegexConstraint: *NewRegex(Url, message),
	}
}

type EmailConstraint struct {
	RegexConstraint
}

func NewEmail(message string) *EmailConstraint {
	return &EmailConstraint{
		RegexConstraint: *NewRegex(Email, message),
	}
}

type MacAddressConstraint struct {
	RegexConstraint
}

func NewMacAddress(message string) *MacAddressConstraint {
	return &MacAddressConstraint{
		RegexConstraint: *NewRegex(MacAddress, message),
	}
}

type UuidConstraint struct {
	RegexConstraint
}

func NewUuid(message string) *UuidConstraint {
	return &UuidConstraint{
		RegexConstraint: *NewRegex(Uuid, message),
	}
}

type HostnameConstraint struct {
	RegexConstraint
}

func NewHostname(message string) *HostnameConstraint {
	return &HostnameConstraint{
		RegexConstraint: *NewRegex(Hostname, message),
	}
}

type IpConstraint struct {
	RegexConstraint
}

func NewIp(message string) *IpConstraint {
	return &IpConstraint{
		RegexConstraint: *NewRegex(IPv4, message),
	}
}

type PasswordStrengthConstraint struct {
	minScore int
	baseConstraint
}

func (p PasswordStrengthConstraint) MinScore() int {
	return p.minScore
}

func NewPasswordStrength(message string, minSCore int) *PasswordStrengthConstraint {
	return &PasswordStrengthConstraint{
		minScore: minSCore,
		baseConstraint: baseConstraint{
			message:           message,
			processValidators: []contract.Validator{validatorprocess.NewPasswordStrengthValidator()},
		},
	}
}
