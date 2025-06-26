package tag

import (
	"fmt"
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/enum"
	"github.com/Pashgunt/Validator/pkg/factory"
)

func BlankTag(fieldName string, v contract.CacheInterface) {
	message := fmt.Sprintf(enum.ConstraintMessageDefault[enum.Blank], fieldName)

	if ok := v.Exist(string(enum.Blank)); ok {
		constraint := v.Get(string(enum.Blank))
		constraint.SetMessage(message)
	}

	v.Set(string(enum.Blank), factory.NewBlank(message))
}
