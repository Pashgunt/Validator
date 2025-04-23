package factory

import (
	"github.com/Pashgunt/Validator/internal/contract"
	validatorprocess "github.com/Pashgunt/Validator/internal/validator"
	"github.com/Pashgunt/Validator/pkg"
	"github.com/Pashgunt/Validator/pkg/domain"
	pkginterface "github.com/Pashgunt/Validator/pkg/interface"
	"regexp"
)

func NewRegex(
	pattern string,
	message string,
) contract.ConstraintRegexInterface {
	regex := &domain.RegexConstraint{
		ConstraintInterface: NewBaseConstraint(
			message,
			[]contract.Validator{validatorprocess.NewRegexValidator()},
		),
	}
	regex.SetPattern(regexp.MustCompile(pattern))

	return regex
}

func NewLength(min int, max int, minMessage string, maxMessage string) contract.ConstraintLengthInterface {
	return &domain.LengthConstraint{
		MinConstraintLengthInterface: NewMinRangeConstraint(min, minMessage),
		MaxConstraintLengthInterface: NewMaxRangeConstraint(max, maxMessage),
		ConstraintInterface: NewBaseConstraint(
			"",
			[]contract.Validator{validatorprocess.NewLengthValidator()},
		),
	}
}

func NewSpecialRegex(regexType pkg.RegexType, message string) contract.ConstraintRegexInterface {
	switch regexType {
	case pkg.Email:
		return &domain.EmailConstraint{
			ConstraintRegexInterface: NewRegex(pkg.RegexTypeAssoc[regexType], message),
		}
	case pkg.Url:
		return &domain.UrlConstraint{
			ConstraintRegexInterface: NewRegex(pkg.RegexTypeAssoc[regexType], message),
		}
	case pkg.MacAddress:
		return &domain.MacAddressConstraint{
			ConstraintRegexInterface: NewRegex(pkg.RegexTypeAssoc[regexType], message),
		}
	case pkg.Uuid:
		return &domain.UuidConstraint{
			ConstraintRegexInterface: NewRegex(pkg.RegexTypeAssoc[regexType], message),
		}
	case pkg.Hostname:
		return &domain.HostnameConstraint{
			ConstraintRegexInterface: NewRegex(pkg.RegexTypeAssoc[regexType], message),
		}
	case pkg.IPv4:
		return &domain.IpConstraint{
			ConstraintRegexInterface: NewRegex(pkg.RegexTypeAssoc[regexType], message),
		}
	}

	return nil
}

func NewPasswordStrength(message string, minScore int) contract.PasswordStrengthInterface {
	passwordStrength := &domain.PasswordStrengthConstraint{
		ConstraintInterface: NewBaseConstraint(
			message,
			[]contract.Validator{validatorprocess.NewPasswordStrengthValidator()},
		),
	}
	passwordStrength.SetMinScore(minScore)

	return passwordStrength
}

func NewWordCount(min int, max int, minMessage string, maxMessage string) contract.ConstraintLengthInterface {
	return &domain.WordCountConstraint{
		MinConstraintLengthInterface: NewMinRangeConstraint(min, minMessage),
		MaxConstraintLengthInterface: NewMaxRangeConstraint(max, maxMessage),
		ConstraintInterface: NewBaseConstraint(
			"",
			[]contract.Validator{validatorprocess.NewWordCountValidator()},
		),
	}
}

func NewNotCompromisedPassword(message string) contract.ConstraintInterface {
	return &domain.NotCompromisedPasswordConstraint{
		ConstraintInterface: NewBaseConstraint(
			message,
			[]contract.Validator{validatorprocess.NewCompromisedPasswordValidator()},
		),
	}
}

func NewSpoof(message string) contract.ConstraintInterface {
	return &domain.SpoofConstraint{
		ConstraintInterface: NewBaseConstraint(
			message,
			[]contract.Validator{validatorprocess.NewSpoofValidator()},
		),
	}
}

func NewUserPassword(
	message string,
	passwordHasher pkginterface.PasswordHasherInterface,
) contract.UserPasswordInterface {
	userPassword := &domain.UserPasswordConstraint{
		ConstraintInterface: NewBaseConstraint(
			message,
			[]contract.Validator{validatorprocess.NewUserPasswordValidator()},
		),
	}
	userPassword.SetPasswordHasher(passwordHasher)

	return userPassword
}
