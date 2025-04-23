package usecase

import (
	"fmt"
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/enum"
	"github.com/Pashgunt/Validator/pkg/factory"
)

func NotBlankTag(fieldName string, v contract.CacheInterface, tagItem string) {
	message := fmt.Sprintf(enum.ConstraintMessageDefault[enum.NotBlank], fieldName)

	if ok := v.Exist(tagItem); ok {
		constraint := v.Get(tagItem)
		constraint.SetMessage(message)
	}

	v.Set(tagItem, factory.NewNotBlank(message))
}
