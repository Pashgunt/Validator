package validator

import (
	"github.com/Pashgunt/Validator/internal/cache"
	"github.com/Pashgunt/Validator/internal/contract"
	"github.com/Pashgunt/Validator/internal/enum"
	maphelper "github.com/Pashgunt/Validator/internal/helper/map"
	structhelper "github.com/Pashgunt/Validator/internal/helper/struct"
	"github.com/Pashgunt/Validator/internal/service"
	"github.com/Pashgunt/Validator/internal/violation"
	"github.com/Pashgunt/Validator/pkg/interface"
	"github.com/Pashgunt/Validator/pkg/usecase"
	"reflect"
)

type ValidatorInterface interface {
	Validate(value interface{}, constraints Collection)
	ValidateValue(value interface{}, constraints AssertListValue)
}

type SimpleValidator struct {
	exception pkginterface.ValidationFailedExceptionInterface
	cache     contract.CacheInterface
}

func (v *SimpleValidator) Exception() pkginterface.ValidationFailedExceptionInterface {
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
			panic("Expect structure")
		}

		for i := 0; i < reflectValue.NumField(); i++ {
			tag := reflectValue.Type().Field(i).Tag.Get(enum.KeyAssert)

			if tag == "" {
				continue
			}

			v.doProcessConstraintValidate(
				v.getOrCreateValidator(tag, reflectValue.Type().Field(i).Name),
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
	case reflect.Struct:
		v.processInitValidate(
			structhelper.GetFillValidateData(reflectValue),
			*constraints,
			reflectValue.String(),
		)
	case reflect.Map:
		v.processInitValidate(
			maphelper.GetFillValidateData(reflectValue),
			*constraints,
			reflectValue.String(),
		)
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

		for _, validatorItem := range constraint.ProcessValidators() {
			validatorItem.Process(constraint, value, v.exception)
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

func (v *SimpleValidator) getOrCreateValidator(tag string, fieldName string) AssertListValue {
	var validators AssertListValue

	for _, tagItem := range service.GetTags(tag) {
		switch tagItem {
		case string(enum.NotBlank):
			usecase.NotBlankTag(fieldName, v.cache, tagItem)
		}
	}

	return validators
}
