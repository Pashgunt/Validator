package validator

import (
	"github.com/Pashgunt/Validator/internal/contract"
	structhelper "github.com/Pashgunt/Validator/internal/helper/struct"
	"github.com/Pashgunt/Validator/internal/violation"
	"reflect"
)

type ValidatorInterface interface {
	Validate(value interface{}, constraints Collection)
}

type ValidatorExceptionInterface interface {
	Exception() contract.ValidationFailedExceptionInterface
}

type SimpleValidator struct {
	exception contract.ValidationFailedExceptionInterface
}

func (v *SimpleValidator) Exception() contract.ValidationFailedExceptionInterface {
	return v.exception
}

func NewSimpleValidator() *SimpleValidator {
	return &SimpleValidator{exception: &violation.ValidationFailedException{}}
}

func (v *SimpleValidator) Validate(value interface{}, constraints Collection) {
	reflectValue := reflect.ValueOf(value)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	if reflectValue.Kind() == reflect.Struct {
		v.processInitValidate(structhelper.GetFillValidateData(reflectValue), constraints, reflectValue.String())

		return
	}

	switch data := value.(type) {
	case map[string]interface{}:
		v.processInitValidate(data, constraints, reflectValue.String())
		break
	}
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

		for _, constraint := range constraintList {
			v.setConstraintMainData(constraint, root, property)
			for _, validator := range constraint.ProcessValidators() {
				validator.Process(constraint, propertyValue, v.exception)
			}
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
