package tag

import (
	"errors"
	"fmt"
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/enum"
	"github.com/Pashgunt/Validator/pkg/factory"
	"reflect"
	"strings"
)

func IsTypeTag(fieldName string, v contract.CacheInterface, tagItem string) {
	message := fmt.Sprintf(enum.ConstraintMessageDefault[enum.IsType], fieldName)

	if ok := v.Exist(string(enum.IsType)); ok {
		constraint := v.Get(string(enum.IsType))
		constraint.SetMessage(message)
	}

	v.Set(string(enum.IsType), factory.NewIsType(message, getDataType(tagItem)))
}

func getDataType(tagItem string) reflect.Kind {
	dataType := strings.Split(tagItem, enum.AssertParamDelimiter)
	res := dataType[len(dataType)-1]

	if res == "" {
		panic(errors.New("<UNK>"))
	}

	return map[string]reflect.Kind{
		"Bool":       reflect.Bool,
		"Int":        reflect.Int,
		"Int8":       reflect.Int8,
		"Int16":      reflect.Int16,
		"Int32":      reflect.Int32,
		"Int64":      reflect.Int64,
		"Uint":       reflect.Uint,
		"Uint8":      reflect.Uint8,
		"Uint16":     reflect.Uint16,
		"Uint32":     reflect.Uint32,
		"Uint64":     reflect.Uint64,
		"Float32":    reflect.Float32,
		"Float64":    reflect.Float64,
		"Complex64":  reflect.Complex64,
		"Complex128": reflect.Complex128,
		"String":     reflect.String,
	}[res]
}
