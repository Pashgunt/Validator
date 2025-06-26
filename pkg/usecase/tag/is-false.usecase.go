package tag

import (
	"fmt"
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/enum"
	"github.com/Pashgunt/Validator/pkg/factory"
)

func IsFalseTag(fieldName string, v contract.CacheInterface) {
	message := fmt.Sprintf(enum.ConstraintMessageDefault[enum.IsFalse], fieldName)

	if ok := v.Exist(string(enum.IsFalse)); ok {
		constraint := v.Get(string(enum.IsFalse))
		constraint.SetMessage(message)
	}

	v.Set(string(enum.IsFalse), factory.NewIsFalse(message))
}
