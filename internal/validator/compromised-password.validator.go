package validatorprocess

import (
	"bufio"
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/factory"
	stringhelper "github.com/Pashgunt/Validator/internal/helper/string"
	"net/http"
	"strings"
)

const (
	DefaultApiEndpoint    = "https://api.pwnedpasswords.com/range/"
	DelimiterHashResponse = ":"
)

type CompromisedPasswordValidator struct {
}

func NewCompromisedPasswordValidator() *CompromisedPasswordValidator {
	return &CompromisedPasswordValidator{}
}

func (v *CompromisedPasswordValidator) Process(
	constraint contract.ConstraintInterface,
	value interface{},
	exception contract.ValidationFailedExceptionInterface,
) {
	checkPassword := func(password string) bool {
		hash := strings.ToUpper(stringhelper.HashPassword(password))
		response, err := http.Get(DefaultApiEndpoint + hash[:5])

		if err != nil {
			return false
		}

		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			return false
		}

		scanner := bufio.NewScanner(response.Body)

		for scanner.Scan() {
			parts := strings.Split(scanner.Text(), DelimiterHashResponse)

			if len(parts) > 0 && parts[0] == hash[5:] {
				return true
			}
		}

		if err = scanner.Err(); err != nil {
			return false
		}

		return false
	}

	if checkPassword(value.(string)) == false {
		return
	}

	exception.AppendMessageGeneral(constraint.Message())
	exception.AddViolations([]contract.ConstraintViolationInterface{factory.ConstraintViolationFactory(
		constraint,
		value,
		"Message",
	)})
}
