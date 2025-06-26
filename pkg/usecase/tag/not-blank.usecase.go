package tag

import (
	"fmt"
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/enum"
	"github.com/Pashgunt/Validator/pkg/factory"
)

func NotBlankTag(fieldName string, v contract.CacheInterface) {
	message := fmt.Sprintf(enum.ConstraintMessageDefault[enum.NotBlank], fieldName)

	if ok := v.Exist(string(enum.NotBlank)); ok {
		constraint := v.Get(string(enum.NotBlank))
		constraint.SetMessage(message)
	}

	v.Set(string(enum.NotBlank), factory.NewNotBlank(message))
}
