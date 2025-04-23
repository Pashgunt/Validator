package domain

import (
	"github.com/Pashgunt/Validator/internal/contract"
	pkginterface "github.com/Pashgunt/Validator/pkg/interface"
	"regexp"
)

type RegexConstraint struct {
	pattern *regexp.Regexp
	contract.ConstraintInterface
}

func (r *RegexConstraint) SetPattern(pattern *regexp.Regexp) {
	r.pattern = pattern
}

func (r *RegexConstraint) Pattern() *regexp.Regexp {
	return r.pattern
}

type LengthConstraint struct {
	contract.MinConstraintLengthInterface
	contract.MaxConstraintLengthInterface
	contract.ConstraintInterface
}

type UrlConstraint struct {
	contract.ConstraintRegexInterface
}

type EmailConstraint struct {
	contract.ConstraintRegexInterface
}

type MacAddressConstraint struct {
	contract.ConstraintRegexInterface
}

type UuidConstraint struct {
	contract.ConstraintRegexInterface
}

type HostnameConstraint struct {
	contract.ConstraintRegexInterface
}

type IpConstraint struct {
	contract.ConstraintRegexInterface
}

type PasswordStrengthConstraint struct {
	minScore int
	contract.ConstraintInterface
}

func (p *PasswordStrengthConstraint) SetMinScore(minScore int) {
	p.minScore = minScore
}

func (p *PasswordStrengthConstraint) MinScore() int {
	return p.minScore
}

type WordCountConstraint struct {
	contract.MinConstraintLengthInterface
	contract.MaxConstraintLengthInterface
	contract.ConstraintInterface
}

type NotCompromisedPasswordConstraint struct {
	contract.ConstraintInterface
}

type SpoofConstraint struct {
	contract.ConstraintInterface
}

type UserPasswordConstraint struct {
	contract.ConstraintInterface
	passwordHasherInterface pkginterface.PasswordHasherInterface
}

func (u *UserPasswordConstraint) SetPasswordHasher(passwordHasherInterface pkginterface.PasswordHasherInterface) {
	u.passwordHasherInterface = passwordHasherInterface
}

func (u *UserPasswordConstraint) PasswordHasher() pkginterface.PasswordHasherInterface {
	return u.passwordHasherInterface
}
