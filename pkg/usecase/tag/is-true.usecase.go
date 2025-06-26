package tag

import (
	"fmt"
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/enum"
	"github.com/Pashgunt/Validator/pkg/factory"
)

func IsTrueTag(fieldName string, v contract.CacheInterface) {
	message := fmt.Sprintf(enum.ConstraintMessageDefault[enum.IsTrue], fieldName)

	if ok := v.Exist(string(enum.IsTrue)); ok {
		constraint := v.Get(string(enum.IsTrue))
		constraint.SetMessage(message)
	}

	v.Set(string(enum.IsTrue), factory.NewIsTrue(message))
}
