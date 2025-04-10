package validator

import (
	"github.com/Pashgunt/Validator/internal/cache"
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/enum"
	maphelper "github.com/Pashgunt/Validator/internal/helper/map"
	structhelper "github.com/Pashgunt/Validator/internal/helper/struct"
	"github.com/Pashgunt/Validator/internal/violation"
	"reflect"
	"strings"
)

type ValidatorInterface interface {
	Validate(value interface{}, constraints Collection)
	ValidateValue(value interface{}, constraints AssertListValue)
}

type ValidatorExceptionInterface interface {
	Exception() contract.ValidationFailedExceptionInterface
}

type SimpleValidator struct {
	exception contract.ValidationFailedExceptionInterface
	cache     contract.CacheInterface
}

func (v *SimpleValidator) Exception() contract.ValidationFailedExceptionInterface {
	return v.exception
}

func NewSimpleValidator() *SimpleValidator {
	return &SimpleValidator{
		exception: &violation.ValidationFailedException{},
		cache:     cache.NewCache(),
	}
}

func (v *SimpleValidator) Validate(value interface{}, constraints *Collection) {
	if constraints == nil {
		reflectValue := reflect.ValueOf(value)

		if reflectValue.Type().Kind() != reflect.Struct {
			panic("ожидалась структура")
		}

		getOrCreateValidators := func(tag string) AssertListValue {
			tagSlice := strings.Split(tag, "|")
			var validators AssertListValue

			for _, tagItem := range tagSlice {
				switch tagItem {
				case "not_blank":
					if ok := v.cache.Exist(tagItem); ok {
						validators = append(validators, v.cache.Get(tagItem))
						continue
					}

					notBlank := NewNotBlank("TEST")
					v.cache.Set(tagItem, notBlank)
					validators = append(validators, notBlank)
				}
			}

			return validators
		}

		for i := 0; i < reflectValue.NumField(); i++ {
			tag := reflectValue.Type().Field(i).Tag.Get("assert")

			if tag == "" {
				continue
			}

			v.doProcessConstraintValidate(
				getOrCreateValidators(tag),
				reflectValue.Field(i).Interface(),
				reflectValue.String(),
				reflectValue.Type().Field(i).Name,
			)
		}

		return
	}

	reflectValue := reflect.ValueOf(value)

	switch reflectValue.Kind() {
	case reflect.Ptr:
		reflectValue = reflectValue.Elem()
		v.processInitValidate(
			structhelper.GetFillValidateData(reflectValue),
			*constraints,
			reflectValue.String(),
		)
		break
	case reflect.Struct:
		v.processInitValidate(
			structhelper.GetFillValidateData(reflectValue),
			*constraints,
			reflectValue.String(),
		)
		break
	case reflect.Map:
		v.processInitValidate(
			maphelper.GetFillValidateData(reflectValue),
			*constraints,
			reflectValue.String(),
		)
		break
	default:
		panic("Unsupported type for validate: " + reflectValue.Kind().String())
	}
}

func (v *SimpleValidator) ValidateValue(value interface{}, constraints AssertListValue) {
	reflectValue := reflect.ValueOf(value)

	if reflectValue.Kind() == reflect.Array || reflectValue.Kind() == reflect.Slice {
		for i := 0; i < reflectValue.Len(); i++ {
			v.ValidateValue(reflectValue.Index(i).Interface(), constraints)
		}

		return
	}

	v.doProcessConstraintValidate(constraints, value, enum.ValueRootAnonymous, enum.ValuePropertyAnonymous)
}

func (v *SimpleValidator) processInitValidate(
	data map[string]interface{},
	constraints Collection,
	root string,
) {
	for property, propertyValue := range data {
		constraintList, isset := constraints.Asserts()[property]

		if !isset {
			continue
		}

		reflectPropertyValue := reflect.ValueOf(propertyValue)

		if reflectPropertyValue.Kind() == reflect.Slice || reflectPropertyValue.Kind() == reflect.Array {
			v.ValidateValue(reflectPropertyValue.Interface(), constraintList)

			continue
		}

		v.doProcessConstraintValidate(constraintList, propertyValue, root, property)
	}
}

func (v *SimpleValidator) doProcessConstraintValidate(
	constraintList AssertListValue,
	value interface{},
	root, property string,
) {
	for _, constraint := range constraintList {
		v.setConstraintMainData(constraint, root, property)

		for _, validator := range constraint.ProcessValidators() {
			validator.Process(constraint, value, v.exception)
		}
	}
}

func (v *SimpleValidator) setConstraintMainData(
	constraint contract.ConstraintMainDataInterface,
	root, property string,
) {
	constraint.SetRoot(root)
	constraint.SetPropertyPath(property)
}
