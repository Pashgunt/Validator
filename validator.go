package validator

import (
	structhelper "github.com/Pashgunt/Validator/internal/helper/struct"
	"reflect"
)

type ValidatorInterface interface {
	Validate(value interface{}, constraints Collection)
}

type SimpleValidator struct{}

func NewSimpleValidator() *SimpleValidator {
	return &SimpleValidator{}
}

func (v *SimpleValidator) Validate(value interface{}, constraints Collection) {
	reflectValue := reflect.ValueOf(value)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	if reflectValue.Kind() == reflect.Struct {
		v.processInitValidate(structhelper.GetFillValidateData(reflectValue), constraints)

		return
	}

	switch data := value.(type) {
	case map[string]interface{}:
		v.processInitValidate(data, constraints)
		break
	}
}

func (v *SimpleValidator) processInitValidate(data map[string]interface{}, constraints Collection) {
	for property, propertyValue := range data {
		constraintList, isset := constraints.Asserts()[property]

		if !isset {
			continue
		}

		for _, constraint := range constraintList {
			for _, validator := range constraint.ProcessValidators() {
				validator.Process(constraint, propertyValue)
			}
		}
	}
}
