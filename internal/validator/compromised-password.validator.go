package validatorprocess

import (
	"bufio"
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/enum"
	"github.com/Pashgunt/Validator/internal/factory"
	stringhelper "github.com/Pashgunt/Validator/internal/helper/string"
	"github.com/Pashgunt/Validator/pkg/interface"
	"net/http"
	"strings"
)

const (
	defaultApiEndpoint    = "https://api.pwnedpasswords.com/range/"
	delimiterHashResponse = ":"
	firstKIndex           = 5
	indexForPart          = 0
	blankParts
)

type CompromisedPasswordValidator struct {
}

func NewCompromisedPasswordValidator() contract.Validator {
	return &CompromisedPasswordValidator{}
}

func (v *CompromisedPasswordValidator) Process(
	constraint contract.ConstraintInterface,
	value interface{},
	exception pkginterface.ValidationFailedExceptionInterface,
) {
	checkPassword := func(password string) bool {
		hash := strings.ToUpper(stringhelper.HashPassword(password))
		response, err := http.Get(defaultApiEndpoint + hash[:firstKIndex])

		if err != nil {
			return false
		}

		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			return false
		}

		scanner := bufio.NewScanner(response.Body)

		for scanner.Scan() {
			parts := strings.Split(scanner.Text(), delimiterHashResponse)

			if len(parts) > blankParts && parts[indexForPart] == hash[firstKIndex:] {
				return true
			}
		}

		if err = scanner.Err(); err != nil {
			return false
		}

		return false
	}

	if !checkPassword(value.(string)) {
		return
	}

	exception.AppendMessageGeneral(constraint.Message())
	exception.AddViolations([]pkginterface.ConstraintViolationInterface{factory.ConstraintViolationFactory(
		constraint,
		value,
		enum.MessageMethod,
	)})
}
