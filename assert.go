package validator

import (
	"github.com/Pashgunt/Validator/internal/contract"
	validatorprocess "github.com/Pashgunt/Validator/internal/validator"
	"reflect"
	"regexp"
)

type RegexConstraint struct {
	pattern                     *regexp.Regexp
	message, propertyPath, root string
	processValidators           []contract.Validator
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
	message, propertyPath, root string
	processValidators           []contract.Validator
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
	message, propertyPath, root string
	processValidators           []contract.Validator
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
	message, propertyPath, root string
	processValidators           []contract.Validator
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
	message, propertyPath, root string
	processValidators           []contract.Validator
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
	dataType                    reflect.Kind
	message, propertyPath, root string
	processValidators           []contract.Validator
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

type LengthConstraint struct {
	min, max                                            int
	minMessage, maxMessage, message, propertyPath, root string
	processValidators                                   []contract.Validator
}

func NewLength(min int, max int, minMessage string, maxMessage string) *LengthConstraint {
	length := &LengthConstraint{min: min, max: max, minMessage: minMessage, maxMessage: maxMessage}
	length.processValidators = []contract.Validator{validatorprocess.NewLengthValidator()}

	return length
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

func (l *LengthConstraint) Message() string {
	return l.message
}

func (l *LengthConstraint) PropertyPath() string {
	return l.propertyPath
}

func (l *LengthConstraint) SetPropertyPath(propertyPath string) {
	l.propertyPath = propertyPath
}

func (l *LengthConstraint) Root() string {
	return l.root
}

func (l *LengthConstraint) SetRoot(root string) {
	l.root = root
}

func (l *LengthConstraint) ProcessValidators() []contract.Validator {
	return l.processValidators
}

type UrlConstraint struct {
	message, propertyPath, root string
	pattern                     *regexp.Regexp
	processValidators           []contract.Validator
}

func (c *UrlConstraint) Pattern() regexp.Regexp {
	return *c.pattern
}

func NewUrl(message string) *UrlConstraint {
	url := &UrlConstraint{message: message, pattern: regexp.MustCompile(Url)}
	url.processValidators = []contract.Validator{validatorprocess.NewRegexValidator()}

	return url
}

func (c *UrlConstraint) Message() string {
	return c.message
}

func (c *UrlConstraint) PropertyPath() string {
	return c.propertyPath
}

func (c *UrlConstraint) SetPropertyPath(propertyPath string) {
	c.propertyPath = propertyPath
}

func (c *UrlConstraint) Root() string {
	return c.root
}

func (c *UrlConstraint) SetRoot(root string) {
	c.root = root
}

func (c *UrlConstraint) ProcessValidators() []contract.Validator {
	return c.processValidators
}

type EmailConstraint struct {
	message, propertyPath, root string
	pattern                     *regexp.Regexp
	processValidators           []contract.Validator
}

func (c *EmailConstraint) Pattern() regexp.Regexp {
	return *c.pattern
}

func NewEmail(message string) *EmailConstraint {
	email := &EmailConstraint{message: message, pattern: regexp.MustCompile(Email)}
	email.processValidators = []contract.Validator{validatorprocess.NewRegexValidator()}

	return email
}

func (c *EmailConstraint) Message() string {
	return c.message
}

func (c *EmailConstraint) PropertyPath() string {
	return c.propertyPath
}

func (c *EmailConstraint) SetPropertyPath(propertyPath string) {
	c.propertyPath = propertyPath
}

func (c *EmailConstraint) Root() string {
	return c.root
}

func (c *EmailConstraint) SetRoot(root string) {
	c.root = root
}

func (c *EmailConstraint) ProcessValidators() []contract.Validator {
	return c.processValidators
}

type MacAddressConstraint struct {
	message, propertyPath, root string
	pattern                     *regexp.Regexp
	processValidators           []contract.Validator
}

func (c *MacAddressConstraint) Pattern() regexp.Regexp {
	return *c.pattern
}

func NewMacAddress(message string) *MacAddressConstraint {
	email := &MacAddressConstraint{message: message, pattern: regexp.MustCompile(MacAddress)}
	email.processValidators = []contract.Validator{validatorprocess.NewRegexValidator()}

	return email
}

func (c *MacAddressConstraint) Message() string {
	return c.message
}

func (c *MacAddressConstraint) PropertyPath() string {
	return c.propertyPath
}

func (c *MacAddressConstraint) SetPropertyPath(propertyPath string) {
	c.propertyPath = propertyPath
}

func (c *MacAddressConstraint) Root() string {
	return c.root
}

func (c *MacAddressConstraint) SetRoot(root string) {
	c.root = root
}

func (c *MacAddressConstraint) ProcessValidators() []contract.Validator {
	return c.processValidators
}

type UuidConstraint struct {
	message, propertyPath, root string
	pattern                     *regexp.Regexp
	processValidators           []contract.Validator
}

func (c *UuidConstraint) Pattern() regexp.Regexp {
	return *c.pattern
}

func NewUuid(message string) *UuidConstraint {
	uuid := &UuidConstraint{message: message, pattern: regexp.MustCompile(Uuid)}
	uuid.processValidators = []contract.Validator{validatorprocess.NewRegexValidator()}

	return uuid
}

func (c *UuidConstraint) Message() string {
	return c.message
}

func (c *UuidConstraint) PropertyPath() string {
	return c.propertyPath
}

func (c *UuidConstraint) SetPropertyPath(propertyPath string) {
	c.propertyPath = propertyPath
}

func (c *UuidConstraint) Root() string {
	return c.root
}

func (c *UuidConstraint) SetRoot(root string) {
	c.root = root
}

func (c *UuidConstraint) ProcessValidators() []contract.Validator {
	return c.processValidators
}

type HostnameConstraint struct {
	message, propertyPath, root string
	pattern                     *regexp.Regexp
	processValidators           []contract.Validator
}

func (c *HostnameConstraint) Pattern() regexp.Regexp {
	return *c.pattern
}

func NewHostname(message string) *HostnameConstraint {
	hostname := &HostnameConstraint{message: message, pattern: regexp.MustCompile(Hostname)}
	hostname.processValidators = []contract.Validator{validatorprocess.NewRegexValidator()}

	return hostname
}

func (c *HostnameConstraint) Message() string {
	return c.message
}

func (c *HostnameConstraint) PropertyPath() string {
	return c.propertyPath
}

func (c *HostnameConstraint) SetPropertyPath(propertyPath string) {
	c.propertyPath = propertyPath
}

func (c *HostnameConstraint) Root() string {
	return c.root
}

func (c *HostnameConstraint) SetRoot(root string) {
	c.root = root
}

func (c *HostnameConstraint) ProcessValidators() []contract.Validator {
	return c.processValidators
}

type IpConstraint struct {
	message, propertyPath, root string
	pattern                     *regexp.Regexp
	processValidators           []contract.Validator
}

func (c *IpConstraint) Pattern() regexp.Regexp {
	return *c.pattern
}

func NewIp(message string) *IpConstraint {
	ip := &IpConstraint{message: message, pattern: regexp.MustCompile(IPv4)}
	ip.processValidators = []contract.Validator{validatorprocess.NewRegexValidator()}

	return ip
}

func (c *IpConstraint) Message() string {
	return c.message
}

func (c *IpConstraint) PropertyPath() string {
	return c.propertyPath
}

func (c *IpConstraint) SetPropertyPath(propertyPath string) {
	c.propertyPath = propertyPath
}

func (c *IpConstraint) Root() string {
	return c.root
}

func (c *IpConstraint) SetRoot(root string) {
	c.root = root
}

func (c *IpConstraint) ProcessValidators() []contract.Validator {
	return c.processValidators
}
